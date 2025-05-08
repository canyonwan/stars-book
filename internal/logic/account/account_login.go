package account

import (
	"context"
	"time"

	"gf-playground-1/internal/consts"
	"gf-playground-1/internal/dao"
	"gf-playground-1/internal/model"
	"gf-playground-1/internal/model/entity"
	"gf-playground-1/utility"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/golang-jwt/jwt/v5"
)

func (a *Account) Login(ctx context.Context, username, password string) (tokenString string, err error) {
	var userEntity *entity.Users
	err = dao.Users.Ctx(ctx).
		Where(dao.Users.Columns().Username, username).
		Scan(&userEntity)
	if err != nil {
		return "", err
	}
	if userEntity == nil {
		return "", gerror.New("帐号不存在,请先注册!")
	}
	if userEntity.Password != utility.EncryptPassword(password) {
		return "", gerror.New("帐号或密码错误!")
	}

	uc := model.JwtClaimsInput{
		UserId:   userEntity.Id,
		UserName: userEntity.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(6 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err = token.SignedString([]byte(consts.JwtKey))
	return
}
