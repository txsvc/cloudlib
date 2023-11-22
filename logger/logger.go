package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/alecthomas/chroma/quick"
	"github.com/logrusorgru/aurora"

	"github.com/txsvc/stdlib/v2"
)

// FIXME: replace with https://github.com/labstack/gommon/blob/master/log/log.go

type Level int

const (
	Debug Level = iota
	Info
	Warn
	Error

	LogLevelEnv = "LOG_LEVEL"
)

type Logger struct {
	out   io.Writer
	level Level
}

func New() Logger {
	return NewLogger(os.Stdout, "")
}

func NewLogger(out io.Writer, lvl string) Logger {
	return Logger{
		out:   out,
		level: levelFromEnv(lvl),
	}
}

// levelFromEnv looks for ENV['LOG_LEVEL'], returns level Info by default
func levelFromEnv(lvl string) Level {

	lit := strings.ToLower(stdlib.GetString(LogLevelEnv, lvl))

	switch lit {
	default:
		return Info
	case "debug":
		return Debug
	case "warn":
		return Warn
	case "error":
		return Error
	}
}

func (l *Logger) debug(v ...interface{}) {
	if str, ok := v[0].(string); ok {
		byteString := []byte(str)
		if json.Valid(byteString) {
			var prettyJSON bytes.Buffer
			err := json.Indent(&prettyJSON, byteString, "", "  ")
			if err == nil {
				_ = quick.Highlight(l.out, prettyJSON.String()+"\n", "json", "terminal", "monokai")

				return
			}
		}
	}

	fmt.Fprintln(
		l.out,
		aurora.Faint("DEBUG"),
		fmt.Sprint(v...),
	)
}

func (l *Logger) Debug(v ...interface{}) {
	if l.level <= Debug {
		l.debug(v...)
	}
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.level <= Debug {
		l.debug(fmt.Sprintf(format, v...))
	}
}

func (l *Logger) info(v ...interface{}) {
	fmt.Fprintln(
		l.out,
		aurora.Faint("INFO"),
		fmt.Sprint(v...),
	)
}

func (l *Logger) Info(v ...interface{}) {
	if l.level <= Info {
		l.info(v...)
	}
}

func (l *Logger) Infof(format string, v ...interface{}) {
	if l.level <= Info {
		l.info(fmt.Sprintf(format, v...))
	}
}

func (l *Logger) warn(v ...interface{}) {
	fmt.Fprintln(
		l.out,
		aurora.Yellow("WARN"),
		fmt.Sprint(v...),
	)
}

func (l *Logger) Warn(v ...interface{}) {
	if l.level <= Warn {
		l.warn(v...)
	}
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	if l.level <= Warn {
		l.warn(fmt.Sprintf(format, v...))
	}
}

func (l *Logger) error(v ...interface{}) {
	fmt.Fprintln(
		l.out,
		aurora.Red("ERROR"),
		fmt.Sprint(v...),
	)
}

func (l *Logger) Error(v ...interface{}) {
	if l.level <= Error {
		l.error(v...)
	}
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	if l.level <= Error {
		l.error(fmt.Sprintf(format, v...))
	}
}
