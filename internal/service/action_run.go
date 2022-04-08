// Copyright (c) 2022 Vitaliy Tolmachov.

package service

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

func Run(ctx *cli.Context) error {
	log.Info().Str("log_level", zerolog.GlobalLevel().String()).Msg("initializing application")
	panic("not implemented")
}
