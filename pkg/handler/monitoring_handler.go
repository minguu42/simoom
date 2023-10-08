package handler

import (
	"context"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/gen/simoompb/v1"
	"github.com/minguu42/simoom/pkg/usecase"
)

type monitoringHandler struct {
	uc usecase.MonitoringUsecase
}

func (h monitoringHandler) CheckHealth(_ context.Context, _ *connect.Request[simoompb.CheckHealthRequest]) (*connect.Response[simoompb.CheckHealthResponse], error) {
	out := h.uc.CheckHealth()
	return connect.NewResponse(&simoompb.CheckHealthResponse{Revision: out.Revision}), nil
}
