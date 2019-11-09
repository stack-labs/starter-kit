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

	v1 := app.Party("/v1")
	{
		// Method:   GET
		// Resource: http://localhost:8080
		v1.Get("/", func(ctx iris.Context) {
			ctx.HTML("<h1>Welcome</h1>")
		})

		// same as app.Handle("GET", "/ping", [...])
		// Method:   GET
		// Resource: http://localhost:8080/ping
		v1.Get("/ping", func(ctx iris.Context) {
			ctx.WriteString("pong")
		})

		// Method:   GET
		// Resource: http://localhost:8080/hello
		v1.Get("/hello", func(ctx iris.Context) {
			ctx.JSON(iris.Map{"message": "Hello Iris!"})
		})
	}

	err := app.Build()

	return app, err
}
