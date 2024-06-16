package user

import (
	"context"

	"github.com/foodi-org/foodi-user-proxy/internal/svc"
	"github.com/foodi-org/foodi-user-proxy/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdayeUserNikenameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新用户昵称
func NewUpdayeUserNikenameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdayeUserNikenameLogic {
	return &UpdayeUserNikenameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdayeUserNikenameLogic) UpdayeUserNikename(req *types.UpdateNikenameRequest) (resp *types.Base, err error) {
	// todo: add your logic here and delete this line

	return
}
