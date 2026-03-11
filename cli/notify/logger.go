package main

import (
	lg "github.com/lucas59356/go-logger"
	"os"
)

const (
	debug = true
)

// New Create new logger object
func New(module string) lg.Logger {
	g := lg.NewGenerator(os.Stdout)
	if debug {
		g.SetDebugLevel(lg.LogLevelDebug1)
	}
	return g.New(module)
}