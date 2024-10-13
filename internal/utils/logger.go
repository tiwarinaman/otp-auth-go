package utils

import (
	"github.com/sirupsen/logrus"
	"os"
)

var log = logrus.New()

func InitLogger() {
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.InfoLevel)
}

func LogInfo(message string, fields map[string]interface{}) {
	log.WithFields(fields).Info(message)
}

func LogError(message string, err error, fields map[string]interface{}) {
	if err != nil {
		fields["error"] = err.Error()
	}
	log.WithFields(fields).Error(message)
}
