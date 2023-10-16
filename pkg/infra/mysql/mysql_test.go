package mysql

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/minguu42/simoom/pkg/env"
)

var tc *Client

func TestMain(m *testing.M) {
	appEnv, err := env.Load()
	if err != nil {
		log.Fatalf("%+v", err)
	}

	mysql := appEnv.MySQL
	tc, err = NewClient(env.MySQL{
		Host:               "localhost",
		Port:               mysql.Port,
		Database:           fmt.Sprintf("%s_test", mysql.Database),
		User:               mysql.User,
		Password:           mysql.Password,
		ConnMaxLifetimeMin: mysql.ConnMaxLifetimeMin,
		MaxOpenConns:       mysql.MaxOpenConns,
		MaxIdleConns:       mysql.MaxIdleConns,
	})
	if err != nil {
		log.Fatalf("%+v", err)
	}
	defer tc.Close()

	ctx := context.Background()
	if err := initAllData(ctx, tc.db); err != nil {
		log.Fatalf("%+v", err)
	}

	m.Run()
}
