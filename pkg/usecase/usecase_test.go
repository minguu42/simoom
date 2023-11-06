package usecase

import (
	"context"
	"log"
	"testing"

	"github.com/minguu42/simoom/pkg/config"
	"github.com/minguu42/simoom/pkg/infra/mysql"
)

var tc *mysql.Client

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

	if err := mysql.InitAllData(context.Background(), tc); err != nil {
		log.Fatalf("%+v", err)
	}

	m.Run()
}
