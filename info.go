package main

import (
	"encoding/json"
	"fmt"

	"github.com/otlet/mikruscli/flags"
	"github.com/otlet/mikruscli/structs"
	"github.com/urfave/cli/v2"
)

func getInfoCommands() *cli.Command {
	return &cli.Command{
		Name:    "info",
		Aliases: []string{"i"},
		Usage:   "Get information about Mikrus Web Services server",
		Flags:   flags.OutPutFlags(),
		Action: func(c *cli.Context) error {
			data, err := makeMikrusApiCall("info")
			if err != nil {
				return err
			}
			if c.Bool("json") {
				fmt.Fprintln(w, string(data))
				return nil
			} else if c.Bool("table") {
				fmt.Println("Not implemented yet.")
				return nil
			}

			d := structs.Info{}
			json.Unmarshal(data, &d)
			fmt.Println("Server_ID:", d.Server_Id)
			fmt.Println("Server_Name:", d.Server_Name)
			fmt.Println("Expires:", d.Expires)
			fmt.Println("Expires_Cytrus:", d.Expires_Cytrus)
			fmt.Println("Expires_Storage:", d.Expires_Storage)
			fmt.Println("Param_Ram:", fmt.Sprintf("%dM", d.Param_Ram))
			fmt.Println("Param_Disk:", fmt.Sprintf("%dG", d.Param_Disk))
			fmt.Println("Lastlog_Panel:", d.Lastlog_Panel)
			return nil
		},
		Subcommands: []*cli.Command{
			{
				Name:    "stats",
				Aliases: []string{"s"},
				Usage:   "Get statistics about Mikrus Web Services server",
				Flags:   flags.OutPutFlags(),
				Action: func(c *cli.Context) error {
					data, err := makeMikrusApiCall("stats")
					if err != nil {
						return err
					}

					if c.Bool("json") {
						fmt.Fprintln(w, string(data))
						return nil
					} else if c.Bool("table") {
						fmt.Println("Not implemented yet.")
						return nil
					}

					d := structs.Stats{}
					json.Unmarshal(data, &d)

					fmt.Println("[free]")
					fmt.Println(d.Free)

					fmt.Println("[df]")
					fmt.Println(d.Df)

					fmt.Println("[uptime]")
					fmt.Println(d.Uptime)

					fmt.Println("[ps]")
					fmt.Println(d.Ps)

					return nil
				},
			},
			{
				Name:    "logs",
				Aliases: []string{"l"},
				Usage:   "Get logs from Mikrus Web Services server",
				Action: func(c *cli.Context) error {
					data, err := makeMikrusApiCall("logs")
					if err != nil {
						return err
					}

					if c.Bool("json") {
						fmt.Fprintln(w, string(data))
						return nil
					} else if c.Bool("table") {
						fmt.Println("Not implemented yet.")
						return nil
					}

					d := []structs.Logs{}
					json.Unmarshal(data, &d)

					for _, v := range d {
						fmt.Println("[ TASK:", v.ID, "]")
						fmt.Println("Server_ID:", v.ServerID)
						fmt.Println("Task:", v.Task)
						fmt.Println("When created: ", v.WhenCreated)
						fmt.Println("When done: ", v.WhenDone)
						fmt.Println("Output: ", v.Output)
					}

					return nil
				},
			},
			{
				Name:    "ports",
				Aliases: []string{"p"},
				Usage:   "Get ports from Mikrus Web Services server",
				Action: func(c *cli.Context) error {
					data, err := makeMikrusApiCall("porty")
					if err != nil {
						return err
					}

					var d []int
					json.Unmarshal(data, &d)

					for _, v := range d {
						fmt.Println(v)
					}

					return nil

				},
			},
			{
				Name:  "db",
				Usage: "Get database information from Mikrus Web Services server",
				Action: func(c *cli.Context) error {
					data, err := makeMikrusApiCall("db")
					if err != nil {
						return err
					}

					d := structs.Db{}
					json.Unmarshal(data, &d)

					fmt.Println("[MySQL]")
					fmt.Println(d.Mysql)

					fmt.Println("[PostgreSQL]")
					fmt.Println(d.Postgresql)

					fmt.Println("[MongoDB]")
					fmt.Println(d.Mongo)

					return nil
				},
			},
		},
	}
}
