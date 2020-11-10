// +build statik

package statik

import (
	"net/http"

	"github.com/rakyll/statik/fs"
	"github.com/stack-labs/stack-rpc/util/log"
)

func Handler() http.Handler {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	return http.FileServer(statikFS)
}
