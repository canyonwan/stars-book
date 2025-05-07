package v1

import (
	"gf-playground-1/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Status marks users status.

type RegisterReq struct {
	g.Meta   `path:"/users/register" method:"post" tags:"用户" summary:"注册用户"`
	Username string `json:"username" v:"required|length:2,12"`
	Password string `json:"password" v:"required|length:6,16" dc:"密码"`
	Email    string `json:"email" v:"required|email"`
}
type RegisterRes struct{}

type LoginReq struct {
	g.Meta   `path:"/login" method:"post" tags:"用户" summary:"用户登录"`
	Username string `json:"username" v:"required|length:2,12"`
	Password string `json:"password" v:"required|length:6,16"`
}
type LoginRes struct {
	Token string `json:"token" dc:"在需要鉴权的接口中header加入Authorization: token"`
}

type UserInfoReq struct {
	g.Meta `path:"/user/info" method:"get" tags:"用户" summary:"获取用户信息"`
}
type UserInfoRes struct {
	Id        string      `json:"id"`
	Username  string      `json:"username"`
	Email     string      `json:"email"`
	CreatedAt *gtime.Time `json:"created_at"`
	UpdatedAt *gtime.Time `json:"updated_at"`
}

type UserListReq struct {
	g.Meta   `path:"/users" method:"get" tags:"用户" summary:"获取用户列表"`
	Username string `json:"username" dc:"用户名"`
}
type UserListRes struct {
	Users []entity.Users `json:"users"`
}

type DeleteUserReq struct {
	g.Meta `path:"/users/delete/:id" method:"delete" tags:"用户" summary:"删除用户"`
	Id     string `json:"id" v:"required"`
}
type DeleteUserRes struct{}
