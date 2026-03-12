package gntp

import (
	"fmt"
	"github.com/lucas59356/notify/log"
	"github.com/mattn/go-gntp"
	"github.com/urfave/cli"
)

var (
	// Plugin Instance
	Plugin GNTP
	// AppName Name of the program
	AppName = "notify"
)

// GNTP Plugin definition
type GNTP struct {
	client gntp.Client
}

// SetUP Sets the plugin up for work
func (g GNTP) SetUP() (cli.Command, error) {
	name := "gntp"
	description := "Envia notificações através do protocolo GNTP (Growl)"
	return cli.Command{
		HelpName:    name,
		Description: description,
		Name:        name,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "title, t",
				Value: AppName,
				Usage: "Título da notificação",
			},
			cli.StringFlag{
				Name:  "body, text, b",
				Value: "",
				Usage: "Mensagem da notificação",
			},
			cli.StringFlag{
				Name:  "passwd, password, pw",
				Value: "",
				Usage: "Senha para conectar (opcional)",
			},
			cli.StringFlag{
				Name:   "host, c",
				Value:  "localhost",
				Usage:  "Computador no qual será enviada a notificação (padrão localhost ou $GROWL_HOST)",
				EnvVar: "GROWL_HOST",
			},
			cli.IntFlag{
				Name:   "port, p",
				Value:  23053,
				Usage:  "Porta para o qual será enviada a notificação (padrão 23053 ou $GROWL_PORT)",
				EnvVar: "GROWL_PORT",
			},
			cli.BoolFlag{
				Name:  "sticky, s",
				Usage: "Manter notificação? (padrão: false)",
			},
			cli.StringFlag{
				Name:  "icon, i",
				Usage: "URL para ícone da notificação",
			},
		},
	}, nil
}

// Handler Quem recebe o sinal pra começar o trab
func (g GNTP) Handler(ctx *cli.Context) {
	log := logger.New("gntp-handler")
	log.Debug("Preparando mensagem")
	g.client.Server = fmt.Sprintf("%s:%d", ctx.String("host"), ctx.Int("port"))
	n := gntp.Notification{
		DisplayName: AppName,
		Enabled:     true,
		Event:       "default",
	}
	m := gntp.Message{
		DisplayName: n.DisplayName,
		Event:       n.Event,
		Sticky:      ctx.Bool("sticky"),
		Title:       ctx.String("title"),
		Text:        ctx.String("text"),
	}
	g.client.AppName = AppName
	if ctx.String("password") != "" {
		g.client.Password = ctx.String("password")
	}
	if ctx.String("icon") != "" {
		m.Icon = ctx.String("icon")
	}
	err := g.client.Register([]gntp.Notification{n})
	if err != nil {
		log.Error(err)
	}
	err = g.client.Notify(&m)
	if err != nil {
		log.Error(err)
	}
	log.Debug("Enviado: %v", m)
}
