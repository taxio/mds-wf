package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"golang.org/x/xerrors"
)

func NewSummarySubCmd(ctx *CommandContext) *cobra.Command {
	summaryCmd := &cobra.Command{
		Use:   "summary",
		Short: "generate summary string for markdown format",
		Long:  "generate summary string for markdown format",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return xerrors.New("arguments not correct")
			}
			smm, err := generateSummaryTmpl(args[0])
			if err != nil {
				return xerrors.Errorf(": %w", err)
			}
			_, err = fmt.Fprintln(ctx.Out, smm)
			if err != nil {
				return xerrors.Errorf(": %w", err)
			}

			return nil
		},
	}

	return summaryCmd
}

func generateSummaryTmpl(title string) (string, error) {
	return fmt.Sprintf("<details><summary>%s</summary><div></div></details>", title), nil
}
