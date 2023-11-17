package command

import (
	"net/http"
	_ "net/http/pprof"

	"github.com/spf13/cobra"
	"go.octolab.org/unsafe"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"go.octolab.org/template/service/internal/api/service/v1/servicev1connect"
	"go.octolab.org/template/service/internal/config"
	"go.octolab.org/template/service/internal/server"
)

// NewServer returns the new server command.
func NewServer() *cobra.Command {
	var cnf config.Server

	command := cobra.Command{
		Use:   "pong",
		Short: "server of service",
		Long:  "Execute remote client commands.",

		Args: cobra.NoArgs,

		SilenceErrors: false,
		SilenceUsage:  true,

		RunE: func(cmd *cobra.Command, args []string) error {
			mux := http.NewServeMux()
			path, handler := servicev1connect.NewGreeterServiceHandler(
				new(server.Connect),
			)
			twirp := server.Twirp()
			mux.Handle(path, handler)
			mux.Handle(twirp.PathPrefix(), twirp)

			if cnf.Debug {
				go func() {
					unsafe.Ignore(http.ListenAndServe("localhost:3360", http.DefaultServeMux))
				}()
			}

			return http.ListenAndServe(cnf.Host, h2c.NewHandler(mux, new(http2.Server)))
		},
	}
	flags := command.Flags()
	flags.StringVar(&cnf.Host, "host", "localhost:8890", "server address")
	flags.BoolVar(&cnf.Debug, "debug", false, "enable pprof")

	command.AddCommand( /* add related commands */ )
	return &command
}
