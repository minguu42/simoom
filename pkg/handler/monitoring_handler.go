package handler

import (
	"context"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/gen/simoompb/v1"
)

func (s simoom) CheckHealth(_ context.Context, _ *connect.Request[simoompb.CheckHealthRequest]) (*connect.Response[simoompb.CheckHealthResponse], error) {
	out := s.monitoring.CheckHealth()
	return connect.NewResponse(&simoompb.CheckHealthResponse{Revision: out.Revision}), nil
}
