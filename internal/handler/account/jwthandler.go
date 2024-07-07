package account

import (
	"github.com/foodi-org/foodi-user-proxy/common/response"
	"net/http"

	"github.com/foodi-org/foodi-user-proxy/internal/logic/account"
	"github.com/foodi-org/foodi-user-proxy/internal/svc"
	"github.com/foodi-org/foodi-user-proxy/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// get user token
func JwtHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.JwtTokenRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := account.NewJwtLogic(r.Context(), svcCtx)
		resp, err := l.Jwt(&req)
		response.Response(w, resp, err)
	}
}
