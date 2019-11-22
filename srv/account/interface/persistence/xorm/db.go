package xorm

import (
	"fmt"
	"sync"

	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/util/log"
	"xorm.io/xorm"
	"xorm.io/xorm/migrate"

	"github.com/micro-in-cn/starter-kit/srv/account/conf"
)

var dbConf conf.Database
var db *xorm.Engine
var once sync.Once

func InitDB() {
	once.Do(func() {
		dbConf = conf.Database{}
		err := config.Get("database").Scan(&dbConf)
		if err != nil {
			log.Fatal(err)
		}

		db, err = xorm.NewEngine("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
			dbConf.User,
			dbConf.Password,
			dbConf.Host,
			dbConf.Port,
			dbConf.Name,
		))
		if err != nil {
			log.Fatal(err)
		}

		// TODO xorm migrate问题，mysql创建migrations表出错
		// Specified key was too long; max key length is 767 bytes
		options := migrate.DefaultOptions
		exists, err := db.IsTableExist(options.TableName)
		if err != nil {
			panic(err)
		}
		if !exists {
			sql := fmt.Sprintf("CREATE TABLE %s (%s VARCHAR(64) PRIMARY KEY)", options.TableName, options.IDColumnName)
			if _, err := db.Exec(sql); err != nil {
				panic(err)
			}
		}

		m := migrate.New(db, migrate.DefaultOptions, migrations)
		err = m.Migrate()
		if err != nil {
			panic(err)
		}
	})
}

func DB() *xorm.Engine {
	if db == nil {
		InitDB()
	}
	return db
}
