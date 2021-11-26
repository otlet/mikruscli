package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var (
	config Config
)

func main() {
	config = Config{}
	config.init()
	app := &cli.App{
		Name:     "mikruscli",
		Usage:    "connect and manage Mikrus Web Services Instances",
		Commands: getCommands(),
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
