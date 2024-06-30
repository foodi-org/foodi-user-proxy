package main

import (
	"flag"
	"fmt"
	"github.com/foodi-org/foodi-user-proxy/internal/config"
	"github.com/foodi-org/foodi-user-proxy/internal/handler"
	"github.com/foodi-org/foodi-user-proxy/internal/svc"
	"github.com/zeromicro/go-zero/rest"
	_ "github.com/zeromicro/zero-contrib/zrpc/registry/consul"
	"os"
)

var configFile = flag.String("f", "dev.yaml", "the config file")

func main() {
	var c = config.ServConf()
	dir, _ := os.Getwd()
	flag.Parse()
	if err := svc.NewServiceContext(c, dir, *configFile); err != nil {
		panic(err)
	}
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, svc.SVC())

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
