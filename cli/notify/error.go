package main

import (
	"os"

	logger "github.com/lucas59356/notify/log"
)

var errorLog = logger.New("error-reporter")

// reportError is the centralized error handling function.
// All unexpected/unhandled errors should funnel through here.
func reportError(err error) {
	if err != nil {
		errorLog.Error(err)
		os.Exit(1)
	}
}
