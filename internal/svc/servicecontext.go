package svc

import (
	"github.com/foodi-org/foodi-user-proxy/internal/config"
	"github.com/foodi-org/foodi-user-service/client/account"
	"github.com/foodi-org/foodi-user-service/client/article"
	"github.com/foodi-org/foodi-user-service/client/user"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

var svc = ServiceContext{}

type ServiceContext struct {
	Config        *config.Config
	UserClient    user.User
	AccountClient account.Account
	ArticleClient article.Article
}

func SVC() *ServiceContext {
	return &svc
}

func NewServiceContext(c *config.Config, path string, filename string) error {

	// 加载配置
	if err := config.InitServConf(path, filename); err != nil {
		return err
	} else {
		svc.Config = c
	}

	// 日志配置
	logx.MustSetup(logx.LogConf{
		ServiceName: c.Name,
		Mode:        "file",
		//TimeFormat:  "",
		Path:       path,
		Level:      "info",
		KeepDays:   10,
		MaxBackups: 10,
		Rotation:   "daily",
	})
	logx.DisableStat() // disables the stat logs.

	svc.UserClient = user.NewUser(zrpc.MustNewClient(c.UserRpc))
	svc.AccountClient = account.NewAccount(zrpc.MustNewClient(c.UserRpc))
	svc.ArticleClient = article.NewArticle(zrpc.MustNewClient(c.UserRpc))

	return nil
}
