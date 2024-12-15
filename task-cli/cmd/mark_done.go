package cmd

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/sixty4bee/task-cli/tasker"
)

type MarkDoneCommand struct {
	fs *flag.FlagSet
}

func NewMarkDoneCommand() *MarkDoneCommand {
	return &MarkDoneCommand{
		fs: flag.NewFlagSet("mark-done", flag.ContinueOnError),
	}
}

func (c *MarkDoneCommand) Name() string {
	return c.fs.Name()
}

func (c *MarkDoneCommand) Init(args []string) error {
	return c.fs.Parse(args)
}

func (c *MarkDoneCommand) Run() error {

	id, err := strconv.ParseUint(c.fs.Arg(0), 10, 0)
	if err != nil {
		return err
	}

	err = tasker.Update(uint(id), "done", "status")
	if err != nil {
		return err
	}

	fmt.Printf("task with id: %d updated successfully\n", id)

	return nil
}
