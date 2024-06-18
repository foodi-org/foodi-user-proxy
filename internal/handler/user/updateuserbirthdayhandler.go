package user

import (
	"github.com/foodi-org/foodi-user-proxy/common/response"
	"net/http"

	"github.com/foodi-org/foodi-user-proxy/internal/logic/user"
	"github.com/foodi-org/foodi-user-proxy/internal/svc"
	"github.com/foodi-org/foodi-user-proxy/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 更新用户生日
func UpdateUserBirthdayHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateBirthdayRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.BindFailResponse(w, err) // 统一错误返回, todo 添加到template
			return
		}

		l := user.NewUpdateUserBirthdayLogic(r.Context(), svcCtx)
		err := l.UpdateUserBirthday(&req)
		response.Response(w, nil, err) // 统一错误返回, todo 添加到template
	}
}
