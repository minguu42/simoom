package mysql_test

import (
	"context"
	"errors"
	"log"
	"testing"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/minguu42/simoom/api/config"
	"github.com/minguu42/simoom/api/domain/model"
	"github.com/minguu42/simoom/api/infra/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var (
	tc       *mysql.Client
	fixtures *testfixtures.Loader
)

func TestMain(m *testing.M) {
	ctx := context.Background()
	mysqlC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image: "mysql:8.0.32",
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
	fixtures = mysql.NewFixtureLoader(tc)
	if err := fixtures.Load(); err != nil {
		log.Fatalf("failed to load test data: %s", err)
	}

	m.Run()
}

func TestClient_Transaction(t *testing.T) {
	t.Run("トランザクション中にエラーが発生した場合はロールバックされる", func(t *testing.T) {
		ctx := context.Background()
		t.Cleanup(func() {
			_ = fixtures.Load()
		})
		err := tc.Transaction(ctx, func(transactionCtx context.Context) error {
			err := tc.DeleteProject(transactionCtx, "project_01")
			require.NoError(t, err)
			return errors.New("some error occurred")
		})
		require.Error(t, err)

		if got, err := tc.GetProjectByID(ctx, "project_01"); assert.NoError(t, err) {
			assert.Equal(t, model.Project{
				ID:         "project_01",
				UserID:     "user_01",
				Name:       "プロジェクト1",
				Color:      "#1a2b3c",
				IsArchived: false,
			}, got)
		}
	})
}
