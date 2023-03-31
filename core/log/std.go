package log

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"sync"
)

var _ Logger = (*stdLogger)(nil)

type stdLogger struct {
	log  *log.Logger
	pool *sync.Pool
}

// NewStdLogger new a logger with writer.
func NewStdLogger(w io.Writer) Logger {
	return &stdLogger{
		log: log.New(w, "", log.Ldate|log.Lmicroseconds|log.Lshortfile),
		pool: &sync.Pool{
			New: func() interface{} {
				return new(bytes.Buffer)
			},
		},
	}
}

// Log print the kv pairs log.
func (l *stdLogger) Log(level Level, v ...interface{}) error {
	if len(v) == 0 {
		return nil
	}

	buf := l.pool.Get().(*bytes.Buffer)
	buf.WriteString(level.String())
	buf.WriteByte(' ')
	buf.WriteString(fmt.Sprint(v...))
	_ = l.log.Output(3, buf.String())
	buf.Reset()
	l.pool.Put(buf)
	return nil
}
