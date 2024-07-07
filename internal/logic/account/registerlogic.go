package account

import (
	"context"
	"fmt"
	"github.com/foodi-org/foodi-user-proxy/common/xerror"
	"github.com/foodi-org/foodi-user-proxy/internal/svc"
	"github.com/foodi-org/foodi-user-proxy/internal/types"
	foodiuserservice "github.com/foodi-org/foodi-user-service/github.com/foodi-org/foodi-user-service"
	"regexp"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewRegisterLogic user register
func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// mobileCheck
//
//	@Description: 手机号码校验
//	@param phone 手机号码
//	@return bool
func (r *RegisterLogic) mobileCheck(phone int64) bool {
	str := fmt.Sprintf("%d", phone)

	// 匹配规则
	// ^1 第一位为1
	// [345789]{1} 第二位位 345789
	// 后接9位数字并且结束
	regRuler := "^1[345789]{1}\\d{9}$"

	reg := regexp.MustCompile(regRuler)

	return reg.MatchString(str)
}

func (r *RegisterLogic) idCardCheck(card string) bool {
	//18位身份证 ^(\d{17})([0-9]|X)$
	// 匹配规则
	// (^\d{15}$) 15位身份证
	// (^\d{18}$) 18位身份证
	// (^\d{17}(\d|X|x)$) 18位身份证 最后一位为X的用户
	regRuler := "(^\\d{15}$)|(^\\d{18}$)|(^\\d{17}(\\d|X|x)$)"

	// 正则调用规则
	reg := regexp.MustCompile(regRuler)

	// 返回 MatchString 是否匹配
	return reg.MatchString(card)
}

func (r *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterReply, err error) {
	var (
		reply *foodiuserservice.RegisterReply
	)
	if req.Password == "" {
		return nil, xerror.MissPhone
	}
	if !r.mobileCheck(req.Phone) {
		return nil, xerror.InvalidPhone
	}

	switch foodiuserservice.RegisterCoup(req.RegisterType) {
	case foodiuserservice.RegisterCoup_Phone:
		if req.Code == "" {
			return nil, xerror.MissCode
		}
		if reply, err = r.svcCtx.AccountClient.Register(r.ctx, &foodiuserservice.RegisterRequest{
			Type:         foodiuserservice.UserCoup(req.UserType),
			RegisterType: foodiuserservice.RegisterCoup_Phone,
			Phone:        req.Phone,
			Code:         req.Code,
		}); err != nil {
			goto RPCError
		} else if reply.Ok {
			goto RegisterSuccess
		} else {
			goto RegisterFail
		}
	case foodiuserservice.RegisterCoup_Code:
		if req.Code == "" {
			return nil, xerror.MissCode
		}
		if reply, err = r.svcCtx.AccountClient.Register(r.ctx, &foodiuserservice.RegisterRequest{
			Type:         foodiuserservice.UserCoup(req.UserType),
			RegisterType: foodiuserservice.RegisterCoup_Code,
			Phone:        req.Phone,
			Code:         req.Code,
		}); err != nil {
			goto RPCError
		} else if reply.Ok {
			goto RegisterSuccess
		} else {
			goto RegisterFail
		}
	case foodiuserservice.RegisterCoup_Password:
		if req.Password == "" {
			return nil, xerror.MissPassword
		}
		if reply, err = r.svcCtx.AccountClient.Register(r.ctx, &foodiuserservice.RegisterRequest{
			Type:         foodiuserservice.UserCoup(req.UserType),
			RegisterType: foodiuserservice.RegisterCoup_Password,
			Phone:        req.Phone,
			Password:     req.Password,
			Length:       0,
		}); err != nil {
			goto RPCError
		} else if reply.Ok {
			goto RegisterSuccess
		} else {
			goto RegisterFail
		}
	default:
		return nil, xerror.InvalidRegisterType
	}

RegisterFail:
	return nil, xerror.RegisterFail
RPCError:
	return nil, err
RegisterSuccess:
	// 生成token
	return &types.RegisterReply{
		Token: reply.GetToken(),
		Uid:   reply.GetUid(),
	}, nil
}
