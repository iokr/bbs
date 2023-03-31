package orm

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

func NewMySQL(c *Config) (db *gorm.DB, err error) {
	dataSourceName := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?%s",
		c.UserName, c.Password, c.Protocol, c.Address, c.Port, c.DbName, c.Params)
	db, err = gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
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
