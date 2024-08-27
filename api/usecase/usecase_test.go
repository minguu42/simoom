package usecase_test

import (
	"context"
	"log"
	"path"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/minguu42/simoom/api/adapter/mysql"
	"github.com/minguu42/simoom/api/config"
	"github.com/minguu42/simoom/api/domain/model"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var (
	tc   *mysql.Client
	tctx = model.ContextWithUser(context.Background(), model.User{
		ID:    "user_01",
		Name:  "ユーザ1",
		Email: "testuser1@example.com",
	})
	tctxUser2 = model.ContextWithUser(context.Background(), model.User{
		ID:    "user_02",
		Name:  "ユーザ2",
		Email: "testuser2@example.com",
	})
	fixtures *testfixtures.Loader
)

func TestMain(m *testing.M) {
	ctx := context.Background()
	mysqlC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image: "mysql:8.0.28",
			Env: map[string]string{
				"MYSQL_DATABASE":             "simoomdb_test",
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

	port, err := mysqlC.MappedPort(ctx, "3306/tcp")
	if err != nil {
		log.Fatalf("failed to get externally mapped port: %s", err)
	}
	tc, err = mysql.NewClient(config.DB{
		Host:               "localhost",
		Port:               port.Int(),
		Database:           "simoomdb_test",
		User:               "root",
		Password:           "",
		ConnMaxLifetimeMin: 5,
		MaxOpenConns:       25,
		MaxIdleConns:       25,
	})
	if err != nil {
		log.Fatalf("failed to create test mysql client: %s", err)
	}
	defer tc.Close()

	mysql.Migrate(tc)
	_, f, _, _ := runtime.Caller(0)
	fixtures = mysql.NewFixtureLoader(tc, filepath.Join(path.Dir(f), "testdata"))
	if err := fixtures.Load(); err != nil {
		log.Fatalf("failed to load test data: %s", err)
	}

	m.Run()
}
