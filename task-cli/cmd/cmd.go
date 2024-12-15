package cmd

import (
	"errors"
	"fmt"
)

type Runner interface {
	Init([]string) error
	Run() error
	Name() string
}

func Root(args []string) error {

	if len(args) < 1 {
		return errors.New("you must pass a sub-command")
	}

	commands := []Runner{
		NewAddCommand(),
		NewListCommand(),
		NewUpdateCommand(),
		NewMarkDoneCommand(),
		NewMarkProgressCommand(),
		NewDeleteCommand(),
	}

	subcommand := args[0]

	for _, cmd := range commands {
		if cmd.Name() == subcommand {
			cmd.Init(args[1:])
			return cmd.Run()
		}
	}

	return fmt.Errorf("unknown subcommand: %s", subcommand)
}
