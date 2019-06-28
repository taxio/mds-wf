package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"golang.org/x/xerrors"
)

func NewLinkSubCmd() *cobra.Command {
	linkCmd := &cobra.Command{
		Use:   "link",
		Short: "generate link string for markdown format",
		Long:  "generate link string for markdown format",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				return xerrors.New("arguments not correct")
			}
			fmt.Printf("[%s](%s)", args[0], args[1])
			return nil
		},
	}

	return linkCmd
}
