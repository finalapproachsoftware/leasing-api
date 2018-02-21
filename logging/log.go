package logging

import (
	"os"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Entry
var log *logrus.Logger

func init() {
	log = logrus.New()
	log.Formatter = &logrus.JSONFormatter{}
	log.Out = os.Stdout

	logger = log.WithFields(logrus.Fields{
		"app": "leasing",
	})
}

func Log() *logrus.Entry {
	return logger
}
