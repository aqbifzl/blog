package main

import (
	"blog/cli"
	"log"
)

func main() {
	if err := cli.Entrypoint(); err != nil {
		log.Fatal(err)
	}
}
