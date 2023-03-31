package log

import (
	"bytes"
	"fmt"
)

// Entry is logger entry
type Entry struct {
	logger Logger

	traceId string
	message string
}

func NewEntry(logger Logger) *Entry {
	return &Entry{
		logger: logger,
	}
}

// WithTraceId set trace id
func (e *Entry) WithTraceId(traceId string) *Entry {
	return &Entry{logger: e.logger, traceId: traceId}
}

func (e *Entry) format() []byte {
	b := new(bytes.Buffer)

	b.WriteString("trace_id:")
	if e.traceId == "" {
		e.traceId = "null"
	}
	b.WriteString(e.traceId)

	appendKeyValue(b, "", e.message)

	return b.Bytes()
}

func (e *Entry) Debug(v ...interface{}) {
	e.message = fmt.Sprint(v...)
	e.logger.Log(Debug, string(e.format()))
}

func (e *Entry) Debugf(format string, v ...interface{}) {
	e.message = fmt.Sprintf(format, v...)
	e.logger.Log(Debug, string(e.format()))
}

func (e *Entry) Info(v ...interface{}) {
	e.message = fmt.Sprint(v...)
	e.logger.Log(Info, string(e.format()))
}

func (e *Entry) Infof(format string, v ...interface{}) {
	e.message = fmt.Sprintf(format, v...)
	e.logger.Log(Info, string(e.format()))
}

func appendKeyValue(b *bytes.Buffer, key string, value interface{}) {
	if b.Len() > 0 {
		b.WriteByte(' ')
	}
	if key != "" {
		b.WriteString(key)
		b.WriteByte('=')
	}
	appendValue(b, value)
}

func appendValue(b *bytes.Buffer, value interface{}) {
	stringVal, ok := value.(string)
	if !ok {
		stringVal = fmt.Sprint(value)
	}
	b.WriteString(stringVal)
}
