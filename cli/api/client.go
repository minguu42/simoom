package api

import (
	"context"
	"fmt"
	"net/http"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/lib/go/simoompb/v1/simoompbconnect"
)

type Client interface {
	simoompbconnect.SimoomServiceClient
	CheckCredentials() bool
	GetRefreshToken() string
}

type ServiceClient struct {
	simoompbconnect.SimoomServiceClient
	Credentials Credentials
}

func NewClient() (*ServiceClient, error) {
	credentials, err := NewCredentials()
	if err != nil {
		return nil, fmt.Errorf("failed to create credentials: %w", err)
	}

	authInterceptor := connect.UnaryInterceptorFunc(func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			req.Header().Set("Authorization", fmt.Sprintf("Bearer %s", credentials.AccessToken))
			return next(ctx, req)
		}
	})
	c := simoompbconnect.NewSimoomServiceClient(
		http.DefaultClient,
		"http://localhost:8080",
		connect.WithInterceptors(authInterceptor),
	)
	return &ServiceClient{
		SimoomServiceClient: c,
		Credentials:         credentials,
	}, nil
}

func (c *ServiceClient) CheckCredentials() bool {
	if c.Credentials.AccessToken == "" && c.Credentials.RefreshToken == "" {
		return false
	}
	return true
}

func (c *ServiceClient) GetRefreshToken() string {
	return c.Credentials.RefreshToken
}
