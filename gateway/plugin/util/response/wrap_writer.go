package response

import (
	"net/http"
)

type WrapWriter struct {
	StatusCode  int
	Size        int64
	wroteHeader bool

	http.ResponseWriter
}

func (ww *WrapWriter) WriteHeader(statusCode int) {
	ww.wroteHeader = true
	ww.StatusCode = statusCode
	ww.ResponseWriter.WriteHeader(statusCode)
}

func (ww *WrapWriter) Write(b []byte) (n int, err error) {
	// 默认200
	if !ww.wroteHeader {
		ww.WriteHeader(http.StatusOK)
	}

	n, err = ww.ResponseWriter.Write(b)
	ww.Size += int64(n)
	return
}
