package logger

import (
	"go.uber.org/zap"
	"log"
)

var logger *zap.Logger

func Init() {
	logg, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}

	logger = logg

	defer logger.Sync()
}

func Get() *zap.Logger {
	return logger
}
