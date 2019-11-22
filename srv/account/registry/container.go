package registry

import (
	"github.com/micro-in-cn/starter-kit/srv/account/conf"
	"github.com/micro-in-cn/starter-kit/srv/account/domain/service"
	"github.com/micro-in-cn/starter-kit/srv/account/interface/persistence/memory"
	"github.com/micro-in-cn/starter-kit/srv/account/interface/persistence/xorm"
	"github.com/micro-in-cn/starter-kit/srv/account/usecase"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/util/log"
	"github.com/sarulabs/di"
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
	dbConf := conf.Database{}
	err := config.Get("database").Scan(&dbConf)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// TODO ORM选择，gorm、xorm...
	if dbConf.Engine == "mysql" {
		repo := xorm.NewUserRepository()
		service := service.NewUserService(repo)
		return usecase.NewUserUsecase(repo, service), nil
	} else {
		// 默认memory作为mock
		repo := memory.NewUserRepository()
		service := service.NewUserService(repo)
		return usecase.NewUserUsecase(repo, service), nil
	}
}
