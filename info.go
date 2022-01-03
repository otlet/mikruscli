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
	}
}
