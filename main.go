package main

import (
	"log"
	"os"
	"text/tabwriter"

	"github.com/urfave/cli/v2"
)

var (
	version string = "0.0.1"
	config  Config
	w       *tabwriter.Writer
)

func main() {
	config = Config{}
	config.init()
	w = tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	app := &cli.App{
		Name:     "mikruscli",
		Version:  version,
		Usage:    "connect and manage Mikrus Web Services Instances",
		Commands: getCommands(),
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
