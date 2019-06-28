package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/taxio/mds-wf/cmd/mds/cmd"
)

func main() {
	// for Alfred Workflow query format
	osArgs := os.Args[1:]
	query := strings.Join(osArgs, " ")
	splited := strings.Split(query, " ")
	var args []string
	for _, s := range splited {
		if len(s) == 0 {
			continue
		}
		args = append(args, s)
	}

	cmdCtx := &cmd.CommandContext{
		Out: os.Stdout,
	}
	rootCmd := cmd.NewRootCmd(args, cmdCtx)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}
