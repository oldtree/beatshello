package main

import (
	"os"

	"github.com/oldtree/beatshello/cmd"
)

//go build -ldflags=s and go build -ldflags=-w both produce working binaries, which suggests the problem is DWARF-related, as in #23046.
//pip install --upgrade virtualenv
func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
