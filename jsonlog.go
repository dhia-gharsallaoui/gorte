package gorte

import (
	gofmt "fmt"
	golog "log"
	"os"
	"strconv"
	"time"
)

type jsonLogger struct {
	out *golog.Logger
	err *golog.Logger
	*LoggerConfiguration
}

func NewLogger(cfg *LoggerConfiguration) Logger {
	return &jsonLogger{out: golog.New(os.Stdout, "", 0), err: golog.New(os.Stderr, "", 0), LoggerConfiguration: cfg}
}

func (l *jsonLogger) SetVerbosity(Verbosity string) {
	l.Verbosity = GetVerbosityFromString(Verbosity)
}

func (l *jsonLogger) Debug(fmt string, v ...interface{}) {
	if l.Verbosity <= Debug {
		l.out.Println(l.format("debug", fmt, v...))
	}
}

func (l *jsonLogger) Info(fmt string, v ...interface{}) {
	if l.Verbosity <= Info {
		l.out.Println(l.format("info", fmt, v...))
	}
}

func (l *jsonLogger) Warn(fmt string, v ...interface{}) {
	if l.Verbosity <= Warn {
		l.out.Println(l.format("warn", fmt, v...))
	}
}

func (l *jsonLogger) Err(fmt string, v ...interface{}) {
	if l.Verbosity <= Err {
		l.err.Println(l.format("error", fmt, v...))
	}
}

func (l *jsonLogger) Panic(fmt string, v ...interface{}) {
	l.err.Panicln(l.format("panic", fmt, v...))
}

func (l *jsonLogger) Fatal(fmt string, v ...interface{}) {
	l.err.Fatalln(l.format("fatal", fmt, v...))
}

func (l *jsonLogger) format(level string, fmt string, v ...interface{}) string {
	return `{"time": "` + time.Now().Format(time.RFC3339Nano) + `", "level": "` + level + `", "message": ` + strconv.Quote(gofmt.Sprintf("%s"+fmt, append([]interface{}{l.Prefix}, v...)...)) + `}`
}
