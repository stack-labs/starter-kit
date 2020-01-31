package beego

import (
	"net/http"

	"github.com/astaxie/beego"
	bctx "github.com/astaxie/beego/context"
	"github.com/micro/go-micro/v2/util/log"
)

//New return a beego handler
func New() (http.Handler, error) {
	beego.Get("/v1/beego", func(ctx *bctx.Context) {
		log.Info("Received Get Request")
		ctx.Output.JSON(map[string]string{
			"message": "BeeGo Here",
		}, false, true)
	})
	return beego.BeeApp.Handlers, nil
}
