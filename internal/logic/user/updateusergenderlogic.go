package user

import (
	"context"

	"github.com/foodi-org/foodi-user-proxy/internal/svc"
	"github.com/foodi-org/foodi-user-proxy/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserGenderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新用户性别
func NewUpdateUserGenderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserGenderLogic {
	return &UpdateUserGenderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserGenderLogic) UpdateUserGender(req *types.UpdateGenderRequest) (resp *types.Base, err error) {
	// todo: add your logic here and delete this line

	return
}
