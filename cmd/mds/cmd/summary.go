package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"golang.org/x/xerrors"
)

func NewSummarySubCmd() *cobra.Command {
	summaryCmd := &cobra.Command{
		Use:   "summary",
		Short: "generate summary string for markdown format",
		Long:  "generate summary string for markdown format",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return xerrors.New("arguments not correct")
			}
			fmt.Printf("<details><summary>%s</summary><div></div></details>", args[0])
			return nil
		},
	}

	return summaryCmd
}
