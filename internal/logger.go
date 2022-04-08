// Copyright (c) 2022 Vitaliy Tolmachov.

package internal

import (
	"fmt"
	"io"
	"time"

	"cloud.google.com/go/logging"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/tolmachov/georgian/internal/settings"
	"github.com/urfave/cli/v2"
)

func init() {
	zerolog.TimeFieldFormat = time.RFC3339Nano
	zerolog.ErrorStackFieldName = "stack_trace"
	zerolog.LevelFieldMarshalFunc = func(l zerolog.Level) string {
		switch l {
		case zerolog.NoLevel:
			return ""
		case zerolog.TraceLevel:
			return logging.Default.String()
		case zerolog.DebugLevel:
			return logging.Debug.String()
		case zerolog.InfoLevel:
			return logging.Info.String()
		case zerolog.WarnLevel:
			return logging.Warning.String()
		case zerolog.ErrorLevel:
			return logging.Error.String()
		case zerolog.PanicLevel:
			return logging.Critical.String()
		case zerolog.FatalLevel:
			return logging.Emergency.String()
		default:
			panic(fmt.Errorf("unknown logging level: %v", l))
		}
	}
}

func errorReportingHook(ctx *cli.Context) zerolog.HookFunc {
	return func(e *zerolog.Event, level zerolog.Level, message string) {
		serviceContext := zerolog.Dict()
		serviceContext = serviceContext.Str("service", ctx.String(settings.ServiceName.Name))
		serviceContext = serviceContext.Str("version", ctx.App.Version)
		e.Dict("service_context", serviceContext)
		e.Caller(4)
		if level == zerolog.ErrorLevel || level == zerolog.FatalLevel || level == zerolog.PanicLevel {
			e.Str("@type", "type.googleapis.com/google.devtools.clouderrorreporting.v1beta1.ReportedErrorEvent")
		}
	}
}

func createLogger(ctx *cli.Context, writer io.Writer) zerolog.Logger {
	return zerolog.New(writer).With().
		Str("service", ctx.String(settings.ServiceName.Name)).
		Str("instance", ctx.String(settings.InstanceID.Name)).
		Str("version", ctx.App.Version).
		Logger().Hook(errorReportingHook(ctx))
}

// initLogger initialises an logger.
func initLogger(ctx *cli.Context) error {

	log.Logger = createLogger(ctx, ctx.App.Writer)

	if levelStr := ctx.String(settings.LogLevel.Name); levelStr != "" {
		if level, err := zerolog.ParseLevel(levelStr); err == nil {
			zerolog.SetGlobalLevel(level)
		} else {
			log.Debug().
				Strs("available_levels", []string{
					zerolog.DebugLevel.String(),
					zerolog.InfoLevel.String(),
					zerolog.WarnLevel.String(),
					zerolog.ErrorLevel.String(),
					zerolog.FatalLevel.String(),
					zerolog.PanicLevel.String(),
					zerolog.NoLevel.String(),
					zerolog.Disabled.String(),
					zerolog.TraceLevel.String(),
				}).
				Str("specified_level", levelStr).
				Msg("specified invalid logger level")
			return fmt.Errorf("parsing logging level: %w", err)
		}
	}

	return nil
}
