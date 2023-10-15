package mysql

import (
	"context"
	"log"
	"testing"

	"github.com/minguu42/simoom/pkg/env"
)

var tc *Client

func TestMain(m *testing.M) {
	var err error
	tc, err = NewClient(env.MySQL{
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

	ctx := context.Background()
	if err := initAllData(ctx, tc.db); err != nil {
		log.Fatalf("%+v", err)
	}

	m.Run()
}
