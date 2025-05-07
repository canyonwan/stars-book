package users

import (
	"context"

	"gf-playground-1/api/users/v1"
)

func (c *ControllerV1) DeleteUser(ctx context.Context, req *v1.DeleteUserReq) (res *v1.DeleteUserRes, err error) {
	err = c.users.DeleteUser(ctx, req.Id)
	return
}
