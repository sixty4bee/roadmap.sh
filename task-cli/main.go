package main

import (
	"fmt"
	"os"

	"github.com/kobiedanquah/task-cli/cmd"
)

func main() {
	if err:= cmd.Root(os.Args[1:]); err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
}
