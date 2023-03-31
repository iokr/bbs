package trace

import (
	"context"

	"github.com/iokr/bbs/core/utils"
)

type Trace struct {
	TraceId string
	SpanId  string
}

func New(traceId string) *Trace {
	t := &Trace{
		TraceId: traceId,
	}

	if t.TraceId == "" {
		t.TraceId = utils.NewPureUuid()
	}
	return t
}

func WithContext(ctx context.Context) *Trace {
	t, ok := ctx.Value(TracingKey).(*Trace)
	if !ok {
		return &Trace{}
	}
	return t
}
