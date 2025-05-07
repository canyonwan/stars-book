package users

import (
	"context"

	v1 "gf-playground-1/api/users/v1"
	"gf-playground-1/internal/dao"
	"gf-playground-1/internal/model/do"

	"github.com/bwmarrin/snowflake"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (u *Users) Register(ctx context.Context, req *v1.RegisterReq) error {
	// 先检查用户是否已注册过
	if err := u.checkUser(ctx, req.Username); err != nil {
		return err
	}

	node, err := snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}

	id := node.Generate()
	_, err = dao.Users.Ctx(ctx).Data(do.Users{
		Id:       id,
		Username: req.Username,
		Password: u.encryptPassword(req.Password),
		Email:    req.Email,
	}).Insert()
	if err != nil {
		return err
	}
	return nil
}

func (u *Users) checkUser(ctx context.Context, username string) error {
	count, err := dao.Users.Ctx(ctx).Where(dao.Users.Columns().Username, username).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.New("用户已存在")
	}
	return nil
}
