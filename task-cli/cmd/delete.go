package cmd

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/sixty4bee/task-cli/tasker"
)

type DeleteCommand struct {
	fs *flag.FlagSet
}

func NewDeleteCommand() *DeleteCommand {
	return &DeleteCommand{
		fs: flag.NewFlagSet("delete", flag.ContinueOnError),
	}
}

func (c *DeleteCommand) Name() string {
	return c.fs.Name()
}

func (c *DeleteCommand) Init(args []string) error {
	return c.fs.Parse(args)
}

func (c *DeleteCommand) Run() error {

	id, err := strconv.ParseUint(c.fs.Arg(0), 10, 0)
	if err != nil {
		return err
	}

	err = tasker.Delete(uint(id))
	if err != nil {
		return err
	}

	fmt.Println("task successfully deleted")

	return nil
}
