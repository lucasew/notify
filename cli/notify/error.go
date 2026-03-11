package main

import (
	"os"
)

// reportError is the centralized error reporting function.
// It reports unexpected errors instead of directly logging them at the call site.
func reportError(err error) {
	if err == nil {
		return
	}
	log := New("error-reporter")
	log.Error(err)
	os.Exit(1)
}