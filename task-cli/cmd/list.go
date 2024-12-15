package cmd

import (
	"flag"
	"fmt"

	"github.com/sixty4bee/task-cli/tasker"
)

type ListCommand struct {
	fs *flag.FlagSet
}

func NewListCommand() *ListCommand {
	return &ListCommand{
		fs: flag.NewFlagSet("list", flag.ContinueOnError),
	}
}

func (c *ListCommand) Name() string {
	return c.fs.Name()
}

func (c *ListCommand) Init(args []string) error {
	return c.fs.Parse(args)
}


func (c *ListCommand) Run() error {

	tasks, err :=tasker.GetAll()
	if err != nil {
		return err
	}


	for _, task := range tasks {
		fmt.Printf("id: %d\tdescription: %s\tstatus: %s\n", task.Id, task.Description, task.Status)
	}

	return nil
}
