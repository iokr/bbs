package log

import (
	"strings"
)

type Level uint8

const (
	Debug Level = iota
	Info
	Warn
	Error
	Fatal
)

func (l Level) String() string {
	switch l {
	case Debug:
		return "DEBUG"
	case Info:
		return "INFO"
	case Warn:
		return "WARN"
	case Error:
		return "ERROR"
	case Fatal:
		return "FATAL"
	}
	return ""
}

// ParseLevel .
func ParseLevel(s string) Level {
	switch strings.ToUpper(s) {
	case "DEBUG":
		return Debug
	case "INFO":
		return Info
	case "WARN":
		return Warn
	case "ERROR":
		return Error
	case "FATAL":
		return Fatal
	}
	return Info
}
