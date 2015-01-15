package log

import "github.com/Sirupsen/logrus"

var Logger = logrus.New()

func Info(args ...interface{}) {
	Logger.Info(args...)
}

func Infof(fmt string, args ...interface{}) {
	Logger.Infof(fmt, args...)
}

func Error(args ...interface{}) {
	Logger.Error(args...)
}

func Errorf(fmt string, args ...interface{}) {
	Logger.Errorf(fmt, args...)
}

func Fatal(args ...interface{}) {
	Logger.Fatal(args...)
}

func Fatalf(fmt string, args ...interface{}) {
	Logger.Fatalf(fmt, args...)
}

func Debug(args ...interface{}) {
	Logger.Debug(args...)
}

func Debugf(fmt string, args ...interface{}) {
	Logger.Debugf(fmt, args...)
}
