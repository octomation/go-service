package cmd

import "github.com/spf13/cobra"

// New returns the new server command.
func NewServer() *cobra.Command {
	command := cobra.Command{
		Use:   "%template%",
		Short: "%template%",
		Long:  "%template%",

		SilenceErrors: false,
		SilenceUsage:  true,
	}
	/* configure instance */
	command.AddCommand( /* related commands */ )
	return &command
}
