package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	query := os.Args[1]
	splited := strings.Split(query, " ")
	args := []string{}
	for _, s := range splited {
		if len(s) == 0 {
			continue
		}
		args = append(args, s)
	}
	if len(args) != 2 {
		fmt.Fprintln(os.Stderr, "arguments not correct")
		os.Exit(2)
	}
	fmt.Fprintf(os.Stdout, "[%s](%s)", args[0], args[1])
	os.Exit(0)
}
