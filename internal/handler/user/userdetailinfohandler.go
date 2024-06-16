package user

import (
	"net/http"

	"github.com/foodi-org/foodi-user-proxy/internal/logic/user"
	"github.com/foodi-org/foodi-user-proxy/internal/svc"
	"github.com/foodi-org/foodi-user-proxy/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取用户详细信息
func UserDetailInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserBaseInfoRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewUserDetailInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserDetailInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
