// Package handler はハンドラを定義する
package handler

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"connectrpc.com/connect"
	"github.com/bufbuild/protovalidate-go"
	"github.com/minguu42/simoom/api/config"
	"github.com/minguu42/simoom/api/domain/auth"
	"github.com/minguu42/simoom/api/domain/model"
	"github.com/minguu42/simoom/api/domain/repository"
	"github.com/minguu42/simoom/api/handler/interceptor"
	"github.com/minguu42/simoom/api/usecase"
	"github.com/minguu42/simoom/lib/go/simoompb/v1"
	"github.com/minguu42/simoom/lib/go/simoompb/v1/simoompbconnect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var ErrInvalidRequest = errors.New("the entered value is incorrect")

type handler struct {
	validator  *protovalidate.Validator
	auth       usecase.Auth
	monitoring usecase.Monitoring
	project    usecase.Project
	step       usecase.Step
	tag        usecase.Tag
	task       usecase.Task
}

// New はハンドラを生成する
func New(authenticator auth.Authenticator, repo repository.Repository, conf config.Config, idgen model.IDGenerator) (http.Handler, error) {
	validator, err := protovalidate.New()
	if err != nil {
		return nil, fmt.Errorf("failed to create validator: %w", err)
	}

	opt := connect.WithInterceptors(
		interceptor.NewSetContext(),
		interceptor.NewArrangeErrorAndRecordAccess(),
		interceptor.NewAuthenticate(authenticator, conf.Auth.AccessTokenSecret),
	)

	mux := http.NewServeMux()
	mux.Handle(simoompbconnect.NewSimoomServiceHandler(handler{
		validator:  validator,
		auth:       usecase.NewAuth(authenticator, repo, conf.Auth, idgen),
		monitoring: usecase.Monitoring{},
		project:    usecase.NewProject(repo, idgen),
		step:       usecase.NewStep(repo, idgen),
		tag:        usecase.NewTag(repo, idgen),
		task:       usecase.NewTask(repo, idgen),
	}, opt))

	return h2c.NewHandler(mux, &http2.Server{}), nil
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
