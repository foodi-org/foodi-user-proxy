package login

import (
	"context"
	"github.com/foodi-org/foodi-user-proxy/internal/handler/pkg/AES"

	"github.com/foodi-org/foodi-user-proxy/internal/svc"
	"github.com/foodi-org/foodi-user-proxy/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// user login
func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginReply, err error) {
	// todo: add your logic here and delete this line

	var password string
	if req.Password != "" {
		password, err = AES.DecryptAES(req.Password, req.Length)
		if err != nil {
			return nil, err
		}
	}

	return
}
