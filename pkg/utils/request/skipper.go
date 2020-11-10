package request

import (
	"net/http"
)

//SkipperFunc ...
type SkipperFunc func(r *http.Request) bool

//DefaultSkipperFunc ...
var DefaultSkipperFunc = func(r *http.Request) bool {
	return false
}
