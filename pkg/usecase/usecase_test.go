package usecase_test

import (
	"context"
	"log"
	"testing"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/minguu42/simoom/pkg/config"
	"github.com/minguu42/simoom/pkg/domain/auth"
	"github.com/minguu42/simoom/pkg/infra/mysql"
	"github.com/minguu42/simoom/pkg/infra/ulidgen"
	"github.com/minguu42/simoom/pkg/usecase"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var (
	tc       *mysql.Client
	tctx     = auth.SetUserID(context.Background(), "user_01")
	project  usecase.Project
	step     usecase.Step
	tag      usecase.Tag
	task     usecase.Task
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
	project = usecase.NewProject(tc, ulidgen.Generator{})
	step = usecase.NewStep(tc, ulidgen.Generator{})
	tag = usecase.NewTag(tc, ulidgen.Generator{})
	task = usecase.NewTask(tc, ulidgen.Generator{})

	mysql.Migrate(tc)
	fixtures = mysql.NewFixtureLoader(tc)
	if err := fixtures.Load(); err != nil {
		log.Fatalf("failed to load test data: %s", err)
	}

	m.Run()
}
