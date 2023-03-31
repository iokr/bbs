package log

import (
	"fmt"
	"sync"
)

// Helper is a helper logger
type Helper struct {
	// logger interface
	logger Logger

	// Reusable empty entry
	entryPool sync.Pool
}

func NewHelper(logger Logger) *Helper {
	return &Helper{
		logger: logger,
	}
}

func (h *Helper) newEntry() *Entry {
	entry, ok := h.entryPool.Get().(*Entry)
	if ok {
		return entry
	}
	return NewEntry(h.logger)
}

func (h *Helper) releaseEntry(entry *Entry) {
	h.entryPool.Put(entry)
}

func (h *Helper) WithTraceId(traceId string) *Entry {
	entry := h.newEntry()
	defer h.releaseEntry(entry)
	return entry.WithTraceId(traceId)
}

func (h *Helper) Log(level Level, v ...interface{}) error {
	h.logger.Log(level, v...)
	return nil
}

func (h *Helper) Debug(v ...interface{}) {
	h.logger.Log(Debug, v...)
}

func (h *Helper) Debugf(format string, v ...interface{}) {
	h.logger.Log(Debug, fmt.Sprintf(format, v...))
}

func (h *Helper) Info(v ...interface{}) {
	h.logger.Log(Info, v...)
}

func (h *Helper) Infof(format string, v ...interface{}) {
	h.logger.Log(Info, fmt.Sprintf(format, v...))
}

func (h *Helper) Warn(v ...interface{}) {
	h.logger.Log(Warn, v...)
}

func (h *Helper) Warnf(format string, v ...interface{}) {
	h.logger.Log(Warn, fmt.Sprintf(format, v...))
}

func (h *Helper) Error(v ...interface{}) {
	h.logger.Log(Error, v...)
}

func (h *Helper) Errorf(format string, v ...interface{}) {
	h.logger.Log(Error, fmt.Sprintf(format, v...))
}
