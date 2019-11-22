package model

type User struct {
	Id       int64  `xorm:"bigint pk autoincr"`
	Name     string `xorm:"varchar(64)"`
	Password string `xorm:"varchar(64)"`
}
