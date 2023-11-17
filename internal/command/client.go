package command

import (
	"net/http"
	"strings"

	"connectrpc.com/connect"
	"github.com/spf13/cobra"

	v1 "go.octolab.org/template/service/internal/api/service/v1"
	"go.octolab.org/template/service/internal/api/service/v1/servicev1connect"
	"go.octolab.org/template/service/internal/config"
)

// NewClient returns the new client command.
func NewClient() *cobra.Command {
	var cnf config.Server

	command := cobra.Command{
		Use:   "ping",
		Short: "client to the service",
		Long:  "Run service commands remotely.",

		Args: cobra.NoArgs,

		SilenceErrors: false,
		SilenceUsage:  true,
	}
	flags := command.PersistentFlags()
	flags.StringVar(&cnf.Host, "host", "localhost:8890", "remote host")

	command.AddCommand(Hello(&cnf), Sign(&cnf))
	return &command
}

func Hello(cnf *config.Server) *cobra.Command {
	return &cobra.Command{
		Use:   "hello",
		Short: "send a message to the service",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			client := servicev1connect.NewGreeterServiceClient(http.DefaultClient, cnf.BaseURL())
			resp, err := client.Hello(cmd.Context(), connect.NewRequest(&v1.HelloRequest{
				Subject: strings.Join(args, " "),
			}))
			if err != nil {
				return err
			}
			cmd.Println(resp.Msg.GetText())
			return nil
		},
	}
}

func Sign(cnf *config.Server) *cobra.Command {
	return &cobra.Command{
		Use:   "sign",
		Short: "send a license request to the service",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			client := servicev1connect.NewGreeterServiceClient(http.DefaultClient, cnf.BaseURL())
			resp, err := client.Sign(cmd.Context(), connect.NewRequest(new(v1.SignRequest)))
			if err != nil {
				return err
			}
			cmd.Println(string(resp.Msg.GetLicense()))
			return nil
		},
	}
}
