package response

import (
	"net/http"

	"github.com/micro/go-micro/v2/errors"
)

//Handler for http
type Handler func(w http.ResponseWriter, r *http.Request, err error)

//DefaultResponseHandler of http
func DefaultResponseHandler(w http.ResponseWriter, r *http.Request, err error) {
	mErr, ok := err.(*errors.Error)
	if !ok {
		mErr = &errors.Error{
			Code:   http.StatusInternalServerError,
			Detail: err.Error(),
			Status: http.StatusText(http.StatusInternalServerError),
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(mErr.Code))
	w.Write([]byte(mErr.Error()))
}
