package cmd

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/sixty4bee/task-cli/tasker"
)

type MarkProgressCommand struct {
	fs *flag.FlagSet
}

func NewMarkProgressCommand() *MarkProgressCommand {
	return &MarkProgressCommand{
		fs: flag.NewFlagSet("mark-in-progress", flag.ContinueOnError),
	}
}

func (c *MarkProgressCommand) Name() string {
	return c.fs.Name()
}

func (c *MarkProgressCommand) Init(args []string) error {
	return c.fs.Parse(args)
}

func (c *MarkProgressCommand) Run() error {

	id, err := strconv.ParseUint(c.fs.Arg(0), 10, 0)
	if err != nil {
		return err
	}

	err = tasker.Update(uint(id), "in-progress", "status")
	if err != nil {
		return err
	}

	fmt.Printf("task with id: %d updated successfully\n", id)

	return nil
}
