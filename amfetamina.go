package main

import (
	"encoding/json"
	"fmt"

	"github.com/otlet/mikruscli/structs"
	"github.com/urfave/cli/v2"
)

func getAmfetaminaCommands() *cli.Command {
	return &cli.Command{
		Name:    "amfetamina",
		Aliases: []string{"a"},
		Usage:   "Run amfetamina!",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "yes",
				Aliases: []string{"y"},
				Usage:   "Pass yes :P",
			},
		},
		Action: func(c *cli.Context) error {
			if !c.Bool("yes") {
				for ok := true; ok; {
					var response string
					fmt.Printf("Are you sure? (Y/n) ")
					fmt.Scanln(&response)
					if response == "Y" || response == "n" || response == "yes" || response == "no" {
						if response == "Y" || response == "yes" {
							break
						} else {
							return nil
						}
					}
				}
			}

			data, err := makeMikrusApiCall("amfetamina")
			if err != nil {
				return err
			}

			var response structs.Task
			json.Unmarshal(data, &response)
			if response.Error != "" {
				fmt.Println(response.Error)
			} else {
				fmt.Println("Uruchomiono Amfetaminę!")
				fmt.Println("TaskID:", response.TaskID)
			}
			return nil
		},
	}
}
