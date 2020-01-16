// +build !statik

package statik

import (
	"net/http"
)

func Handler() http.Handler {
	return http.FileServer(http.Dir("vue/dist"))
}
