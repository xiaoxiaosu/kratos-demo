package main

import (
	"demo/api/merchant"
	"demo/api/user"
	"demo/internal/conf"
	ms "demo/internal/merchant/service"
	"demo/internal/pkg"
	us "demo/internal/user/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func main() {
	// 初始化配置
	c := config.New(
		config.WithSource(
			file.NewSource("../../configs/config.yaml"),
		),
	)
	if err := c.Load(); err != nil {
		panic(err)
	}

	var appConf conf.App
	if err := c.Value("app").Scan(&appConf); err != nil {
		panic(err)
	}
	var dataConf conf.Data
	if err := c.Value("data").Scan(&dataConf); err != nil {
		panic(err)
	}

	var serverConf conf.Server
	if err := c.Value("server").Scan(&serverConf); err != nil {
		panic(err)
	}

	// 初始化数据库
	data, err := pkg.NewData(&dataConf)
	if err != nil {
		panic(err)
	}

	// 创建http服务器
	hs := http.NewServer(http.Address(serverConf.Http.Addr))

	//注册服务
	merchant.RegisterMerchantHTTPServer(hs, ms.NewMerchantService(data))
	user.RegisterUserHTTPServer(hs, us.NewUserService(data))

	// 实例化应用
	app := kratos.New(
		kratos.Name(appConf.Name),
		kratos.Server(hs),
	)

	app.Run()
}
