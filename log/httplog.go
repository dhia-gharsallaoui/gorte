package log

import (
	"strings"
)

type httpLogger struct {
	logger Logger
}

func NewHTTPLogger(l Logger) HTTPLogger {
	return &httpLogger{logger: l}
}

func (l *httpLogger) Debug(msg string, keysAndValues ...interface{}) {
	l.logger.Debug(msgToFormat(msg, keysAndValues...), keysAndValues...)
}

func (l *httpLogger) Info(msg string, keysAndValues ...interface{}) {
	l.logger.Info(msgToFormat(msg, keysAndValues...), keysAndValues...)
}

func (l *httpLogger) Warn(msg string, keysAndValues ...interface{}) {
	l.logger.Warn(msgToFormat(msg, keysAndValues...), keysAndValues...)
}

func (l *httpLogger) Error(msg string, keysAndValues ...interface{}) {
	l.logger.Err(msgToFormat(msg, keysAndValues...), keysAndValues...)
}

func msgToFormat(msg string, keysAndValues ...interface{}) string {
	return msg + strings.Repeat(" - %v: %v", len(keysAndValues)/2)
}
