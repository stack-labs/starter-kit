package conf

import (
	"time"
)

type Database struct {
	Engine   string
	Host     string
	Port     string
	User     string
	Password string
	Name     string

	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}
