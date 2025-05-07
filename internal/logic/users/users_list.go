package users

import (
	"context"

	v1 "gf-playground-1/api/users/v1"
	"gf-playground-1/internal/dao"
	"gf-playground-1/internal/model/entity"
)

func (u *Users) UserList(ctx context.Context, req *v1.UserListReq) ([]entity.Users, error) {
	var list []entity.Users

	username := "%" + req.Username + "%"
	err := dao.Users.Ctx(ctx).
		WhereLike(dao.Users.Columns().Username, username).
		WhereNull(dao.Users.Columns().DeletedAt).
		Scan(&list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
