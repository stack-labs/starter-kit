package gorm

import (
	"sync"

	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/util/log"

	"github.com/micro-in-cn/starter-kit/console/account/conf"
	"github.com/micro-in-cn/starter-kit/console/account/domain/model"
)

var (
	dbConf conf.Database
	db     *gorm.DB
	once   sync.Once
)

func InitDB() {
	once.Do(func() {
		dbConf = conf.Database{}
		err := config.Get("database").Scan(&dbConf)
		if err != nil {
			log.Fatal(err)
		}

		sqlConnection := dbConf.User + ":" + dbConf.Password + "@tcp(" + dbConf.Host + ":" + dbConf.Port + ")/" + dbConf.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
		db, err = gorm.Open(dbConf.Engine, sqlConnection)
		if err != nil {
			log.Fatal(err)
		}

		db.DB().SetMaxOpenConns(dbConf.MaxOpenConns)
		db.DB().SetMaxIdleConns(dbConf.MaxIdleConns)
		db.DB().SetConnMaxLifetime(dbConf.ConnMaxLifetime)

		db.SingularTable(true)
		err = db.AutoMigrate(&model.User{}).Error
		if err != nil {
			log.Fatal(err)
		}

		// TODO 数据初始化
	})
}
