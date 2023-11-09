// Package handler はハンドラを定義する
package handler

import (
	"net/http"
	"time"

	"connectrpc.com/connect"
	"github.com/cockroachdb/errors"
	"github.com/minguu42/simoom/gen/simoompb/v1"
	"github.com/minguu42/simoom/gen/simoompb/v1/simoompbconnect"
	"github.com/minguu42/simoom/pkg/config"
	"github.com/minguu42/simoom/pkg/domain/auth"
	"github.com/minguu42/simoom/pkg/domain/repository"
	"github.com/minguu42/simoom/pkg/handler/interceptor"
	"github.com/minguu42/simoom/pkg/usecase"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type handler struct {
	auth       usecase.AuthUsecase
	monitoring usecase.MonitoringUsecase
	project    usecase.ProjectUsecase
	step       usecase.StepUsecase
	tag        usecase.TagUsecase
	task       usecase.TaskUsecase
}

// New はハンドラを生成する
func New(authenticator auth.Authenticator, repo repository.Repository, conf config.Env) http.Handler {
	opt := connect.WithInterceptors(
		interceptor.NewSetContext(),
		interceptor.NewErrorJudge(),
		interceptor.NewAccessLog(),
		interceptor.NewAuth(authenticator, conf.API.AccessTokenSecret),
	)

	mux := http.NewServeMux()
	mux.Handle(simoompbconnect.NewSimoomServiceHandler(handler{
		auth:       usecase.AuthUsecase{Repo: repo, Env: conf},
		monitoring: usecase.MonitoringUsecase{},
		project:    usecase.NewProject(repo),
		step:       usecase.NewStep(repo),
		tag:        usecase.NewTag(repo),
		task:       usecase.NewTask(repo),
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
