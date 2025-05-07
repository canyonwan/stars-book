package users

import (
	"context"

	"gf-playground-1/internal/dao"
)

func (u *Users) DeleteUser(ctx context.Context, id string) error {
	_, err := dao.Users.Ctx(ctx).WherePri(id).Delete()
	if err != nil {
		return err
	}
	return nil
}
