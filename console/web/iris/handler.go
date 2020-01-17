package iris

import (
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

func New() (http.Handler, error) {
	// Iris
	app := iris.New()
	app.Logger().SetLevel("debug")

	// 避免'/'被消除后重定向
	app.Configure(iris.WithoutPathCorrectionRedirection)

	app.Use(recover.New())
	app.Use(logger.New())

	v1 := app.Party("/v1/iris")
	{
		v1.Get("/", func(ctx iris.Context) {
			ctx.JSON(iris.Map{"message": "Hello Iris!"})
		})
	}

	err := app.Build()

	return app, err
}
