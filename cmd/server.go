package cmd

import (
	"github.com/mitchellh/cli"
	"github.com/patricklecuyer/cr460-h2017test/config"
	"github.com/patricklecuyer/cr460-h2017test/router"
)

// CR460ServerCommand Server API Command
type CR460ServerCommand struct{}

// Run runs the command
func (c CR460ServerCommand) Run(args []string) int {
	config.LoadConfig()
	r := router.Init()
	r.Run(":" + config.AppConfig.Port)
	return 0
}

// Help display helps
func (c CR460ServerCommand) Help() string {
	return "Starts the Planifio API server"
}

// Synopsis display helps
func (c CR460ServerCommand) Synopsis() string {
	return "Starts the Planifio API server"
}

//CR460ServerCommandFactory Command Factory for Planifio API Server
func CR460ServerCommandFactory() (cli.Command, error) {
	return CR460ServerCommand{}, nil
}
