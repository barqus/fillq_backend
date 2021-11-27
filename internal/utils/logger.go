package utils

import (
	"github.com/sirupsen/logrus"
)

func NewLogger() *logrus.Logger {
	logger := logrus.New()
	logger.Formatter = &logrus.TextFormatter{}
	logger.SetLevel(logrus.DebugLevel)
	return logger
}
