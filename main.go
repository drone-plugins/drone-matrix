// Copyright (c) 2020, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

// DO NOT MODIFY THIS FILE DIRECTLY

package main

import (
	"os"

	"github.com/drone-plugins/drone-matrix/plugin"
	"github.com/drone-plugins/drone-plugin-lib/errors"
	"github.com/drone-plugins/drone-plugin-lib/urfave"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
)

var version = "unknown"

func main() {
	settings := &plugin.Settings{}

	if _, err := os.Stat("/run/drone/env"); err == nil {
		_ = godotenv.Overload("/run/drone/env")
	}

	app := &cli.App{
		Name:    "drone-matrix",
		Usage:   "build notifications for matrix",
		Version: version,
		Flags:   append(settingsFlags(settings), urfave.Flags()...),
		Action:  run(settings),
	}

	if err := app.Run(os.Args); err != nil {
		errors.HandleExit(err)
	}
}

func run(settings *plugin.Settings) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		urfave.LoggingFromContext(ctx)

		plugin := plugin.New(
			*settings,
			urfave.PipelineFromContext(ctx),
			urfave.NetworkFromContext(ctx),
		)

		if err := plugin.Validate(); err != nil {
			if e, ok := err.(errors.ExitCoder); ok {
				return e
			}

			return errors.ExitMessagef("validation failed: %w", err)
		}

		if err := plugin.Execute(); err != nil {
			if e, ok := err.(errors.ExitCoder); ok {
				return e
			}

			return errors.ExitMessagef("execution failed: %w", err)
		}

		return nil
	}
}

// settingsFlags has the cli.Flags for the plugin.Settings.
func settingsFlags(settings *plugin.Settings) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "username",
			Usage:       "username for authentication",
			EnvVars:     []string{"PLUGIN_USERNAME", "MATRIX_USERNAME"},
			Destination: &settings.Username,
		},
		&cli.StringFlag{
			Name:        "password",
			Usage:       "password for authentication",
			EnvVars:     []string{"PLUGIN_PASSWORD", "MATRIX_PASSWORD"},
			Destination: &settings.Password,
		},
		&cli.StringFlag{
			Name:        "userid",
			Usage:       "userid for authentication",
			EnvVars:     []string{"PLUGIN_USERID,PLUGIN_USER_ID", "MATRIX_USERID", "MATRIX_USER_ID"},
			Destination: &settings.UserID,
		},
		&cli.StringFlag{
			Name:        "accesstoken",
			Usage:       "accesstoken for authentication",
			EnvVars:     []string{"PLUGIN_ACCESSTOKEN,PLUGIN_ACCESS_TOKEN", "MATRIX_ACCESSTOKEN", "MATRIX_ACCESS_TOKEN"},
			Destination: &settings.AccessToken,
		},
		&cli.StringFlag{
			Name:        "homeserver",
			Usage:       "matrix home server",
			EnvVars:     []string{"PLUGIN_HOMESERVER", "MATRIX_HOMESERVER"},
			Value:       "https://matrix.org",
			Destination: &settings.Homeserver,
		},
		&cli.StringFlag{
			Name:        "roomid",
			Usage:       "roomid to send messages",
			EnvVars:     []string{"PLUGIN_ROOMID", "MATRIX_ROOMID"},
			Destination: &settings.RoomID,
		},
		&cli.StringFlag{
			Name:        "template",
			Usage:       "template for the message",
			EnvVars:     []string{"PLUGIN_TEMPLATE", "MATRIX_TEMPLATE"},
			Value:       "Build {{ build.status }} [{{ repo.Owner }}/{{ repo.Name }}#{{ truncate build.commit 8 }}]({{ build.link }}) ({{ build.branch }}) by {{ build.author }}",
			Destination: &settings.Template,
		},
	}
}
