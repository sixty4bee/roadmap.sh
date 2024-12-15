package cmd

import (
	"errors"
	"flag"
	"fmt"
	"time"

	"github.com/sixty4bee/task-cli/tasker"
)

type AddCommand struct {
	fs *flag.FlagSet
}

func NewAddCommand() *AddCommand {
	cmd := &AddCommand{
		fs: flag.NewFlagSet("add", flag.ContinueOnError),
	}

	return cmd
}

func (c *AddCommand) Name() string {
	return c.fs.Name()
}

func (c *AddCommand) Init(args []string) error {
	return c.fs.Parse(args)
}

func (c *AddCommand) Run() error {
	if len(c.fs.Args()) < 1 {
		return errors.New("add command requires the description of task")
	}

	task := &tasker.Task{
		Description: c.fs.Arg(0),
		Status:      "todo",
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}

	id, err := tasker.Add(task)
	if err != nil {
		return err
	}

	fmt.Printf("Task added successfully (ID: %d)\n", id)

	return nil
}
