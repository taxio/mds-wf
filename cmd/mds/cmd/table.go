package cmd

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func NewTableSubCmd(ctx *CommandContext) *cobra.Command {
	tableCmd := &cobra.Command{
		Use:   "table",
		Short: "generate table string for markdown format",
		Long:  "generate table string for markdown format",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				return errors.New("arguments not correct")
			}
			row, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}
			col, err := strconv.Atoi(args[1])
			if err != nil {
				return err
			}

			tableStr, err := generateTableStr(row, col)
			if err != nil {
				return err
			}

			_, err = fmt.Fprint(ctx.Out, tableStr)
			if err != nil {
				return err
			}

			return nil
		},
	}

	return tableCmd
}

func generateTableStr(row, col int) (string, error) {
	zap.L().Info("table param", zap.Int("row", row), zap.Int("col", col))

	if row <= 1 {
		return "", errors.New("invalid row value")
	}
	if col <= 0 {
		return "", errors.New("invalid col value")
	}

	out := []string{}

	colSplits := []string{}
	for i := 0; i <= col; i++ {
		colSplits = append(colSplits, "|")
	}
	line := strings.Join(colSplits, " ")
	baseLine := strings.Join(colSplits, " --- ")

	out = append(out, line)
	out = append(out, baseLine)
	for i := 1; i < row; i++ {
		out = append(out, line)
	}

	return strings.Join(out, "\n"), nil
}
