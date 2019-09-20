package log

import (
	golog "log"
)

type logLogger struct{}

func (t *logLogger) Log(v ...interface{}) {
	golog.Print(v...)
}

func (t *logLogger) Logf(format string, v ...interface{}) {
	golog.Printf(format, v...)
}

func New() *logLogger {
	return &logLogger{}
}
