package account

import (
	"context"
	"github.com/foodi-org/foodi-user-proxy/common/xerror"
	foodi_user_service "github.com/foodi-org/foodi-user-service/github.com/foodi-org/foodi-user-service"

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

	switch foodi_user_service.RegisterCoup(req.RegisterType) {
	case foodi_user_service.RegisterCoup_Phone:
		// todo 接入一键注册验证 code
		var check bool
		if !check {
			return nil, err
		}
		reply, err := l.svcCtx.AccountClient.Login(l.ctx, &foodi_user_service.LoginRequest{
			Type:      foodi_user_service.UserCoup(req.AccountType),
			LoginType: foodi_user_service.RegisterCoup_Phone,
			Phone:     req.Phone,
			Code:      req.Code,
			Password:  req.Password,
			Length:    int64(req.Length),
		})
		if err != nil {
			return nil, err
		}
		resp.Token = reply.GetToken()
		return resp, nil
	case foodi_user_service.RegisterCoup_Password:
		reply, err := l.svcCtx.AccountClient.Login(l.ctx, &foodi_user_service.LoginRequest{
			Type:      foodi_user_service.UserCoup(req.AccountType),
			LoginType: foodi_user_service.RegisterCoup_Phone,
			Phone:     req.Phone,
			Code:      req.Code,
			Password:  req.Password,
			Length:    int64(req.Length),
		})
		if err != nil {
			return nil, err
		}
		resp.Token = reply.GetToken()
		return resp, nil
	case foodi_user_service.RegisterCoup_Code:
		// todo 接入验证码验证 code
		var check bool
		if !check {
			return nil, err
		}
		reply, err := l.svcCtx.AccountClient.Login(l.ctx, &foodi_user_service.LoginRequest{
			Type:      foodi_user_service.UserCoup(req.AccountType),
			LoginType: foodi_user_service.RegisterCoup_Phone,
			Phone:     req.Phone,
			Code:      req.Code,
			Password:  req.Password,
			Length:    int64(req.Length),
		})
		if err != nil {
			return nil, err
		}
		resp.Token = reply.GetToken()
		return resp, nil
	default:
		return nil, xerror.InvalidRegisterType
	}
}
