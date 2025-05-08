package users

import (
	"context"

	v1 "gf-playground-1/api/users/v1"
	"gf-playground-1/internal/logic/middleware"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (c *ControllerV1) UserList(ctx context.Context, req *v1.UserListReq) (res *v1.UserListRes, err error) {
	r := ghttp.RequestFromCtx(ctx)
	if r == nil {
		return
	}

	list, err := c.users.UserList(ctx, req)
	if err != nil {
		middleware.ResponseFail(r, err)
		return
	}
	middleware.ResponseSuccess(r, list)
	return
}
