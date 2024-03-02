package api

import (
	"context"
	"fmt"
	"net/http"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/lib/go/simoompb/v1/simoompbconnect"
)

type Client struct {
	simoompbconnect.SimoomServiceClient
	Credentials Credentials
}

func NewClient() (*Client, error) {
	credentials, err := NewCredentials()
	if err != nil {
		return nil, fmt.Errorf("failed to create credentials: %w", err)
	}
	c := simoompbconnect.NewSimoomServiceClient(
		http.DefaultClient,
		"http://localhost:8080",
		connect.WithInterceptors(setAccessToken(credentials.AccessToken)),
	)
	return &Client{
		SimoomServiceClient: c,
		Credentials:         credentials,
	}, nil
}

func setAccessToken(token string) connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			req.Header().Set("Authorization", fmt.Sprintf("Bearer %s", token))
			return next(ctx, req)
		}
	}
}
