package command

import "github.com/spf13/cobra"

// NewClient returns the new client command.
func NewClient() *cobra.Command {
	command := cobra.Command{
		Use:   "%template%",
		Short: "%template%",
		Long:  "%template%",

		Args: cobra.NoArgs,

		SilenceErrors: false,
		SilenceUsage:  true,
	}
	/* configure instance */
	command.AddCommand( /* related commands */ )
	return &command
}
