package flags

import "github.com/urfave/cli/v2"

func OutPutFlags() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:  "json",
			Usage: "Print output as JSON",
		},
		&cli.BoolFlag{
			Name:  "table",
			Usage: "Print output in table",
		},
	}
}
