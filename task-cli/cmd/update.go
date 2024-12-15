package cmd

import (
	"errors"
	"flag"
	"fmt"
	"strconv"

	"github.com/sixty4bee/task-cli/tasker"
)

type UpdateCommand struct {
	fs *flag.FlagSet
}

func NewUpdateCommand() *UpdateCommand {
	return &UpdateCommand{
		fs: flag.NewFlagSet("update", flag.ContinueOnError),
	}
}

func (c *UpdateCommand) Name() string {
	return c.fs.Name()
}

func (c *UpdateCommand) Init(args []string) error {
	return c.fs.Parse(args)
}

func (c *UpdateCommand) Run() error {

	id, err := strconv.ParseUint(c.fs.Arg(0), 10, 0)
	if err != nil {
		return err
	}
	description := c.fs.Arg(1)
	if description == "" {
		return errors.New("task description is empty")
	}

	err = tasker.Update(uint(id), description, "description")
	if err != nil {
		return err
	}

	fmt.Printf("task with id: %d updated successfully\n", id)

	return nil
}
