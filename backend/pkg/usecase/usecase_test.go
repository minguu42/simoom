package usecase_test

import (
	"context"
	"log"
	"testing"

	"github.com/minguu42/simoom/backend/pkg/config"
	"github.com/minguu42/simoom/backend/pkg/domain/auth"
	mysql2 "github.com/minguu42/simoom/backend/pkg/infra/mysql"
	usecase2 "github.com/minguu42/simoom/backend/pkg/usecase"
)

var (
	tc      *mysql2.Client
	ctx     = auth.SetUserID(context.Background(), "user_01")
	project usecase2.Project
	step    usecase2.Step
	tag     usecase2.Tag
	task    usecase2.Task
)

func TestMain(m *testing.M) {
	var err error
	tc, err = mysql2.NewClient(config.MySQL{
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
	project = usecase2.NewProject(tc)
	step = usecase2.NewStep(tc)
	tag = usecase2.NewTag(tc)
	task = usecase2.NewTask(tc)

	if err := mysql2.InitAllData(context.Background(), tc); err != nil {
		log.Fatalf("%+v", err)
	}

	m.Run()
}
