package cmd

import (
	"errors"
	"flag"
	"fmt"
	"time"

	"github.com/kobiedanquah/task-cli/task"
)

type AddCommand struct {
	fs          *flag.FlagSet
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
	if len(c.fs.Args()) < 1{
		return errors.New("add command requires the description of task")
	}

	task := task.Task{
		Id: 8,
		Description: c.fs.Arg(0),
		Status: "todo",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	fmt.Printf("task: %v\n", task)
	
	
	return nil
}
