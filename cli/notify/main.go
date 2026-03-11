package main

import (
	"github.com/urfave/cli"
	"os"
)

var (
	app = cli.NewApp()
)

func main () {
	log := New("main")
	log.Debug("Iniciando...")

	app.Name = "notify"
	app.Usage = "Envia notificações para diversos destinos"
	app.Author = "lucas59356"
	app.Version = "0.1"

	app.Commands = []cli.Command{
		GNTPCmd(),
	}

	err := app.Run(os.Args)
	if err != nil {
		reportError(err)
	}
}