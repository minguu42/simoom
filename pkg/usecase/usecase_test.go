package usecase_test

import (
	"context"
	"log"
	"testing"

	"github.com/minguu42/simoom/pkg/config"
	"github.com/minguu42/simoom/pkg/domain/auth"
	"github.com/minguu42/simoom/pkg/infra/mysql"
	"github.com/minguu42/simoom/pkg/infra/ulidgen"
	"github.com/minguu42/simoom/pkg/usecase"
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
	tc, err = mysql.NewClient(config.DB{
		Host:               "localhost",
		Port:               3306,
		Database:           "testdb",
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
	project = usecase.NewProject(tc, ulidgen.Generator{})
	step = usecase.NewStep(tc, ulidgen.Generator{})
	tag = usecase.NewTag(tc, ulidgen.Generator{})
	task = usecase.NewTask(tc, ulidgen.Generator{})

	mysql.InitAllData(tc)

	m.Run()
}
