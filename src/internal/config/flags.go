package config

import "github.com/urfave/cli/v2"

var Flags = []cli.Flag{
	&cli.StringFlag{
		Name:        "app_environment",
		Destination: &App.Environment,
		EnvVars:     []string{"APP_ENV"},
		Value:       Development,
	},
	&cli.StringFlag{
		Name:        "app_address",
		Destination: &App.Address,
		EnvVars:     []string{"APP_ADDRESS"},
		Value:       ":8080",
	},
	&cli.StringFlag{
		Name:        "app_log_level",
		Destination: &App.LogLevel,
		EnvVars:     []string{"APP_LOG_LEVEL"},
		Value:       "debug",
	},
	&cli.StringFlag{
		Name:        "app_jwt_secret",
		Destination: &App.JWTSecret,
		EnvVars:     []string{"APP_JWT_SECRET"},
		Value:       "20994458adf248e6a2e2034235e3a0f4",
	},
	&cli.DurationFlag{
		Name:        "app_jwt_exp",
		Destination: &App.JWTExp,
		EnvVars:     []string{"APP_JWT_EXP"},
		Value:       1 * Year,
	},
	&cli.BoolFlag{
		Name:        "app_log_pretty",
		Destination: &App.LogPrettyPrint,
		EnvVars:     []string{"APP_LOG_PRETTY"},
		Value:       false,
	},
	&cli.StringFlag{
		Name:        "app_public_folder",
		Destination: &App.PublicFolder,
		EnvVars:     []string{"APP_PUBLIC_FOLDER"},
		Value:       "/home/ppcamp/Desktop/me/go-websocket/public/index.html",
	},
}
