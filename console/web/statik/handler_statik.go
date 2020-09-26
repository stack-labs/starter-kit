// +build statik

package statik

import (
	"net/http"

	log "github.com/micro/go-micro/v3/logger"
	"github.com/rakyll/statik/fs"
)

func Handler() http.Handler {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	return http.FileServer(statikFS)
}
