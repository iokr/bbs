package server

import (
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/iokr/bbs/core/log"
	"github.com/iokr/bbs/internal/conf"
	"github.com/iokr/bbs/internal/middleware"
)

func NewHTTPServer(c *conf.Server, logger *log.Helper) (srv *http.Server, engine *gin.Engine) {
	opts := []http.ServerOption{
		http.Address(":80"),
	}

	if c.HTTP.Addr != "" {
		opts = append(opts, http.Address(c.HTTP.Addr))
	}
	gin.SetMode(gin.DebugMode)

	ginEngine := gin.New()
	ginEngine.Use(gin.Recovery(), middleware.Logger(logger))

	srv = http.NewServer(opts...)
	srv.HandlePrefix("/", ginEngine)
	return srv, ginEngine
}
