package log

import (
	"log"
)

var (
	DefaultLogger = NewStdLogger(log.Writer())
)

// Logger is a logger interface
type Logger interface {
	Log(level Level, v ...interface{}) error
}
