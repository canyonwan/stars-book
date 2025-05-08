package account

import (
	"context"

	"gf-playground-1/internal/dao"
	"gf-playground-1/internal/model"
	"gf-playground-1/internal/model/do"
	"gf-playground-1/utility"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (a *Account) Register(ctx context.Context, in model.RegisterInput) error {
	err := checkAccountExist(ctx, in.Username)
	if err != nil {
		return err
	}

	_, err = dao.Users.Ctx(ctx).Data(do.Users{
		Id:       utility.GenerateSnowId(),
		Username: in.Username,
		Password: in.Password,
		Email:    in.Email,
	}).Insert()
	if err != nil {
		return err
	}
	return nil
}

func checkAccountExist(ctx context.Context, username string) error {
	count, err := dao.Users.Ctx(ctx).Where(dao.Users.Columns().Username, username).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.New("用户已存在")
	}
	return nil
}
