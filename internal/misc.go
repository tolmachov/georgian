// Copyright (c) 2022 Vitaliy Tolmachov.

package internal

import (
	"os"
	"sort"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

func greeting(ctx *cli.Context) {
	if pid := os.Getpid(); pid != 1 {
		log.Warn().Int("PID", pid).Msg("PID is not 1 - K8S liveness probes will not work correctly")
	}
	event := log.Info().Str("command", ctx.Command.FullName())
	if zerolog.GlobalLevel() == zerolog.DebugLevel {
		if flagNames := ctx.FlagNames(); len(flagNames) > 0 {
			dict := zerolog.Dict()
			for _, flagName := range ctx.FlagNames() {
				dict = dict.Interface(flagName, ctx.Value(flagName))
			}
			event = event.Dict("arguments", dict)
		}
		if environ := os.Environ(); len(environ) > 0 {
			dict := zerolog.Dict()
			for _, envVar := range environ {
				if s := strings.Split(envVar, "="); len(s) == 2 {
					dict = dict.Str(s[0], s[1])
				}
			}
			event.Dict("environment", dict)
		}
	}
	event.Msg("starting")
}

func sortSubCommands(commands []*cli.Command) {
	for _, subCommand := range commands {
		sort.Sort(cli.FlagsByName(subCommand.Flags))
		sort.Sort(cli.CommandsByName(subCommand.Subcommands))
		sortSubCommands(subCommand.Subcommands)
	}
}

// sortAll sorts the authors, flags and commands of the application
func sortAll(app *cli.App) {
	sort.Slice(app.Authors, func(i, j int) bool {
		return app.Authors[i].String() < app.Authors[j].String()
	})
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	sortSubCommands(app.Commands)
}
