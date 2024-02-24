package cmdutil

import "github.com/minguu42/simoom/lib/go/simoompb/v1/simoompbconnect"

type Core struct {
	Client      simoompbconnect.SimoomServiceClient
	Credentials Credentials
}
