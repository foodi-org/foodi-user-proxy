package user

import (
	"context"

	"github.com/foodi-org/foodi-user-proxy/internal/svc"
	"github.com/foodi-org/foodi-user-proxy/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetailInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户详细信息
func NewUserDetailInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailInfoLogic {
	return &UserDetailInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDetailInfoLogic) UserDetailInfo(req *types.UserBaseInfoRequest) (resp *types.UserDetailReply, err error) {
	// todo: add your logic here and delete this line

	return &types.UserDetailReply{
		UID:      1,
		Name:     "this is name",
		Nikename: "this is nike name",
		Gender:   "famle",
	}, nil
}
