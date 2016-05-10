package utils

import (
	"github.com/Sirupsen/logrus"
	"github.com/Sirupsen/logrus/formatters/logstash"
	"os"
	"time"
)

const (
	// LogStashFormatter is constant used to format logs as logstash format
	LogStashFormatter = "logstash"
	// TextFormatter is constant used to format logs as simple text format
	TextFormatter = "text"
)

// InitLog initializes the logrus logger
func InitLog(logLevel, formatter string) error {

	switch formatter {
	case LogStashFormatter:
		logrus.SetFormatter(&logstash.LogstashFormatter{
			TimestampFormat: time.RFC3339,
		})
	default:
		logrus.SetFormatter(&logrus.TextFormatter{
			ForceColors:   true,
			FullTimestamp: true,
		})
	}

	logrus.SetOutput(os.Stdout)

	level, err := logrus.ParseLevel(logLevel)

	if err != nil {
		logrus.SetLevel(logrus.DebugLevel)
		return err
	}

	logrus.SetLevel(level)
	return nil
}
