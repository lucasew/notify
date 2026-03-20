package logger

import (
	lg "github.com/lucas59356/go-logger"
	"os"
)

// New Create new logger object
func New(module string) lg.Logger {
	g := lg.NewGenerator(os.Stdout)
	if os.Getenv("NOTIFY_DEBUG") == "1" || os.Getenv("NOTIFY_DEBUG") == "true" {
		g.SetDebugLevel(lg.LogLevelDebug1)
	}
	return g.New(module)
}
