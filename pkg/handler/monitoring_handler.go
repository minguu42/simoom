package handler

import (
	"context"
	"runtime/debug"
	"slices"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/gen/simoompb/v1"
)

type monitoringHandler struct{}

func (h monitoringHandler) CheckHealth(_ context.Context, _ *connect.Request[simoompb.CheckHealthRequest]) (*connect.Response[simoompb.CheckHealthResponse], error) {
	revision := "xxxxxxx"
	if info, ok := debug.ReadBuildInfo(); ok {
		if i := slices.IndexFunc(info.Settings, func(s debug.BuildSetting) bool {
			return s.Key == "vcs.revision"
		}); i != -1 {
			revision = info.Settings[i].Value[:len(revision)]
		}
	}

	return connect.NewResponse(&simoompb.CheckHealthResponse{Revision: revision}), nil
}
