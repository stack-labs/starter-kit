package response

import (
	"net/http"
)

type WrapWriter struct {
	StatusCode int
	http.ResponseWriter
}

func (ww *WrapWriter) WriteHeader(statusCode int) {
	ww.StatusCode = statusCode
	ww.ResponseWriter.WriteHeader(statusCode)
}
