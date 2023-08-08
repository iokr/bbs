package orm

import (
	"time"
)

type PoolOption func(p *Pool)

var (
	defaultMaxIdleConn = 10
	defaultMaxOpenConn = 20
	defaultKeepAlive   = time.Second * 3600
)

type Pool struct {
	maxIdleConn int
	maxOpenConn int
	keepAlive   time.Duration
}

func NewPool(opts ...PoolOption) *Pool {
	p := &Pool{
		maxIdleConn: defaultMaxIdleConn,
		maxOpenConn: defaultMaxOpenConn,
		keepAlive:   defaultKeepAlive,
	}

	for _, opt := range opts {
		opt(p)
	}

	return p
}
