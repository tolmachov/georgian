// Copyright (c) 2022 Vitaliy Tolmachov.

package internal

import (
	"time"

	"contrib.go.opencensus.io/exporter/stackdriver"
	"contrib.go.opencensus.io/exporter/stackdriver/monitoredresource/gcp"
	"github.com/rs/zerolog/log"
	"github.com/tolmachov/georgian/internal/settings"
	"github.com/urfave/cli/v2"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/trace"
)

// startTracer creates and registers a OpenCensus Stackdriver Trace exporter.
func startTracer(ctx *cli.Context) {
	monitoredResource := gcp.Autodetect()
	// todo add service name, version, etc...
	log.Debug().Interface("monitoredResource", monitoredResource).Msg("identified the environment for tracing")
	exporter, err := stackdriver.NewExporter(stackdriver.Options{
		ProjectID:                ctx.String(settings.ProjectID.Name),
		BundleDelayThreshold:     time.Second,
		BundleCountThreshold:     1000,
		TraceSpansBufferMaxBytes: 32 * 1024 * 1024,
		Timeout:                  time.Second * 15,
		MonitoredResource:        monitoredResource,
		OnError: func(err error) {
			log.Error().Err(err).Msg("failed to upload the stats or tracing data")
		},
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to create and register a OpenCensus Stackdriver Trace exporter")
		return
	}
	view.RegisterExporter(exporter)
	if err := view.Register(ocgrpc.DefaultServerViews...); err != nil {
		log.Fatal().Err(err).Msg("failed to register data watcher")
	}
	view.SetReportingPeriod(time.Second) // Report stats at every second.
	trace.RegisterExporter(exporter)
	trace.ApplyConfig(
		trace.Config{
			DefaultSampler: trace.ProbabilitySampler(ctx.Float64(settings.FractionOfTraces.Name))})
}
