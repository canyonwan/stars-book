package users

import (
	"context"
	"time"

	"gf-playground-1/internal/consts"
	"gf-playground-1/internal/dao"
	"gf-playground-1/internal/model/entity"

	"github.com/golang-jwt/jwt/v5"

	"github.com/gogf/gf/v2/errors/gerror"
)

type jwtClaims struct {
	Id       string
	UserName string
	jwt.RegisteredClaims
}

func (u *Users) Login(ctx context.Context, username, password string) (tokenString string, err error) {
	// 存放查询结果
	var userEntity *entity.Users
	err = dao.Users.Ctx(ctx).Where(dao.Users.Columns().Username, username).Scan(&userEntity)
	if err != nil {
		return "", gerror.New("用户名或密码错误")
	}
	if userEntity == nil {
		return "", gerror.New("用户不存在")
	}
	if userEntity.Password != u.encryptPassword(password) {
		return "", gerror.New("用户名或密码错误")
	}

	//生成token
	uc := jwtClaims{
		Id:       userEntity.Id,
		UserName: userEntity.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(6 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)

	return token.SignedString([]byte(consts.JwtKey))
}
