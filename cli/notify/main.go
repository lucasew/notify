package main

import (
	"github.com/urfave/cli"
	"github.com/lucas59356/notify/log"
	"os"
	"github.com/lucas59356/notify/plugin"
)

var (
	app = cli.NewApp()
)

func main () {
	appLogger := logger.New("main")
	appLogger.Debug("Iniciando...")
	app.Name = "notify"
	app.Usage = "Envia notificações para diversos destinos"
	app.Author = "lucas59356"
	app.Version = "0.1"
	cmds, err := loader.Load(app) // Carrega os módulos, junto com seus comandos
	if err != nil {
		appLogger.Error(err)
	}
	for _, cmd := range(cmds) { // Organiza os comandos do loader junto com os que já tem
		app.Commands = append(app.Commands, cmd)
	}
	app.Run(os.Args)
}