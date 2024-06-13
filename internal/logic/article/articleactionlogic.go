package article

import (
	"context"

	"github.com/foodi-org/foodi-user-proxy/internal/svc"
	"github.com/foodi-org/foodi-user-proxy/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ArticleActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 发布文章
func NewArticleActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleActionLogic {
	return &ArticleActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ArticleActionLogic) ArticleAction(req *types.AddRequest) (resp *types.AddReply, err error) {
	// todo: add your logic here and delete this line

	return
}
