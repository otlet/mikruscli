package main

import (
	"fmt"
	"log"

	"github.com/urfave/cli/v2"
)

func getContextCommands() *cli.Command {
	return &cli.Command{
		Name:    "ctx",
		Aliases: []string{"context"},
		Usage:   "Manage API Context. Context is a information about server",
		Action: func(c *cli.Context) error {
			log.Println("Current context:", config.Ctx[config.CurrentCtx].Name)
			log.Println("Try 'mikruscli ctx help' for more information")
			return nil
		},
		Subcommands: []*cli.Command{
			{
				Name:  "list",
				Usage: "List available context",
				Action: func(c *cli.Context) error {
					if len(config.Ctx) == 0 {
						log.Println("No context available")
						return nil
					}

					log.Println("Available context:")

					for _, ctx := range config.Ctx {
						log.Printf("* %s (%s)\n", ctx.Name, ctx.ServerName)
					}
					return nil
				},
			},
			{
				Name:  "add",
				Usage: "Add new context",
				Action: func(c *cli.Context) error {
					var (
						name   string
						server string
						token  string
					)

					log.Println("Add new context")
					log.Println("Enter name:")
					fmt.Scanln(&name)
					log.Println("Enter server id (like e123):")
					fmt.Scanln(&server)
					log.Println("Enter token:")
					fmt.Scanln(&token)

					ctx := Context{
						Name:       name,
						ServerName: server,
						Token:      token,
					}

					config.Ctx = append(config.Ctx, ctx)
					config.Save()

					return nil
				},
			},
			{
				Name:  "switch",
				Usage: "Switch to another context",
				Action: func(c *cli.Context) error {
					if len(config.Ctx) == 0 {
						log.Println("No context available")
					}

					log.Println("Available context:")
					for i, ctx := range config.Ctx {
						log.Printf("* %d. %s (%s)\n", i+1, ctx.Name, ctx.ServerName)
					}

					var id int
					log.Println("Enter id of context:")
					fmt.Scanln(&id)

					if id > len(config.Ctx) {
						log.Println("Wrong id")
						return nil
					}

					config.CurrentCtx = id - 1
					config.Save()

					return nil
				},
			},
			{
				Name:  "delete",
				Usage: "Delete context",
				Action: func(c *cli.Context) error {
					if len(config.Ctx) == 0 {
						log.Println("No context available")
						return nil
					}

					for key, value := range config.Ctx {
						log.Printf("%v %v\n", key+1, value.Name)
					}

					log.Println("Enter context id to delete:")
					var id int
					fmt.Scanln(&id)

					if id > len(config.Ctx) {
						log.Println("Context id is out of range")
					}

					config.Ctx = append(config.Ctx[:id-1], config.Ctx[id:]...)
					config.Save()

					return nil
				},
			},
		},
	}
}
