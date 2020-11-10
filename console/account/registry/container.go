package registry

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/sarulabs/di"
	"github.com/stack-labs/stack-rpc/pkg/config"

	"github.com/stack-labs/starter-kit/console/account/domain/repository"
	"github.com/stack-labs/starter-kit/console/account/domain/repository/persistence/gorm"
	"github.com/stack-labs/starter-kit/console/account/domain/repository/persistence/memory"
	"github.com/stack-labs/starter-kit/console/account/domain/repository/persistence/xorm"
	"github.com/stack-labs/starter-kit/console/account/domain/service"
	"github.com/stack-labs/starter-kit/console/account/usecase"
)

type Container struct {
	ctn di.Container
}

func NewContainer() (*Container, error) {
	builder, err := di.NewBuilder()
	if err != nil {
		return nil, err
	}

	if err := builder.Add([]di.Def{
		{
			Name:  "user-usecase",
			Build: buildUserUsecase,
		},
	}...); err != nil {
		return nil, err
	}

	return &Container{
		ctn: builder.Build(),
	}, nil
}

func (c *Container) Resolve(name string) interface{} {
	return c.ctn.Get(name)
}

func (c *Container) Clean() error {
	return c.ctn.Clean()
}

func buildUserUsecase(ctn di.Container) (interface{}, error) {
	// TODO stack-rpc
	conf, _ := config.NewConfig()
	persistence := conf.Get("persistence").String("")

	// ORM选择，gorm、xorm...
	var repo repository.UserRepository
	switch persistence {
	case "xorm":
		// DB初始化
		xorm.InitDB()
		repo = xorm.NewUserRepository()
	case "gorm":
		// DB初始化
		gorm.InitDB()
		repo = gorm.NewUserRepository()
	default:
		// 默认memory作为mock
		repo = memory.NewUserRepository()
	}

	service := service.NewUserService(repo)
	return usecase.NewUserUsecase(repo, service), nil
}
