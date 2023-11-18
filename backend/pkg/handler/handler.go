// Package handler はハンドラを定義する
package handler

import (
	"net/http"
	"time"

	"connectrpc.com/connect"
	"github.com/cockroachdb/errors"
	"github.com/minguu42/simoom/backend/pkg/config"
	"github.com/minguu42/simoom/backend/pkg/domain/auth"
	"github.com/minguu42/simoom/backend/pkg/domain/repository"
	interceptor2 "github.com/minguu42/simoom/backend/pkg/handler/interceptor"
	usecase2 "github.com/minguu42/simoom/backend/pkg/usecase"
	"github.com/minguu42/simoom/library/simoompb/v1"
	"github.com/minguu42/simoom/library/simoompb/v1/simoompbconnect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type handler struct {
	auth       usecase2.Auth
	monitoring usecase2.Monitoring
	project    usecase2.Project
	step       usecase2.Step
	tag        usecase2.Tag
	task       usecase2.Task
}

// New はハンドラを生成する
func New(authenticator auth.Authenticator, repo repository.Repository, conf config.Config) http.Handler {
	opt := connect.WithInterceptors(
		interceptor2.NewSetContext(),
		interceptor2.NewJudgeError(),
		interceptor2.NewRecordAccess(),
		interceptor2.NewAuthenticate(authenticator, conf.Auth.AccessTokenSecret),
	)

	mux := http.NewServeMux()
	mux.Handle(simoompbconnect.NewSimoomServiceHandler(handler{
		auth:       usecase2.NewAuth(authenticator, repo, conf.Auth),
		monitoring: usecase2.Monitoring{},
		project:    usecase2.NewProject(repo),
		step:       usecase2.NewStep(repo),
		tag:        usecase2.NewTag(repo),
		task:       usecase2.NewTask(repo),
	}, opt))

	return h2c.NewHandler(mux, &http2.Server{})
}

func newErrInvalidArgument(message string) *connect.Error {
	return connect.NewError(connect.CodeInvalidArgument, errors.New(message))
}

func newDate(t *time.Time) *simoompb.simoompb {
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
