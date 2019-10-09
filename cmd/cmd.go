package cmd

import (
	"fmt"
	"github.com/jgitgud/dot-sync/lib"
)

type ArgError string

func (err ArgError) Error() string {
	return string(err)
}

/*
Example error from Dave Cheney

// PathError records an error and the operation
// and file path that caused it.
type PathError struct {
        Op   string
        Path string
        Err  error // the cause
}

func (e *PathError) Error() string

*/

// Parse command line arguments then call the corresponding command
func ParseArgs(args []string, commands map[string]Command) error {
	if len(args) == 0 {
		return ArgError("No arguments given")
	}

	if command, ok := commands[args[0]]; ok {
		return command.run(args[1:])
	}

	return ArgError("Not an argument")
}

// Struct representing a command line argument and a function to call
type Command struct {
	Name  string
	Usage string
	//Description string
	Fn func(args []string) error
}

// Execute the command, passing string args
// If there is an error with the args passed to the commmand
// print command's usage information
func (c *Command) run(args []string) error {
	err := c.Fn(args)
	if cerr, ok := err.(ArgError); ok {
		return fmt.Errorf("%v\nusage: %v %v", cerr, c.Name, c.Usage)
	}

	return err
}

func Add(args []string) error {
	if len(args) < 1 {
		return ArgError("no app name specified") // not enough args
	}

	appName := args[0]
	paths := args[1:]

	// @fix where is add.usage
	return lib.Add(appName, paths)
}

func List(args []string) error {
	if len(args) > 0 {
		return lib.ListApp(args[0])
	}

	return lib.List()
}

func Track(args []string) error {
	if len(args) < 2 {
		return ArgError("no app name or path specified")
	}

	appName := args[0]
	paths := args[1:]

	return lib.Track(appName, paths)
}

func Sync(args []string) error {
	if len(args) > 0 {
		return lib.Sync(args[0])
	}

	return lib.SyncAll()
}

func Clone(args []string) error {
	if len(args) > 0 {
		return lib.Clone(args[1])
	}

	return lib.CloneAll()
}
