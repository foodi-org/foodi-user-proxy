package account

import (
	"net/http"

	"github.com/foodi-org/foodi-user-proxy/internal/logic/account"
	"github.com/foodi-org/foodi-user-proxy/internal/svc"
	"github.com/foodi-org/foodi-user-proxy/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// RegisterHandler user register
func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := account.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
