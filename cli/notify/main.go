package main

import (
	"github.com/urfave/cli"
	"os"
)

var (
	app = cli.NewApp()
)

func main() {
	log := New("main")
	log.Debug("Iniciando...")
	app.Name = "notify"
	app.Usage = "Envia notificações para diversos destinos"
	app.Author = "lucas59356"
	app.Version = "0.1"
	cmds, err := Load(app) // Carrega os módulos, junto com seus comandos
	if err != nil {
		reportError(err)
	}
	for _, cmd := range cmds { // Organiza os comandos do loader junto com os que já tem
		app.Commands = append(app.Commands, cmd)
	}
	app.Run(os.Args)
}
