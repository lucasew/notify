package loader

import (
	"github.com/lucas59356/notify/log"
	"github.com/lucas59356/notify/plugin/gntp"
	"github.com/urfave/cli"
)

var (
	// Plugins All the loaded plugins
	Plugins = map[string]Plugin{}
)

// Load Called from main loader
func Load(App *cli.App) ([]cli.Command, error) {
	cmds := []cli.Command{}
	log := logger.New("plugin-loader")
	registerPlugins()
	for index, plugin := range Plugins {
		log.Debug("Setting up plugin %s", index)
		cmd, err := plugin.SetUP()
		if err != nil {
			log.Error(err)
			break
		}
		cmd.Action = plugin.Handler
		cmds = append(cmds, cmd)
	}
	return cmds, nil
}

func registerPlugins() {
	LoadPlugin("gntp", gntp.Plugin)
}

// LoadPlugin Função auxiliar que carrega os plugins
func LoadPlugin(name string, plugin Plugin) {
	Plugins[name] = plugin
}

// Plugin Generic plugin
type Plugin interface {
	SetUP() (cli.Command, error)
	Handler(*cli.Context)
}
