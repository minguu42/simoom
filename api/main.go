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

	"github.com/minguu42/simoom/api/applog"
	"github.com/minguu42/simoom/api/config"
	"github.com/minguu42/simoom/api/handler"
	"github.com/minguu42/simoom/api/infra/jwtauth"
	"github.com/minguu42/simoom/api/infra/mysql"
	"github.com/minguu42/simoom/api/infra/ulidgen"
)

func main() {
	time.Local = time.UTC
	applog.Init()

	ctx := context.Background()
	if err := mainRun(ctx); err != nil {
		applog.LogApplicationError(ctx, fmt.Sprintf("failed to run server: %s", err))
		os.Exit(1)
	}
}

func mainRun(ctx context.Context) error {
	conf, err := config.Load()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	c, err := mysql.NewClient(conf.DB)
	if err != nil {
		return fmt.Errorf("failed to create mysql client: %w", err)
	}
	defer c.Close()

	h, err := handler.New(jwtauth.Authenticator{}, c, conf, ulidgen.Generator{})
	if err != nil {
		return fmt.Errorf("failed to create handler: %w", err)
	}
	s := &http.Server{
		Addr:              net.JoinHostPort(conf.API.Host, strconv.Itoa(conf.API.Port)),
		Handler:           h,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}
	go func() {
		applog.LogApplicationEvent(ctx, "Start accepting requests")
		if err := s.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			applog.LogApplicationError(ctx, fmt.Sprintf("failed to listen and handle requests: %s", err))
			return
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)
	<-quit
	if err := s.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to shutdown server: %w", err)
	}
	applog.LogApplicationEvent(ctx, "Stop accepting requests")

	return nil
}
