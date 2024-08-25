package main_test

import (
	"context"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/minguu42/simoom/api/config"
	"github.com/minguu42/simoom/api/factory"
	"github.com/minguu42/simoom/api/handler"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const database = "simoomdb_test"

var port int

func TestMain(m *testing.M) {
	ctx := context.Background()
	mysqlC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image: "mysql:8.0.28",
			Env: map[string]string{
				"MYSQL_DATABASE":             database,
				"MYSQL_ALLOW_EMPTY_PASSWORD": "yes",
			},
			ExposedPorts: []string{"3306/tcp"},
			WaitingFor:   wait.ForListeningPort("3306/tcp"),
		},
		Started: true,
	})
	if err != nil {
		log.Fatalf("failed to create mysql container: %s", err)
	}
	defer func() {
		if err := mysqlC.Terminate(ctx); err != nil {
			log.Fatalf("failed to terminate mysql container: %s", err)
		}
	}()

	portNat, err := mysqlC.MappedPort(ctx, "3306/tcp")
	if err != nil {
		log.Fatalf("failed to get externally mapped port: %s", err)
	}
	port = portNat.Int()

	m.Run()
}

func TestAPI(t *testing.T) {
	f, err := factory.New(config.Config{
		Auth: config.Auth{
			AccessTokenExpiryHour:  2,
			RefreshTokenExpiryHour: 168,
			AccessTokenSecret:      "some-access-token-secret",
			RefreshTokenSecret:     "some-refresh-token-secret",
		},
		DB: config.DB{
			Host:               "localhost",
			Port:               port,
			Database:           database,
			User:               "root",
			Password:           "",
			ConnMaxLifetimeMin: 5,
			MaxOpenConns:       25,
			MaxIdleConns:       25,
		},
	})
	require.NoError(t, err)

	h, err := handler.New(f, 5*time.Second)
	require.NoError(t, err)

	ts := httptest.NewServer(h)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/simoompb.v1.SimoomService/CheckHealth?encoding=json&message=%7b%7d")
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	data, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	assert.Equal(t, `{"revision":"xxxxxxx"}`, string(data))
}
