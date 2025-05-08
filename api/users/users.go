// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package users

import (
	"context"

	"gf-playground-1/api/users/v1"
)

type IUsersV1 interface {
	UserInfo(ctx context.Context, req *v1.UserInfoReq) (res *v1.UserInfoRes, err error)
	UserList(ctx context.Context, req *v1.UserListReq) (res *v1.UserListRes, err error)
	DeleteUser(ctx context.Context, req *v1.DeleteUserReq) (res *v1.DeleteUserRes, err error)
}
