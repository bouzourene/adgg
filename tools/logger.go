package tools

import (
	"log"

	"go.uber.org/zap"
)

// This function provides a logger for the
// rest of the progran
func GetLogger() *zap.Logger {

	// Create new Zap logger with default prod settings
	logger, err := zap.NewProduction()

	// If cannot be create, fatal error using logger from stdlib
	if err != nil {
		log.Fatal("Cannot init logger!")
	}

	// Manage buffer automatically
	defer logger.Sync()

	// Return logger to be used in other methods
	return logger
}
