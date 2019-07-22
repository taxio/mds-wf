package cmd

import "github.com/spf13/cobra"

func NewRootCmd(args []string, ctx *CommandContext) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "mds",
		Short: "markdown shortcut utility",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	rootCmd.SetArgs(args)

	subCmds := []*cobra.Command{
		NewLinkSubCmd(ctx),
		NewSummarySubCmd(ctx),
		NewTableSubCmd(ctx),
	}
	rootCmd.AddCommand(subCmds...)

	return rootCmd
}
