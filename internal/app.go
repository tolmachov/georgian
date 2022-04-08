// Copyright (c) 2022 Vitaliy Tolmachov.

package internal

import (
	"fmt"
	"io"
	"math/rand"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/tolmachov/georgian/internal/service"
	"github.com/tolmachov/georgian/internal/settings"
	"github.com/urfave/cli/v2"
)

// Version contains semantic version number of application.
//
// Do not edit it because it will be overridden by CI.
var Version = "dev"

const serviceName = "georgian"

func init() {
	rand.Seed(time.Now().UnixNano())
	settings.ServiceName.Value = strings.ToLower(serviceName)
}

// Before is a function which must be called before any action.
//
// It initializes the logger and prints the run arguments.
func before(ctx *cli.Context) error {
	if err := initLogger(ctx); err != nil {
		return fmt.Errorf("initing logger: %w", err)
	}
	if err := startProfiler(ctx); err != nil {
		log.Warn().Err(err).Msg("failed to start profiler")
	}
	startTracer(ctx)
	greeting(ctx)
	return nil
}

// After is a function which must be called after any action.
func after(ctx *cli.Context) error {
	log.Info().Str("command", ctx.Command.Name).Msg("stopped")
	return nil
}

// New creates new instance of application.
func New(out, errOut io.Writer) *cli.App {
	app := cli.NewApp()
	app.Name = serviceName
	app.HelpName = strings.ToLower(serviceName)
	app.Version = Version
	app.Usage = "Telegram bot for transliteration and translation between Georgian and Russian."
	app.Writer = out
	app.ErrWriter = errOut
	app.Authors = []*cli.Author{
		VitaliiTolmachov, // Developer
	}
	app.Copyright = "Copyright (c) 2022 Vitaliy Tolmachov."
	app.Flags = []cli.Flag{
		settings.ProjectID,
		settings.ServiceName,
		settings.InstanceID,
		settings.LogLevel,
		settings.FractionOfTraces,
		settings.MutexProfiling,
		settings.NoCPUProfiling,
		settings.NoAllocProfiling,
		settings.NoHeapProfiling,
		settings.NoGoroutineProfiling,
		settings.AllocForceGC,
		settings.EnableOCTelemetry,
	}
	app.Before = before
	app.Action = service.Run
	app.After = after

	sortAll(app)

	return app
}
