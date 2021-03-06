package main

import (
	"flag"
	"fmt"
	"qiyaowu-go-zero/basic"

	"qiyaowu-go-zero/user/api/internal/config"
	"qiyaowu-go-zero/user/api/internal/handler"
	"qiyaowu-go-zero/user/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	basic.Init()

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
