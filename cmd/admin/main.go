package main

import (
	"demo/internal/conf"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

var (
	appConf    conf.App
	dataConf   conf.Data
	serverConf conf.Server
)

func newApp(hs *http.Server, gs *grpc.Server) *kratos.App {
	return kratos.New(
		kratos.Name(appConf.Name),
		kratos.Metadata(map[string]string{}),
		kratos.Server(
			hs,
			gs,
		),
	)
}

func main() {

	// 初始化配置
	c := config.New(
		config.WithSource(
			file.NewSource("../../configs/config.yaml"),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	if err := c.Value("app").Scan(&appConf); err != nil {
		panic(err)
	}
	if err := c.Value("data").Scan(&dataConf); err != nil {
		panic(err)
	}

	var serverConf conf.Server
	if err := c.Value("server").Scan(&serverConf); err != nil {
		panic(err)
	}

	app, cleanup, err := wireApp(&serverConf, &dataConf)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	if err = app.Run(); err != nil {
		panic(err)
	}
}
