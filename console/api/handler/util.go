package handler

import (
	"encoding/json"
	"errors"
	"strconv"

	api "github.com/micro/go-micro/v2/api/proto"
)

type responseBody struct {
	Code   int64       `json:"code"`
	Detail string      `json:"detail"`
	Data   interface{} `json:"data"`
}

func ResponseBody(code int64, data interface{}, detail ...string) (string, error) {
	body := responseBody{
		Code: code,
		Data: data,
	}

	if len(detail) > 0 {
		body.Detail = detail[0]
	}

	b, err := json.Marshal(body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

var ErrPairValueNotfound = errors.New("pair extract value not found")

func extractValue(pair *api.Pair) string {
	if pair == nil {
		return ""
	}
	if len(pair.Values) == 0 {
		return ""
	}
	return pair.Values[0]
}

func extractValueInt64(pair *api.Pair) (int64, error) {
	if pair == nil {
		return 0, ErrPairValueNotfound
	}
	if len(pair.Values) == 0 {
		return 0, ErrPairValueNotfound
	}

	return strconv.ParseInt(pair.Values[0], 10, 64)
}
