package echo

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() (http.Handler, error) {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	r := e.Group("/v1/echo")
	r.GET("/", func(ctx echo.Context) error {
		return ctx.JSON(200, echo.Map{
			"message": "Hello Echo",
		})
	})

	return e, nil
}
