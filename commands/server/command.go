package server

import (
	"github.com/sirupsen/logrus"
	"github.com/finalapproachsoftware/leasing-api/logging"
	"strings"
)

type Command struct {

}

func(c *Command) Synopsis() string{
	return "Runs the leasing api server"
} 

func(c *Command) Help() string{
	helpText := `
	Usage: leasing server [options]
	`

	return strings.TrimSpace(helpText)
}


func(c *Command) Run(args []string) int {
	logging.Log().WithFields(logrus.Fields{
		"event": "start",
		"topic": "something else",
		"key": "a key",
	  }).Info("Server Starting...")

	return 0
}


