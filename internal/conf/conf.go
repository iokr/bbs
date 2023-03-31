package conf

import (
	"fmt"
	"os"
	"time"

	"github.com/iokr/bbs/core/cache/redis"
	"github.com/iokr/bbs/core/database/orm"
	"github.com/iokr/bbs/core/env"
	"gopkg.in/yaml.v3"
)

type (
	Config struct {
		filePath string
		Server   *Server
		Data     *Data
	}

	Server struct {
		HTTP *HTTP
	}

	HTTP struct {
		Addr         string
		ReadTimeout  time.Duration
		WriteTimeout time.Duration
	}

	Data struct {
		Database *orm.Config
		Redis    *redis.Client
	}
)

func New(opts ...Option) *Config {
	if err := env.Init(); err != nil {
		panic(err)
	}

	filePath := fmt.Sprintf("./configs/%s/config.yaml", env.GetEnv())
	c := &Config{
		filePath: filePath,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func (c *Config) Load() error {
	yamlFile, err := os.ReadFile(c.filePath)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(yamlFile, c)
}
