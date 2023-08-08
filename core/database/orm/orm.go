package orm

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	Driver      string
	Protocol    string
	UserName    string
	Password    string
	Address     string
	Port        int
	DbName      string
	Params      string
	MaxIdleConn int
	MaxOpenConn int
	KeepAlive   time.Duration
}

type Client struct {
	username string
	password string
	protocol string
	address  string
	port     int
	dbName   string
	params   string

	db  *gorm.DB
	dsn string

	pool *Pool
}

var (
	defaultProtocol = "tcp"
	defaultAddress  = "127.0.0.1"
	defaultPort     = 3306
	defaultParams   = "charset=utf8mb4&sql_notes=false&timeout=90s&collation=utf8mb4_general_ci&parseTime=True&loc=Local"
)

func NewClient(username, password, dbName string, opts ...Option) (*Client, error) {
	cli := &Client{
		username: username,
		password: password,
		protocol: defaultProtocol,
		address:  defaultAddress,
		port:     defaultPort,
		dbName:   dbName,
		params:   defaultParams,
	}

	for _, opt := range opts {
		opt(cli)
	}

	cli.dsn = fmt.Sprintf("%s:%s@%s(%s:%d)/%s?%s",
		cli.username, cli.password, cli.protocol, cli.address, cli.port, cli.dbName, cli.params)

	var err error
	cli.db, err = gorm.Open(mysql.Open(cli.dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	return cli, nil
}

func NewMySQL(c *Config) (db *gorm.DB, err error) {
	dataSourceName := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?%s",
		c.UserName, c.Password, c.Protocol, c.Address, c.Port, c.DbName, c.Params)
	db, err = gorm.Open(mysql.Open(dataSourceName), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	sqlDb, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDb.SetMaxIdleConns(c.MaxIdleConn)
	sqlDb.SetMaxOpenConns(c.MaxOpenConn)
	sqlDb.SetConnMaxLifetime(c.KeepAlive)
	return db, nil
}
