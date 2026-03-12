package loader

import (
	"github.com/urfave/cli"
	"github.com/lucas59356/notify/plugin/gntp"
	"github.com/lucas59356/notify/log"
)

var (
	// Providers All the loaded providers
	Providers = map[string]Provider{}
)

// Load Called from main loader
func Load(App *cli.App) ([]cli.Command, error) {
	cmds := []cli.Command{}
	loaderLogger := logger.New("plugin-loader")
	registerProviders()
	for index, provider := range Providers {
		loaderLogger.Debug("Setting up provider %s", index)
		cmd, err := provider.SetUP()
		if err != nil {
			loaderLogger.Error(err)
			break
		}
		cmd.Action = provider.Handler
		cmds = append(cmds, cmd)
	}	
	return cmds, nil
}

func registerProviders() {
	LoadProvider("gntp", gntp.Plugin)
}

// LoadProvider Função auxiliar que carrega os providers
func LoadProvider(name string, provider Provider) {
	Providers[name] = provider
}

// Provider defines the interface for notification providers.
// As described by Fowler and standard Go conventions, interfaces
// providing services are typically named with an "-er" suffix or a descriptive noun.
type Provider interface {
	SetUP()(cli.Command, error)
	Handler(*cli.Context)
}