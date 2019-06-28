package cmd

import "github.com/spf13/cobra"

func NewRootCmd(args []string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "mds-wf",
		Short: "markdown shortcut utility",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	rootCmd.SetArgs(args)

	subCmds := []*cobra.Command{
		NewLinkSubCmd(),
		NewSummarySubCmd(),
	}
	rootCmd.AddCommand(subCmds...)

	return rootCmd
}
