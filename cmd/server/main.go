package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/minguu42/simoom/pkg/applog"
	"github.com/minguu42/simoom/pkg/config"
	"github.com/minguu42/simoom/pkg/handler"
	"github.com/minguu42/simoom/pkg/infra/jwtauth"
	"github.com/minguu42/simoom/pkg/infra/mysql"
	"github.com/minguu42/simoom/pkg/infra/ulidgen"
)

func main() {
	time.Local = time.UTC
	applog.InitDefault()

	if err := mainRun(); err != nil {
		applog.Errorf("failed to run server: %s", err)
		os.Exit(1)
	}
}

func mainRun() error {
	conf, err := config.Load()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	c, err := mysql.NewClient(conf.DB)
	if err != nil {
		return fmt.Errorf("failed to create mysql client: %w", err)
	}
	defer c.Close()

	s := &http.Server{
		Addr:              net.JoinHostPort(conf.API.Host, strconv.Itoa(conf.API.Port)),
		Handler:           handler.New(jwtauth.Authenticator{}, c, conf, ulidgen.Generator{}),
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}
	go func() {
		applog.Infof("Start accepting requests")
		if err := s.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			applog.Errorf("failed to listen and handle requests: %s", err)
			return
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)
	<-quit
	if err := s.Shutdown(context.Background()); err != nil {
		return fmt.Errorf("failed to shutdown server: %w", err)
	}
	applog.Infof("Stop accepting requests")

	return nil
}
