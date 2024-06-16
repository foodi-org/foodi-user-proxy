package user

import (
	"net/http"

	"github.com/foodi-org/foodi-user-proxy/internal/logic/user"
	"github.com/foodi-org/foodi-user-proxy/internal/svc"
	"github.com/foodi-org/foodi-user-proxy/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 更新用户头像
func UpdateUserImageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateImageRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewUpdateUserImageLogic(r.Context(), svcCtx)
		resp, err := l.UpdateUserImage(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
