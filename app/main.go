// Copyright (c) 2022 Vitaliy Tolmachov.

package main

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/tolmachov/georgian/internal"
)

func main() {
	if err := internal.New(os.Stdout, os.Stderr).Run(os.Args); err != nil {
		log.Fatal().Msg(err.Error())
	}
}
