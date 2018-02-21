package main

import (
	"github.com/finalapproachsoftware/leasing-api/commands/server"
	"log"
	"os"
	"github.com/mitchellh/cli"
	
)

func main() {
	c := cli.NewCLI("leasing", "0.1.0")
	c.Args = os.Args[1:]
	
	c.Commands = map[string]cli.CommandFactory{
		"server": func() (cli.Command, error){
			return &server.Command{}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}