package logging

import "github.com/sirupsen/logrus"

type ILogger interface {
	LogConsole() *logrus.Logger
	LogFile() *logrus.Logger
}
