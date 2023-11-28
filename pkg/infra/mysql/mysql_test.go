package mysql

import (
	"log"
	"testing"

	"github.com/minguu42/simoom/pkg/config"
)

var tc *Client

func TestMain(m *testing.M) {
	var err error
	tc, err = NewClient(config.DB{
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

	InitAllData(tc)

	m.Run()
}
