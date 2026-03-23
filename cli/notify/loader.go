package main

import (
	"github.com/urfave/cli"
)

var (
	// Plugins All the loaded plugins
	Plugins = map[string]Plugin{}
)

// Load Called from main loader
func Load(App *cli.App) ([]cli.Command, error) {
	cmds := []cli.Command{}
	log := New("plugin-loader")
	registerPlugins()
	for index, plugin := range Plugins {
		log.Debug("Setting up plugin %s", index)
		cmd, err := plugin.SetUP()
		if err != nil {
			reportError(err)
			break
		}
		cmd.Action = func(c *cli.Context) error {
			if err := plugin.Handler(c); err != nil {
				reportError(err)
				return err
			}
			return nil
		}
		cmds = append(cmds, cmd)
	}
	return cmds, nil
}

func registerPlugins() {
	LoadPlugin("gntp", gntpPlugin)
}

// LoadPlugin Função auxiliar que carrega os plugins
func LoadPlugin(name string, plugin Plugin) {
	Plugins[name] = plugin
}

// Plugin Generic plugin
type Plugin interface {
	SetUP() (cli.Command, error)
	Handler(*cli.Context) error
}
