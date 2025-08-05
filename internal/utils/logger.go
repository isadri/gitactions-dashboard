package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

func GetLogger() *logrus.Logger {
	var log = logrus.New()

	log.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	log.Out = os.Stdout
	return log
}
