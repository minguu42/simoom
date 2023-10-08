// Package handler はハンドラを定義する
package handler

import (
	"net/http"
	"time"

	"connectrpc.com/connect"
	"github.com/cockroachdb/errors"
	"github.com/minguu42/simoom/gen/simoompb/v1"
	"github.com/minguu42/simoom/gen/simoompb/v1/simoompbconnect"
	"github.com/minguu42/simoom/pkg/domain/repository"
	"github.com/minguu42/simoom/pkg/usecase"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var errInvalidArgument = connect.NewError(connect.CodeInvalidArgument, errors.New("request contains an error"))

// New はハンドラを生成する
func New(repo repository.Repository) http.Handler {
	mux := http.NewServeMux()
	mux.Handle(simoompbconnect.NewMonitoringServiceHandler(monitoringHandler{}))
	mux.Handle(simoompbconnect.NewProjectServiceHandler(projectHandler{uc: usecase.ProjectUsecase{Repo: repo}}))
	mux.Handle(simoompbconnect.NewStepServiceHandler(stepHandler{uc: usecase.StepUsecase{Repo: repo}}))
	mux.Handle(simoompbconnect.NewTagServiceHandler(tagHandler{uc: usecase.TagUsecase{Repo: repo}}))
	mux.Handle(simoompbconnect.NewTaskServiceHandler(taskHandler{uc: usecase.TaskUsecase{Repo: repo}}))

	return h2c.NewHandler(mux, &http2.Server{})
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
