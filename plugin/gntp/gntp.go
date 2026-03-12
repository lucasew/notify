package gntp

import (
	"fmt"
	"github.com/urfave/cli"
	"github.com/mattn/go-gntp"
	"github.com/lucas59356/notify/log"
)
const (
	// PluginName Default plugin identifier
	PluginName = "gntp"
	// DefaultPort Default GNTP server port
	DefaultPort = 23053
	// DefaultHost Default GNTP server host
	DefaultHost = "localhost"
)

var (
	// Plugin Instance
	Plugin GNTP
	// DefaultAppName Name of the program
	DefaultAppName = "notify"
)

// GNTP Plugin definition
type GNTP struct {
	client gntp.Client
}

// SetUP Sets the plugin up for work
func (g GNTP) SetUP() (cli.Command, error) {
	description := "Envia notificações através do protocolo GNTP (Growl)"
	return cli.Command{
		HelpName: PluginName,
		Description: description,
		Name: PluginName,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name: "title, t",
				Value: DefaultAppName,
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
				Value: DefaultHost,
				Usage: "Computador no qual será enviada a notificação (padrão localhost ou $GROWL_HOST)",
				EnvVar: "GROWL_HOST",
			},
			cli.IntFlag{
				Name: "port, p",
				Value: DefaultPort,
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
	}, nil
}

// Handler Quem recebe o sinal pra começar o trab
func (g GNTP) Handler(ctx *cli.Context) {
	log := logger.New("gntp-handler")
	log.Debug("Preparando mensagem")
	g.client.Server = fmt.Sprintf("%s:%d", ctx.String("host"), ctx.Int("port"))
	n := gntp.Notification{
		DisplayName: DefaultAppName,
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
	g.client.AppName = DefaultAppName
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

