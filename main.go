package main

import (
	"os"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/iokr/bbs/core/database/orm"
	"github.com/iokr/bbs/core/log"
	"github.com/iokr/bbs/internal/conf"
	"github.com/iokr/bbs/internal/server"
)

const (
	Name = "bbs"
)

func main() {
	config := conf.New()
	if err := config.Load(); err != nil {
		panic(err)
	}

	app, cleanup, err := initApp(config)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal.
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func initApp(c *conf.Config) (*kratos.App, func(), error) {
	bbsDb, err := orm.NewMySQL(c.Data.Database)
	if err != nil {
		panic(err)
	}

	_ = bbsDb

	logger := log.NewStdLogger(os.Stdout)

	httpServer := server.NewHTTPServer(c.Server, log.NewHelper(logger))

	app := newApp(httpServer)
	return app, func() {

	}, nil
}

func newApp(hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.Name(Name),
		kratos.Metadata(map[string]string{}),
		kratos.Server(
			hs,
		),
	)
}
