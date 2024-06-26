package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/minguu42/simoom/api/config"
	"github.com/minguu42/simoom/api/factory"
	"github.com/minguu42/simoom/api/handler"
	"github.com/minguu42/simoom/api/logging"
)

func main() {
	time.Local = time.UTC

	ctx := context.Background()
	if err := mainRun(ctx); err != nil {
		logging.Error(ctx, fmt.Sprintf("failed to run server: %s", err))
		os.Exit(1)
	}
}

func mainRun(ctx context.Context) error {
	conf, err := config.Load()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	f, err := factory.New(conf)
	if err != nil {
		return fmt.Errorf("failed to create factory: %w", err)
	}
	defer f.Close()

	h, err := handler.New(f, conf.API.Timeout)
	if err != nil {
		return fmt.Errorf("failed to create handler: %w", err)
	}
	s := &http.Server{
		Addr:              net.JoinHostPort(conf.API.Host, strconv.Itoa(conf.API.Port)),
		Handler:           h,
		ReadTimeout:       conf.API.ReadTimeout,
		ReadHeaderTimeout: conf.API.ReadHeaderTimeout,
	}

	serveErr := make(chan error)
	go func() {
		logging.Event(ctx, "Start accepting requests")
		serveErr <- s.ListenAndServe()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)
	select {
	case err := <-serveErr:
		return fmt.Errorf("failed to listen and serve: %w", err)
	case <-quit:
	}

	if err := s.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to shutdown server: %w", err)
	}
	logging.Event(ctx, "Stop accepting requests")

	return nil
}
