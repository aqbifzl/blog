package cli

import (
	"blog/config"
	"os"

	"github.com/urfave/cli/v2"
)

func Entrypoint() error {
	conf, err := config.NewConfig()
	if err != nil {
		return err
	}

	app := cli.App{
		Name: "blog",
		Commands: []*cli.Command{
			{
				Name:  "serve",
				Usage: "run the api server",
				Action: func(ctx *cli.Context) error {
					return cli_serve(conf)
				},
			},
		},
	}

	return app.Run(os.Args)
}
