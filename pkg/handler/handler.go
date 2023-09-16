package handler

import (
	"net/http"

	"github.com/minguu42/simoom/gen/simoompb/v1/simoompbconnect"
	"github.com/minguu42/simoom/pkg/domain/repository"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

const userID = "01DXF6DT000000000000000000"

// New はハンドラを生成する
func New(repo repository.Repository) http.Handler {
	mux := http.NewServeMux()
	mux.Handle(simoompbconnect.NewMonitoringServiceHandler(monitoringHandler{}))
	mux.Handle(simoompbconnect.NewProjectServiceHandler(projectHandler{repo: repo}))

	return h2c.NewHandler(mux, &http2.Server{})
}
