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
		log.Fatalf("env.Load failed: %s", err)
	}
	tc, err = NewClient(env.MySQL{
		Host:               "localhost",
		Port:               appEnv.MySQL.Port,
		Database:           fmt.Sprintf("%s_test", appEnv.MySQL.Database),
		User:               appEnv.MySQL.User,
		Password:           appEnv.MySQL.Password,
		ConnMaxLifetimeMin: appEnv.MySQL.ConnMaxLifetimeMin,
		MaxOpenConns:       appEnv.MySQL.MaxOpenConns,
		MaxIdleConns:       appEnv.MySQL.MaxIdleConns,
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
