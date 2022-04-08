// Copyright (c) 2022 Vitaliy Tolmachov.

package internal

import (
	"strings"
	"unicode"

	"cloud.google.com/go/profiler"
	"github.com/tolmachov/georgian/internal/settings"
	"github.com/urfave/cli/v2"
)

func startProfiler(ctx *cli.Context) error {
	return profiler.Start(profiler.Config{
		Service: strings.Map(func(r rune) rune {
			switch true {
			case unicode.IsDigit(r):
				return r
			case unicode.IsLetter(r):
				return unicode.ToLower(r)
			case r == '-' || r == '_':
				return r
			default:
				return '-'
			}
		}, ctx.App.HelpName),
		ServiceVersion:       ctx.App.Version,
		MutexProfiling:       ctx.Bool(settings.MutexProfiling.Name),
		NoCPUProfiling:       ctx.Bool(settings.NoCPUProfiling.Name),
		NoAllocProfiling:     ctx.Bool(settings.NoAllocProfiling.Name),
		AllocForceGC:         ctx.Bool(settings.AllocForceGC.Name),
		NoHeapProfiling:      ctx.Bool(settings.NoHeapProfiling.Name),
		NoGoroutineProfiling: ctx.Bool(settings.NoGoroutineProfiling.Name),
		EnableOCTelemetry:    ctx.Bool(settings.EnableOCTelemetry.Name),
		Instance:             ctx.String(settings.InstanceID.Name),
	})
}
