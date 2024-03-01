package cmdutil

import "github.com/minguu42/simoom/lib/go/simoompb/v1/simoompbconnect"

type Factory struct {
	Client      simoompbconnect.SimoomServiceClient
	Credentials Credentials
}
