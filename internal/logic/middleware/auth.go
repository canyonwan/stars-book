package middleware

import (
	"net/http"

	"gf-playground-1/internal/consts"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/golang-jwt/jwt/v5"
)

func Auth(r *ghttp.Request) {
	var tokenString = r.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(consts.JwtKey), nil
	})
	if err != nil || !token.Valid {
		r.Response.WriteJson(g.Map{
			"code":    http.StatusUnauthorized,
			"message": "token失效,请重新登录",
			"data":    nil,
		})
		r.Exit()
		return
	}
	claims, _ := token.Claims.(jwt.MapClaims)
	r.SetCtxVar("userId", claims["UserId"])
	r.Middleware.Next()
}
