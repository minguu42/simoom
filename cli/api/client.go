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
	credentials credentials
}

func NewClient(profile string) (*ServiceClient, error) {
	creds, err := newCredentials(profile)
	if err != nil {
		return nil, fmt.Errorf("failed to create credentials: %w", err)
	}

	authInterceptor := connect.UnaryInterceptorFunc(func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			req.Header().Set("Authorization", fmt.Sprintf("Bearer %s", creds.AccessToken))
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
		credentials:         creds,
	}, nil
}

// CheckCredentials は認証情報を確認する
// 認証情報が正しい場合は true を返し、正しくない場合は false を返す
func (c *ServiceClient) CheckCredentials() bool {
	if c.credentials.Profile == "" || (c.credentials.AccessToken == "" && c.credentials.RefreshToken == "") {
		return false
	}
	return true
}

// GetRefreshToken はリフレッシュトークンを返す
func (c *ServiceClient) GetRefreshToken() string {
	return c.credentials.RefreshToken
}
