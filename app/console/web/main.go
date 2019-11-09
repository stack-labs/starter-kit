package main

//go:generate statik -src=./vue/dist -dest=./ -f
import (
	"net/http"

	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/web"
	"github.com/rakyll/statik/fs"

	"github.com/micro-in-cn/starter-kit/app/console/web/iris"
	_ "github.com/micro-in-cn/starter-kit/app/console/web/statik"
)

func main() {
	// create new web service
	service := web.NewService(
		web.Name("go.micro.web.console"),
		web.Version("latest"),
	)

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	// Iris
	h, err := iris.New()
	if err != nil {
		log.Fatal(err)
	}
	service.Handle("/v1/", h)

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	// register html handler
	service.Handle("/", http.FileServer(statikFS))

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
