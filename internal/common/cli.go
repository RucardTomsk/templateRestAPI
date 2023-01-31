package common

import (
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
)

// InitAppCli returns *cli.App with config name and path as actions.
func InitAppCli() *cli.App {
	return &cli.App{
		Name:  "SetConfigViper",
		Usage: "changing viper configuration variables",
		Action: func(cCtx *cli.Context) error {
			viper.SetConfigName(cCtx.Args().Get(0)) // config.yml
			viper.AddConfigPath(cCtx.Args().Get(1)) // ./cmd/${SERVICE}/config/
			return nil
		},
	}
}
