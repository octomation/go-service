package command

import (
	"net/http"
	"strings"

	"connectrpc.com/connect"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	v1 "go.octolab.org/template/service/api/rpc/v1"
	"go.octolab.org/template/service/api/rpc/v1/v1connect"
	"go.octolab.org/template/service/internal/config"
)

// NewClient returns the new client command.
func NewClient(cnf *config.Service) *cobra.Command {
	var path string
	v := viper.New()
	v.AddConfigPath(".")

	command := cobra.Command{
		Use:   "client",
		Short: "client to the service",
		Long:  "Run service commands remotely.",

		Args: cobra.NoArgs,

		SilenceErrors: false,
		SilenceUsage:  true,

		PersistentPreRunE: func(*cobra.Command, []string) error {
			v.SetConfigFile(path)
			if err := v.ReadInConfig(); err != nil {
				return err
			}
			return v.Unmarshal(cnf)
		},
	}
	flags := command.PersistentFlags()
	flags.StringVarP(&path, "config", "c", "config.toml", "path to config file")
	flags.StringVar(&cnf.Server.Connect.Address, "host", "", "remote rpc host")

	command.AddCommand(
		Hello(&cnf.Server.Connect),
		Sign(&cnf.Server.Connect),
	)
	return &command
}

func Hello(cnf *config.Protocol) *cobra.Command {
	return &cobra.Command{
		Use:   "hello",
		Short: "send a message to the service",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			client := v1connect.NewGreeterServiceClient(http.DefaultClient, cnf.BaseURL())
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

func Sign(cnf *config.Protocol) *cobra.Command {
	return &cobra.Command{
		Use:   "sign",
		Short: "send a license request to the service",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			client := v1connect.NewGreeterServiceClient(http.DefaultClient, cnf.BaseURL())
			resp, err := client.Sign(cmd.Context(), connect.NewRequest(new(v1.SignRequest)))
			if err != nil {
				return err
			}
			cmd.Println(string(resp.Msg.GetLicense()))
			return nil
		},
	}
}
