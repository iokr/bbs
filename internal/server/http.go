package server

import (
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/iokr/bbs/core/log"
	"github.com/iokr/bbs/internal/conf"
	"github.com/iokr/bbs/internal/middleware"
)

func NewHTTPServer(c *conf.Server, logger *log.Helper) *http.Server {
	opts := []http.ServerOption{
		http.Address(":80"),
	}

	if c.HTTP.Addr != "" {
		opts = append(opts, http.Address(c.HTTP.Addr))
	}
	gin.SetMode(gin.DebugMode)

	router := gin.New()
	router.Use(gin.Recovery(), middleware.Logger(logger))

	srv := http.NewServer(opts...)
	srv.HandlePrefix("/", router)
	return srv
}
