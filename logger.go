package StreamDeck

import (
	"log"
)

type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
}

type StdLogger struct{}

func NewStdLogger() *StdLogger {
	return &StdLogger{}
}

func (stdLogger *StdLogger) Debug(args ...interface{}) {
	log.Print(args)
}

func (stdLogger *StdLogger) Debugf(format string, args ...interface{}) {
	log.Printf(format, args)
}

func (stdLogger *StdLogger) Info(args ...interface{}) {
	log.Print(args)
}

func (stdLogger *StdLogger) Infof(format string, args ...interface{}) {
	log.Printf(format, args)
}

func (stdLogger *StdLogger) Warn(args ...interface{}) {
	log.Print(args)
}

func (stdLogger *StdLogger) Warnf(format string, args ...interface{}) {
	log.Printf(format, args)
}

func (stdLogger *StdLogger) Error(args ...interface{}) {
	log.Print(args)
}

func (stdLogger *StdLogger) Errorf(format string, args ...interface{}) {
	log.Printf(format, args)
}
