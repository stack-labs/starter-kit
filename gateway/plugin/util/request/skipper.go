package request

import (
	"net/http"
)

type SkipperFunc func(r *http.Request) bool

var DefaultSkipperFunc = func(r *http.Request) bool {
	return false
}
