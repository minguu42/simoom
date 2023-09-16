package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/minguu42/simoom/pkg/env"
	"github.com/minguu42/simoom/pkg/handler"
	"github.com/minguu42/simoom/pkg/infra/mysql"
)

func main() {
	time.Local = time.UTC

	appEnv, err := env.Load()
	if err != nil {
		log.Fatalf("env.Load() failed: %s", err)
	}

	c, err := mysql.NewClient(appEnv.MySQL)
	if err != nil {
		log.Printf("%+v", err)
	}
	defer c.Close()

	log.Println("Start accepting requests")
	addr := fmt.Sprintf("%s:%d", appEnv.API.Host, appEnv.API.Port)
	if err := http.ListenAndServe(addr, handler.New(c)); errors.Is(err, http.ErrServerClosed) {
		log.Println(err)
		return
	}
	log.Println("server started")
}
