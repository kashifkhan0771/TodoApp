package config

import (
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func log() *logger.Entry {
	level, err := logger.ParseLevel(viper.GetString(LogLevel))
	if err != nil {
		logger.SetLevel(logger.DebugLevel)
	}
	logger.SetLevel(level)

	logger.SetFormatter(&logger.TextFormatter{
		FullTimestamp: true,
	})

	return logger.WithFields(logger.Fields{
		"package": "config",
	})
}
