package registry

import (
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/dig"

	"github.com/micro-in-cn/starter-kit/console/account/domain/repository/persistence/gorm"
	"github.com/micro-in-cn/starter-kit/console/account/domain/repository/persistence/memory"
	"github.com/micro-in-cn/starter-kit/console/account/domain/repository/persistence/xorm"
	"github.com/micro-in-cn/starter-kit/console/account/domain/service"
	"github.com/micro-in-cn/starter-kit/console/account/usecase"
)

func NewContainer() (*dig.Container, error) {
	c := dig.New()

	buildUserUsecase(c)

	return c, nil
}

func buildUserUsecase(c *dig.Container) {
	persistence := "" //config.Get("persistence").String("")

	// ORM选择，gorm、xorm...
	switch persistence {
	case "xorm":
		// DB初始化
		xorm.InitDB()
		c.Provide(xorm.NewUserRepository)
	case "gorm":
		// DB初始化
		gorm.InitDB()
		c.Provide(gorm.NewUserRepository)
	default:
		// 默认memory作为mock
		c.Provide(memory.NewUserRepository)
	}

	c.Provide(service.NewUserService)
	c.Provide(usecase.NewUserUsecase)
}
