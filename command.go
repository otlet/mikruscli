package main

import (
	"github.com/urfave/cli/v2"
)

func getCommands() []*cli.Command {
	return []*cli.Command{
		getContextCommands(),
		getInfoCommands(),
		// getRestartCommands(),
		// getLogsCommands(),
		getAmfetaminaCommands(),
		// getDbCommands(),
		// getExecCommands(),
		getStatsCommands(),
	}
}
