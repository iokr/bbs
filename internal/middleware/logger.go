package middleware

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"

	"github.com/iokr/bbs/core/log"
	"github.com/iokr/bbs/core/trace"
)

func Logger1(logger *log.Helper) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				fmt.Println("operation:", tr.Operation())
				//// Start timer
				//startTime := time.Now()
				//
				//t := trace.New(tr.RequestHeader().Get("X-Trace-ID"))
				//
				//tr.RequestHeader().
			}
			reply, err = handler(ctx, req)

			//logger.WithTraceId(t.TraceId).Infof("method:%s code:%d latency:%v addr:%s path:%s",
			//	c.Request.Method,
			//	c.Writer.Status(),
			//	time.Now().Sub(startTime),
			//	c.ClientIP(),
			//	c.Request.URL.Path,
			//)

			return
		}
	}
}

// Logger instance a Logger middleware with config.
func Logger(logger *log.Helper) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		startTime := time.Now()

		// Set traceId
		t := trace.New(c.Request.Header.Get("X-Trace-ID"))
		c.Set("X-Trace", t)

		// Process request
		c.Next()

		logger.WithTraceId(t.TraceId).Infof("method:%s code:%d latency:%v addr:%s path:%s",
			c.Request.Method,
			c.Writer.Status(),
			time.Now().Sub(startTime),
			c.ClientIP(),
			c.Request.URL.Path,
		)
	}
}
