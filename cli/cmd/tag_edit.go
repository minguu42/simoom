package cmd

import (
	"context"
	"fmt"
	"io"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/cli/api"
	"github.com/minguu42/simoom/cli/factory"
	"github.com/minguu42/simoom/lib/go/simoompb/v1"
	"github.com/spf13/cobra"
)

type TagEditOpts struct {
	Client api.Client

	ID   string
	Name string
}

func NewCmdTagEdit() *cobra.Command {
	var opts TagEditOpts
	cmd := &cobra.Command{
		Use:   "edit",
		Short: "Edit a tag",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			f := factory.FromContext(cmd.Context())
			opts.Client = f.Client

			opts.ID = args[0]
			return TagEditRun(cmd.Context(), f.Out, opts)
		},
	}
	cmd.Flags().StringVar(&opts.Name, "name", "", "tag name")
	return cmd
}

func TagEditRun(ctx context.Context, out io.Writer, opts TagEditOpts) error {
	var name *string
	if opts.Name != "" {
		name = &opts.Name
	}
	resp, err := opts.Client.UpdateTag(ctx, connect.NewRequest(&simoompb.UpdateTagRequest{
		Id:   opts.ID,
		Name: name,
	}))
	if err != nil {
		return fmt.Errorf("failed to call UpdateTag method: %w", err)
	}

	fmt.Fprintf(out, "Tag %s (%s) edited\n", resp.Msg.Name, resp.Msg.Id)
	return nil
}
