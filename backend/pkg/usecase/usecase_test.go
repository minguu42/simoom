package usecase_test

import (
	"context"
	"log"
	"testing"

	"github.com/minguu42/simoom/backend/pkg/config"
	"github.com/minguu42/simoom/backend/pkg/domain/auth"
	"github.com/minguu42/simoom/backend/pkg/infra/mysql"
	"github.com/minguu42/simoom/backend/pkg/usecase"
)

var (
	tc      *mysql.Client
	ctx     = auth.SetUserID(context.Background(), "user_01")
	project usecase.Project
	step    usecase.Step
	tag     usecase.Tag
	task    usecase.Task
)

func TestMain(m *testing.M) {
	var err error
	tc, err = mysql.NewClient(config.MySQL{
		Host:               "localhost",
		Port:               3306,
		Database:           "simoomdb_test",
		User:               "root",
		Password:           "",
		ConnMaxLifetimeMin: 5,
		MaxOpenConns:       25,
		MaxIdleConns:       25,
	})
	if err != nil {
		log.Fatalf("%+v", err)
	}
	defer tc.Close()
	project = usecase.NewProject(tc)
	step = usecase.NewStep(tc)
	tag = usecase.NewTag(tc)
	task = usecase.NewTask(tc)

	if err := mysql.InitAllData(context.Background(), tc); err != nil {
		log.Fatalf("%+v", err)
	}

	m.Run()
}