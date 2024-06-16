package user

import (
	"context"
	"github.com/zeromicro/x/errors"

	"github.com/foodi-org/foodi-user-proxy/internal/svc"
	"github.com/foodi-org/foodi-user-proxy/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserBirthdayLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新用户生日
func NewUpdateUserBirthdayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserBirthdayLogic {
	return &UpdateUserBirthdayLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserBirthdayLogic) UpdateUserBirthday(req *types.UpdateBirthdayRequest) error {
	// todo: add your logic here and delete this line

	return errors.New(100, "aaa")
}
