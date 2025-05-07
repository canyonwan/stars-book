package users

import (
	"context"

	v1 "gf-playground-1/api/users/v1"
)

func (c *ControllerV1) UserInfo(ctx context.Context, req *v1.UserInfoReq) (res *v1.UserInfoRes, err error) {
	info, err := c.users.UserInfo(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.UserInfoRes{
		Id:        info.Id,
		Username:  info.Username,
		Email:     info.Email,
		CreatedAt: info.CreatedAt,
		UpdatedAt: info.UpdatedAt,
	}, nil
}
