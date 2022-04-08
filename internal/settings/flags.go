// Copyright (c) 2022 Vitaliy Tolmachov.

package settings

import (
	"cloud.google.com/go/logging"
	"github.com/urfave/cli/v2"
)

// ProjectID is a Google cloud platform project ID.
var ProjectID = &cli.StringFlag{
	Name:    "project-id",
	Usage:   "google cloud platform project id",
	EnvVars: []string{"PROJECT_ID"},
}

// ServiceName specifies own service name.
var ServiceName = &cli.StringFlag{
	Name:    "service-name",
	Usage:   "own microservice name",
	EnvVars: []string{"SERVICE_NAME"},
}

// InstanceID specifies instance ID.
var InstanceID = &cli.StringFlag{
	Name:    "instance",
	Usage:   "instance ID, for example: dev, stage or production",
	EnvVars: []string{"INSTANCE_ID"},
	Value:   "dev",
}

// LogLevel specifies logging level.
var LogLevel = &cli.StringFlag{
	Name:    "log-level",
	Usage:   "logging level",
	EnvVars: []string{"LOG_LEVEL"},
	Value:   logging.Info.String(),
}

// FractionOfTraces specifies probability whether a trace should be sampled and exported.
var FractionOfTraces = &cli.Float64Flag{
	Name:    "fraction-of-traces",
	Usage:   "Specifies probability whether a trace should be sampled and exported",
	EnvVars: []string{"FRACTION_OF_TRACES"},
	Value:   0.1,
}

// region Profiling
var (
	// MutexProfiling enables mutex profiling. It defaults to false.
	MutexProfiling = &cli.BoolFlag{
		Name:    "mutex-profiling",
		Usage:   "MutexProfiling enables mutex profiling",
		EnvVars: []string{"MUTEX_PROFILING"},
		Value:   false,
	}

	// NoCPUProfiling collecting the CPU profiles is disabled.
	NoCPUProfiling = &cli.BoolFlag{
		Name:    "no-cpu-profiling",
		Usage:   "When true, collecting the CPU profiles is disabled",
		EnvVars: []string{"NO_CPU_PROFILING"},
		Value:   false,
	}

	// NoAllocProfiling collecting the allocation profiles is disabled.
	NoAllocProfiling = &cli.BoolFlag{
		Name:    "no-alloc-profiling",
		Usage:   "When true, collecting the allocation profiles is disabled",
		EnvVars: []string{"NO_ALLOC_PROFILING"},
		Value:   false,
	}

	// NoHeapProfiling collecting the heap profiles is disabled.
	NoHeapProfiling = &cli.BoolFlag{
		Name:    "no-heap-profiling",
		Usage:   "When true, collecting the heap profiles is disabled.",
		EnvVars: []string{"NO_HEAP_PROFILING"},
		Value:   false,
	}

	// NoGoroutineProfiling collecting the goroutine profiles is disabled.
	NoGoroutineProfiling = &cli.BoolFlag{
		Name:    "no-goroutine-profiling",
		Usage:   "When true, collecting the goroutine profiles is disabled.",
		EnvVars: []string{"NO_GOROUTINE_PROFILING"},
		Value:   false,
	}

	// AllocForceGC forces garbage collection before the collection of each heap
	// profile collected to produce the allocation profile. This increases the
	// accuracy of allocation profiling. It defaults to false.
	AllocForceGC = &cli.BoolFlag{
		Name: "alloc-force-gc",
		Usage: "AllocForceGC forces garbage collection before the collection " +
			"of each heap profile collected to produce the allocation profile. " +
			"This increases the accuracy of allocation profiling",
		EnvVars: []string{"ALLOC_FORCE_GC"},
		Value:   false,
	}

	// EnableOCTelemetry sends all telemetries via OpenCensus exporter, which
	// can be viewed in Cloud Trace and Cloud Monitoring.
	EnableOCTelemetry = &cli.BoolFlag{
		Name: "enable-oc-telemetry",
		Usage: "When true, the agent sends all telemetries via OpenCensus exporter, " +
			"which can be viewed in Cloud Trace and Cloud Monitoring.",
		EnvVars: []string{"ENABLE_OC_TELEMETRY"},
		Value:   false,
	}
)

// endregion
