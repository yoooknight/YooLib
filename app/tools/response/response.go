package response

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

const (
	Success       = 200
	Fail          = 400
	InvalidToken  = 401
	InvalidParams = 402
)

type RespBody struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Json(w http.ResponseWriter, data ...interface{}) {
	var body = RespBody{
		Code:    Success,
		Message: "Success",
	}

	for _, datum := range data {
		if v, ok := datum.(int); ok {
			body.Code = v
		} else if v, ok := datum.(string); ok {
			body.Message = v
		} else {
			body.Data = datum
		}
	}

	httpx.OkJson(w, body)
}
