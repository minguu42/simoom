// Package handler はハンドラを定義する
package handler

import (
	"errors"
	"net/http"
	"time"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/pkg/config"
	"github.com/minguu42/simoom/pkg/domain/auth"
	"github.com/minguu42/simoom/pkg/domain/model"
	"github.com/minguu42/simoom/pkg/domain/repository"
	"github.com/minguu42/simoom/pkg/handler/interceptor"
	"github.com/minguu42/simoom/pkg/simoompb/v1"
	"github.com/minguu42/simoom/pkg/simoompb/v1/simoompbconnect"
	"github.com/minguu42/simoom/pkg/usecase"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type handler struct {
	auth       usecase.Auth
	monitoring usecase.Monitoring
	project    usecase.Project
	step       usecase.Step
	tag        usecase.Tag
	task       usecase.Task
}

// New はハンドラを生成する
func New(authenticator auth.Authenticator, repo repository.Repository, conf config.Config, idgen model.IDGenerator) http.Handler {
	opt := connect.WithInterceptors(
		interceptor.NewSetContext(),
		interceptor.NewJudgeError(),
		interceptor.NewRecordAccess(),
		interceptor.NewAuthenticate(authenticator, conf.Auth.AccessTokenSecret),
	)

	mux := http.NewServeMux()
	mux.Handle(simoompbconnect.NewSimoomServiceHandler(handler{
		auth:       usecase.NewAuth(authenticator, repo, conf.Auth, idgen),
		monitoring: usecase.Monitoring{},
		project:    usecase.NewProject(repo, idgen),
		step:       usecase.NewStep(repo, idgen),
		tag:        usecase.NewTag(repo, idgen),
		task:       usecase.NewTask(repo, idgen),
	}, opt))

	return h2c.NewHandler(mux, &http2.Server{})
}

func newErrInvalidArgument(message string) *connect.Error {
	return connect.NewError(connect.CodeInvalidArgument, errors.New(message))
}

func newDate(t *time.Time) *simoompb.Date {
	if t != nil {
		return &simoompb.Date{
			Year:  uint32(t.Year()),
			Month: uint32(t.Month()),
			Day:   uint32(t.Day()),
		}
	}
	return nil
}

func newTimestamp(t *time.Time) *timestamppb.Timestamp {
	if t != nil {
		timestamppb.New(*t)
	}
	return nil
}
