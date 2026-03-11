package main

import (
	"fmt"
	"github.com/urfave/cli"
	"github.com/mattn/go-gntp"
)

const (
	AppName = "notify"
	DefaultGNTPHost = "localhost"
	DefaultGNTPPort = 23053
	GNTPPluginName = "gntp"
	GNTPPluginDesc = "Envia notificações através do protocolo GNTP (Growl)"
)

// GNTPCmd Setups the gntp command
func GNTPCmd() cli.Command {
	return cli.Command{
		HelpName: GNTPPluginName,
		Description: GNTPPluginDesc,
		Name: GNTPPluginName,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name: "title, t",
				Value: AppName,
				Usage: "Título da notificação",
			},
			cli.StringFlag{
				Name: "body, text, b",
				Value: "",
				Usage: "Mensagem da notificação",
			},
			cli.StringFlag{
				Name: "passwd, password, pw",
				Value: "",
				Usage: "Senha para conectar (opcional)",
			},
			cli.StringFlag{
				Name: "host, c",
				Value: DefaultGNTPHost,
				Usage: "Computador no qual será enviada a notificação (padrão localhost ou $GROWL_HOST)",
				EnvVar: "GROWL_HOST",
			},
			cli.IntFlag{
				Name: "port, p",
				Value: DefaultGNTPPort,
				Usage: "Porta para o qual será enviada a notificação (padrão 23053 ou $GROWL_PORT)",
				EnvVar: "GROWL_PORT",
			},
			cli.BoolFlag{
				Name: "sticky, s",
				Usage: "Manter notificação? (padrão: false)",
			},
			cli.StringFlag{
				Name: "icon, i",
				Usage: "URL para ícone da notificação",
			},
		},
		Action: func(ctx *cli.Context) {
			log := New("gntp-handler")
			log.Debug("Preparando mensagem")

			client := gntp.Client{
				Server: fmt.Sprintf("%s:%d", ctx.String("host"), ctx.Int("port")),
				AppName: AppName,
			}

			n := gntp.Notification{
				DisplayName: AppName,
				Enabled: true,
				Event: "default",
			}

			m := gntp.Message{
				DisplayName: n.DisplayName,
				Event: n.Event,
				Sticky: ctx.Bool("sticky"),
				Title: ctx.String("title"),
				Text: ctx.String("text"),
			}

			if ctx.String("password") != "" {
				client.Password = ctx.String("password")
			}
			if ctx.String("icon") != "" {
				m.Icon = ctx.String("icon")
			}

			err := client.Register([]gntp.Notification{n})
			if err != nil {
				reportError(err)
			}

			err = client.Notify(&m)
			if err != nil {
				reportError(err)
			}

			log.Debug("Enviado: %v", m)
		},
	}
}