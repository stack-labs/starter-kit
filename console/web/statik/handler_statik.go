// +build statik

package statik

import (
	"net/http"

	"github.com/micro/go-micro/v2/util/log"
	"github.com/rakyll/statik/fs"
)

func Handler(prefix string) http.Handler {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	return http.FileServer(statikFS)
}
