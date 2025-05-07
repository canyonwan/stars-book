package users

import (
	"context"

	"gf-playground-1/internal/consts"
	"gf-playground-1/internal/dao"
	"gf-playground-1/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/golang-jwt/jwt/v5"
)

func (u *Users) UserInfo(ctx context.Context) (user *entity.Users, err error) {
	tokenString := g.RequestFromCtx(ctx).Request.Header.Get("Authorization")
	tokenClaims, _ := jwt.ParseWithClaims(tokenString, &jwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(consts.JwtKey), nil
	})

	if claims, ok := tokenClaims.Claims.(*jwtClaims); ok && tokenClaims.Valid {
		err = dao.Users.Ctx(ctx).Where(dao.Users.Columns().Id, claims.Id).Scan(&user)
	}
	return
}
