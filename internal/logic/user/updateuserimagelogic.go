package user

import (
	"context"

	"github.com/foodi-org/foodi-user-proxy/internal/svc"
	"github.com/foodi-org/foodi-user-proxy/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserImageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新用户头像
func NewUpdateUserImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserImageLogic {
	return &UpdateUserImageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserImageLogic) UpdateUserImage(req *types.UpdateImageRequest) (resp *types.Base, err error) {
	// todo: add your logic here and delete this line

	return
}
