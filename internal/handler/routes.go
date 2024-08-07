// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	account "github.com/foodi-org/foodi-user-proxy/internal/handler/account"
	article "github.com/foodi-org/foodi-user-proxy/internal/handler/article"
	user "github.com/foodi-org/foodi-user-proxy/internal/handler/user"
	"github.com/foodi-org/foodi-user-proxy/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				// user login
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: account.LoginHandler(serverCtx),
			},
			{
				// user register
				Method:  http.MethodPost,
				Path:    "/path",
				Handler: account.RegisterHandler(serverCtx),
			},
			{
				// get user token
				Method:  http.MethodPost,
				Path:    "/user/token",
				Handler: account.JwtHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 发布文章
				Method:  http.MethodPost,
				Path:    "/add",
				Handler: article.ArticleActionHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 获取用户基础信息
				Method:  http.MethodGet,
				Path:    "/user/baseInfo",
				Handler: user.UserInfoHandler(serverCtx),
			},
			{
				// 更新用户生日
				Method:  http.MethodPut,
				Path:    "/user/birthday",
				Handler: user.UpdateUserBirthdayHandler(serverCtx),
			},
			{
				// 获取用户详细信息
				Method:  http.MethodGet,
				Path:    "/user/detailInfo",
				Handler: user.UserDetailInfoHandler(serverCtx),
			},
			{
				// 更新用户性别
				Method:  http.MethodPut,
				Path:    "/user/gender",
				Handler: user.UpdateUserGenderHandler(serverCtx),
			},
			{
				// 更新用户昵称
				Method:  http.MethodPut,
				Path:    "/user/nikename",
				Handler: user.UpdayeUserNikenameHandler(serverCtx),
			},
			{
				// 更新用户所在地域
				Method:  http.MethodPut,
				Path:    "/user/region",
				Handler: user.UpdateUserRegionHandler(serverCtx),
			},
			{
				// 更新用户头像
				Method:  http.MethodPut,
				Path:    "/user/updateImage",
				Handler: user.UpdateUserImageHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/v1"),
	)
}
