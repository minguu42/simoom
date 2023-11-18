package handler

import (
	"context"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/lib/simoompb/v1"
)

func (h handler) CheckHealth(_ context.Context, _ *connect.Request[simoompb.CheckHealthRequest]) (*connect.Response[simoompb.CheckHealthResponse], error) {
	out := h.monitoring.CheckHealth()
	return connect.NewResponse(&simoompb.CheckHealthResponse{Revision: out.Revision}), nil
}
