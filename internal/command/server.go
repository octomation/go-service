package command

import (
	"net"
	"net/http"
	_ "net/http/pprof"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.octolab.org/errors"
	"go.octolab.org/sync"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	v1 "go.octolab.org/template/service/api/rpc/v1"
	"go.octolab.org/template/service/api/rpc/v1/v1connect"
	"go.octolab.org/template/service/internal/config"
	"go.octolab.org/template/service/internal/server"
)

// NewServer returns the new server command.
func NewServer(cnf *config.Service) *cobra.Command {
	var path string
	v := viper.New()
	v.AddConfigPath(".")

	command := cobra.Command{
		Use:   "server",
		Short: "server of the service",
		Long:  "Execute remote client commands.",

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
	command.AddCommand(
		Run(&cnf.Server),
	)
	return &command
}

func Run(cnf *config.Server) *cobra.Command {
	command := cobra.Command{
		Use:   "run",
		Short: "run the server",
		Long:  "Start listening required protocols.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cnf.Twirp.IsEnabled() {
				twirp := server.Twirp()
				mux := http.NewServeMux()
				mux.Handle(twirp.PathPrefix(), twirp)
				go func() {
					cmd.Println("twirp server starts listening", cnf.Twirp.BaseURL())
					cmd.Println("twirp server status:", http.ListenAndServe(
						cnf.Twirp.Address,
						mux,
					))
				}()
			}
			if cnf.GRPC.IsEnabled() {
				listener, err := net.Listen("tcp", cnf.GRPC.Address)
				if err != nil {
					return err
				}
				srv := grpc.NewServer()
				v1.RegisterGreeterServiceServer(srv, new(server.GRPC))
				go func() {
					cmd.Println("grpc server starts listening", "tcp://"+cnf.GRPC.Address)
					cmd.Println("grpc server status:", srv.Serve(listener))
				}()
				if cnf.Gateway.IsEnabled() {
					mux := runtime.NewServeMux()
					if err := v1.RegisterGreeterServiceHandlerFromEndpoint(
						cmd.Context(),
						mux,
						cnf.GRPC.Address,
						[]grpc.DialOption{
							grpc.WithTransportCredentials(insecure.NewCredentials()),
						},
					); err != nil {
						return err
					}
					go func() {
						cmd.Println("gateway starts listening", cnf.Gateway.BaseURL())
						cmd.Println("gateway status:", http.ListenAndServe(
							cnf.Gateway.Address,
							mux,
						))
					}()
				}
			}
			if cnf.Profile.IsEnabled() {
				mux := http.DefaultServeMux
				go func() {
					cmd.Println("profiler starts listening", cnf.Profile.BaseURL())
					cmd.Println("profiler status:", http.ListenAndServe(
						cnf.Profile.Address,
						mux,
					))
				}()
			}

			mux := http.NewServeMux()
			path, handler := v1connect.NewGreeterServiceHandler(new(server.Connect))
			mux.Handle(path, handler)
			srv := &http.Server{
				Addr:    cnf.Connect.Address,
				Handler: h2c.NewHandler(mux, new(http2.Server)),
			}
			go func() {
				cmd.Println("rpc server starts listening", cnf.Connect.BaseURL())
				cmd.Println("rpc server status:", srv.ListenAndServe())
			}()
			err := sync.Termination().Wait(cmd.Context())
			if errors.Is(err, sync.ErrSignalTrapped) {
				cmd.Println("shutting down rpc server:", srv.Shutdown(cmd.Context()))
				return nil
			}
			return err
		},
	}
	flags := command.Flags()
	flags.StringVar(&cnf.Connect.Address, "host", "", "remote rpc host")
	return &command
}
