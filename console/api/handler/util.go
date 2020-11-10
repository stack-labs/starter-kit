package handler

import (
	"encoding/json"

	api "github.com/stack-labs/stack-rpc/api/proto"
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

func extractValue(pair *api.Pair) string {
	if pair == nil {
		return ""
	}
	if len(pair.Values) == 0 {
		return ""
	}
	return pair.Values[0]
}
