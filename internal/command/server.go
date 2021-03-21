package command

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"service"
	"service/internal/config"
)

// NewServer returns the new server command.
func NewServer() *cobra.Command {
	command := cobra.Command{
		Use:   "%template%",
		Short: "%template%",
		Long:  "%template%",

		Args: cobra.NoArgs,

		SilenceErrors: false,
		SilenceUsage:  true,

		RunE: func(cmd *cobra.Command, args []string) error {
			viper.AutomaticEnv()
			viper.SetFs(service.FS)
			viper.SetConfigFile(".env")
			if err := viper.ReadInConfig(); err != nil {
				return err
			}

			var config config.Server
			return viper.Unmarshal(&config)
		},
	}
	/* configure instance */
	command.AddCommand( /* related commands */ )
	return &command
}
