package middleware

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/net/ghttp"
)

func ResponseSuccess(r *ghttp.Request, data interface{}) {
	r.Response.WriteJson(ghttp.DefaultHandlerResponse{
		Code:    gcode.CodeOK.Code(),
		Message: gcode.CodeOK.Message(),
		Data:    data,
	})
}

func ResponseFail(r *ghttp.Request, err error) {
	r.Response.WriteJson(ghttp.DefaultHandlerResponse{
		Code:    gcode.CodeInternalError.Code(),
		Message: err.Error(),
		Data:    nil,
	})
}
