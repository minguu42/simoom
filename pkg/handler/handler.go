package handler

import (
	"net/http"
	"time"

	"connectrpc.com/connect"
	"github.com/cockroachdb/errors"
	"github.com/minguu42/simoom/gen/simoompb/v1"
	"github.com/minguu42/simoom/gen/simoompb/v1/simoompbconnect"
	"github.com/minguu42/simoom/pkg/domain/repository"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const userID = "01DXF6DT000000000000000000"

var (
	errInvalidArgument = connect.NewError(connect.CodeInvalidArgument, errors.New("request contains an error"))
	errProjectNotFound = connect.NewError(connect.CodeNotFound, errors.New("the specified project is not found"))
	errTaskNotFound    = connect.NewError(connect.CodeNotFound, errors.New("the specified task is not found"))
	errUnimplemented   = connect.NewError(connect.CodeUnimplemented, errors.New("this RPC method is not yet implemented"))
	errInternal        = connect.NewError(connect.CodeInternal, errors.New("an unintentional error occurred on the server"))
)

// New はハンドラを生成する
func New(repo repository.Repository) http.Handler {
	mux := http.NewServeMux()
	mux.Handle(simoompbconnect.NewMonitoringServiceHandler(monitoringHandler{}))
	mux.Handle(simoompbconnect.NewProjectServiceHandler(projectHandler{repo: repo}))
	mux.Handle(simoompbconnect.NewTaskServiceHandler(taskHandler{repo: repo}))

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
