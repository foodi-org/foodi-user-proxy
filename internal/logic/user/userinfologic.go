package user

import (
	"context"
	"github.com/zeromicro/x/errors"

	"github.com/foodi-org/foodi-user-proxy/internal/svc"
	"github.com/foodi-org/foodi-user-proxy/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户基础信息
func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserBaseInfoRequest) (resp *types.UserBaseInfoReply, err error) {
	// todo: add your logic here and delete this line
	if req.UID == 100 {
		return &types.UserBaseInfoReply{
			UID:   1080,
			Name:  "ODIN",
			Image: "https://i.imgur.com/7Y.png",
		}, errors.New(40023, "jfjjs")
	}

	return &types.UserBaseInfoReply{
		UID:   1080,
		Name:  "ODIN",
		Image: "https://i.imgur.com/7Y.png",
	}, nil

}
