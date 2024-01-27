package logging

import (
	"github.com/sirupsen/logrus"
	"os"
)

type AppLogger struct {
}

func NewAppLogger() ILogger {
	return &AppLogger{}
}

func (a *AppLogger) LogConsole() *logrus.Logger {
	var log = logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.DebugLevel)

	return log
}

func (a *AppLogger) LogFile() *logrus.Logger {
	//TODO implement me
	panic("implement me")
}
