package cmdutil

import "github.com/minguu42/simoom/gen/simoompb/v1/simoompbconnect"

type Credentials struct {
	AccessToken  string
	RefreshToken string
}

type Core struct {
	Client      simoompbconnect.SimoomServiceClient
	Credentials Credentials
}
