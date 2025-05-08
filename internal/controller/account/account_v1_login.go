package account

import (
	"context"

	"gf-playground-1/api/account/v1"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	tokenString, err := c.account.Login(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	return &v1.LoginRes{Token: tokenString}, nil
}
