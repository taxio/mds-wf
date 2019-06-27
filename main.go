package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 2 {
		fmt.Fprintln(os.Stderr, "arguments not correct")
		os.Exit(1)
	}
	link := fmt.Sprintf("[%s](%s)", args[0], args[1])
	fmt.Fprintln(os.Stdout, link)
	os.Exit(0)
}
