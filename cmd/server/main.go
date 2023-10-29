package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/minguu42/simoom/pkg/applog"
	"github.com/minguu42/simoom/pkg/env"
	"github.com/minguu42/simoom/pkg/handler"
	"github.com/minguu42/simoom/pkg/infra/mysql"
)

func init() {
	applog.InitDefault()
}

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

	s := &http.Server{
		Addr:              net.JoinHostPort(appEnv.API.Host, strconv.Itoa(appEnv.API.Port)),
		Handler:           handler.New(c),
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}
	go func() {
		log.Println("Start accepting requests")
		if err := s.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Printf("%+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)
	<-quit
	if err := s.Shutdown(context.Background()); err != nil {
		log.Fatalf("s.Shutdown failed: %s", err)
	}
	log.Println("Stop accepting requests")
}
