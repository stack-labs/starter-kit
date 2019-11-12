package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func New() (http.Handler, error) {
	g := gin.Default()
	r := g.Group("/v1/gin")
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Gin",
		})
	})

	return g, nil
}
