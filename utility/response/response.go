package response

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(r *ghttp.Request, data interface{}) {
	r.Response.WriteJson(Response{
		Code:    0,
		Message: "OK",
		Data:    data,
	})
}

func Fail(r *ghttp.Request, err error) {
	r.Response.WriteJson(Response{
		Code:    1,
		Message: err.Error(),
		Data:    nil,
	})
}
