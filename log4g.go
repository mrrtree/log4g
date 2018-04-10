package log4g

import (
	"os"

	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"gopkg.in/natefinch/lumberjack.v2"
)

var loggers map[string]*logrus.Logger
var formatter *prefixed.TextFormatter

func init() {
	loggers = make(map[string]*logrus.Logger)
	loadConfigure("./log4g.json")

	formatter = &prefixed.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05.000",
	}
}

func GetLoggerWithId(t, id string) *logrus.Entry {
	logger := getLogger(t)
	return logger.WithField("requestId", id)
}

func GetLogger(t string) *logrus.Entry {
	logger := getLogger(t)
	return logger.WithFields(logrus.Fields{})
}

func getLogger(t string) *logrus.Logger {
	if logger, ok := loggers[t]; ok {
		return logger
	}

	opt := getOption(t)
	logger := newLogger(opt)
	loggers[t] = logger
	return logger
}

func newLogger(opt *option) *logrus.Logger {
	logger := logrus.New()
	logger.SetLevel(opt.logruslvl)
	logger.Formatter = formatter
	if opt.Xtype == "stdout" {
		logger.Out = os.Stdout
	} else {
		output := &lumberjack.Logger{
			Filename:   opt.FileName,
			MaxSize:    opt.MaxLogSize,
			MaxBackups: opt.Backups,
		}
		logger.Out = output
	}
	return logger
}
