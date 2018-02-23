package server

import (
	"strings"

	"github.com/finalapproachsoftware/leasing-api/logging"
	"github.com/sirupsen/logrus"
)

type Command struct {
}

func (c *Command) Synopsis() string {
	return "Runs the leasing api server"
}

func (c *Command) Help() string {
	helpText := `
	Usage: leasing server [options]
	`

	return strings.TrimSpace(helpText)
}

func (c *Command) Run(args []string) int {
	logging.Log().WithFields(logrus.Fields{
		"event": "start",
		"topic": "something else",
		"key":   "a key",
	}).Info("Server Starting...")

	server, err := newServer()
	if err != nil {
		logging.Log().Error("Could not configure server")
	}

	server.start()

	return 0
}
