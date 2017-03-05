package main

import (
	"log"
	"os"

	"github.com/mitchellh/cli"
	"github.com/patricklecuyer/cr460-h2017test/cmd"
)

func main() {
	c := cli.NewCLI("cr460", "0.0.1")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"server": cmd.CR460ServerCommandFactory,
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
