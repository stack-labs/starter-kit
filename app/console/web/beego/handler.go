package beego

import (
	"context"
	"github.com/micro/go-micro/util/log"
	bctx "github.com/astaxie/beego/context"
	"github.com/astaxie/beego"
)


//New return a beego handler
func New() (http.Handler, error) {

	beego.GET("/", (ctx *bctx.Context) {
		log.Info("Received Get Request")
		ctx.Output.JSON(map[string]string{
			"message": "BeeGo Here",
		}, false, true)
	})
	return 	beego.BeeApp.Handlers, nil
}