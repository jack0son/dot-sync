package main

import (
	"fmt"
	"github.com/jgitgud/dot-sync/cmd"
	"os"
)

// Example usage:
//  $ dotsync track <vim> [<path>]
//  $ dotsync sync
//  $ dotsync sync vim
//  $ dotsync add
//  $ dotsync clone vim

const usage = "dotsync <command> [<args>]"

//type CommandMap map[string]cmd.Command

func main() {
	commandsList := []cmd.Command{
		{"add", "[<app>] [<paths>]", cmd.Add},
		{"list", "[<app>]", cmd.List},
		{"track", "<app> <path> <path>...", cmd.Track},
	}

	commands := make(map[string]cmd.Command)
	for _, command := range commandsList {
		commands[command.Name] = command
	}

	err := cmd.ParseArgs(os.Args[1:], commands)
	if err != nil {
		// @fix this logic doesn't use ArgError correctly
		// but the behaviour is correct for now
		if _, ok := err.(cmd.ArgError); ok {
			fmt.Fprintf(os.Stderr, "usage: %s\n", usage)
		} else {
			fmt.Fprintf(os.Stderr, "dotsync: %s\n", err)
		}
		os.Exit(1)
	}
}

// --- Workflow notes / ideas ---
// On setup
//	- enter path to repository
// Alternative to json config:
// add from file apps/app
//	- files contains paths to files
//	- offer to enter any other paths
//	- create new dir in repository
