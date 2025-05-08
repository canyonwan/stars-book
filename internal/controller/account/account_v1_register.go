package account

import (
	"context"

	"gf-playground-1/api/account/v1"
	"gf-playground-1/internal/model"
	"gf-playground-1/utility"
)

func (c *ControllerV1) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {
	err = c.account.Register(ctx, model.RegisterInput{
		Username: req.Username,
		Password: utility.EncryptPassword(req.Password),
		Email:    req.Email,
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}
