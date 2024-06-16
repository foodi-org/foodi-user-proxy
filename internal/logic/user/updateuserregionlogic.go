package user

import (
	"context"

	"github.com/foodi-org/foodi-user-proxy/internal/svc"
	"github.com/foodi-org/foodi-user-proxy/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserRegionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新用户所在地域
func NewUpdateUserRegionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserRegionLogic {
	return &UpdateUserRegionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserRegionLogic) UpdateUserRegion(req *types.UpdateRegionRequest) (resp *types.Base, err error) {
	// todo: add your logic here and delete this line

	return
}
