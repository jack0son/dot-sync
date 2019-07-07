package cmd

import (
	"fmt"
	"github.com/jgitgud/dot-sync/lib"
)

type ArgError string

func (err ArgError) Error() string {
	return fmt.Sprintf("%v", err)
}

func parseArgs(args []string, commands map[string]Command) error {
	if len(args) == 0 {
		return ArgError("No arguments given")
	}

	if command, ok := commands[args[0]]; ok {
		return command.run(args[1:])
	}

	return ArgError("Not an argument")
}

type Command struct {
	name  string
	usage string
	fn    func(args []string) error
}

// Execute the command, passing string args
// If there is an error with the args passed to the commmand
// print command's usage information
func (c *Command) run(args []string) error {
	err := c.run(args)
	if cerr, ok := err.(ArgError); ok {
		return fmt.Errorf("%v: %v\n%v", c.name, cerr, c.usage)
	}

	return err
}

func Add(args []string) error {
	if len(args) < 1 {
		return ArgError("no app name specified") // not enough args
	}

	appName := args[1]
	paths := args[1:]

	// @fix where is add.usage
	return lib.Add(appName, paths)
}

func Track(args []string) error {
	if len(args) < 2 {
		return errors.New("")
	}

	appName := args[1]
	paths := args[1:]

	return lib.Track(appName, paths)
}

func Sync(args []string) error {
	if len(args) > 0 {
		return lib.Sync(args[1])
	}

	return lib.SyncAll()
}

func Clone(args []string) error {
	if len(args) > 0 {
		return lib.Clone(args[1])
	}

	return lib.CloneAll()
}
