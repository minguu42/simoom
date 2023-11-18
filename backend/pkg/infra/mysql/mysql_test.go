package mysql

import (
	"context"
	"fmt"
	"log"
	"testing"

	config2 "github.com/minguu42/simoom/backend/pkg/config"
)

var tc *Client

func TestMain(m *testing.M) {
	conf, err := config2.Load()
	if err != nil {
		log.Fatalf("%+v", err)
	}

	mysql := conf.MySQL
	tc, err = NewClient(config2.MySQL{
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
	if err := InitAllData(ctx, tc); err != nil {
		log.Fatalf("%+v", err)
	}

	m.Run()
}
