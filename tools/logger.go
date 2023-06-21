package tools

import (
	"log"

	"go.uber.org/zap"
)

func GetLogger() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal("Cannot init logger!")
	}

	defer logger.Sync()

	return logger
}
