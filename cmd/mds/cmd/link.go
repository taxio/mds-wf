package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"golang.org/x/xerrors"
)

func NewLinkSubCmd(ctx *CommandContext) *cobra.Command {
	linkCmd := &cobra.Command{
		Use:   "link",
		Short: "generate link string for markdown format",
		Long:  "generate link string for markdown format",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				return xerrors.New("arguments not correct")
			}
			link, err := generateLink(args[0], args[1])
			if err != nil {
				return xerrors.Errorf(":%w", err)
			}
			_, err = fmt.Fprintln(ctx.Out, link)
			if err != nil {
				return xerrors.Errorf(":%w", err)
			}

			return nil
		},
	}

	return linkCmd
}

func generateLink(title, link string) (string, error) {
	return fmt.Sprintf("[%s](%s)", title, link), nil
}
