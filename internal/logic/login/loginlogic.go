package login

import (
	"context"
	"fmt"
	"github.com/foodi-org/foodi-user-service/client/account"
	foodi_user_service "github.com/foodi-org/foodi-user-service/pb/github.com/foodi-org/foodi-user-service"

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
	//if req.Password != "" {
	//	password, err = AES.DecryptAES(req.Password, req.Length)
	//	if err != nil {
	//		return nil, err
	//	}
	//}
	fmt.Println(password)
	reply, err := l.svcCtx.AccountClient.Login(l.ctx, &account.LoginRequest{
		Phone:     13360017279,
		Password:  password,
		LoginType: foodi_user_service.RegisterCoup_Password,
	})
	if err != nil {
		return nil, err
	}

	return &types.LoginReply{
		Token: reply.GetToken(),
	}, nil
}
