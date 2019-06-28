package cmd

import "io"

type CommandContext struct {
	Out io.Writer
}
