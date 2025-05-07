package cmd

import (
	"context"
	"gf-playground-1/internal/controller/users"
	"gf-playground-1/internal/logic/middleware"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Group("/v1", func(group *ghttp.RouterGroup) {
					group.Bind(
						users.NewV1().Login,
						users.NewV1().Register,
					)
					group.Group("/", func(group *ghttp.RouterGroup) {
						group.Middleware(middleware.Auth)
						group.Bind(
							users.NewV1().UserInfo,
							users.NewV1().DeleteUser,
							users.NewV1().UserList,
						)
					})
				})
			})
			s.Run()
			return nil
		},
	}
)
