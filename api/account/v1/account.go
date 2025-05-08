package v1

import "github.com/gogf/gf/v2/frame/g"

type RegisterReq struct {
	g.Meta   `path:"/account" method:"post" tags:"帐号" summary:"注册帐号"`
	Username string `json:"username" v:"required|length:2,12"`
	Password string `json:"password" v:"required|length:6,16" dc:"密码"`
	Email    string `json:"email" v:"required|email"`
}
type RegisterRes struct{}

type LoginReq struct {
	g.Meta   `path:"/login" method:"post" tags:"帐号" summary:"用户登录"`
	Username string `json:"username" v:"required|length:2,12"`
	Password string `json:"password" v:"required|length:6,16"`
}
type LoginRes struct {
	Token string `json:"token" dc:"在需要鉴权的接口中header加入Authorization: token"`
}
