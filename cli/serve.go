package cli

import (
	"fmt"
	"blog/config"
	"blog/handlers"
)

func cli_serve(conf *config.AppConfig) error {
	mux := handlers.SetupServer(conf)

	fmt.Println("Serving on: http://" + conf.Bind.AsAddress())
	return mux.Run(conf.Bind.AsAddress())
}
