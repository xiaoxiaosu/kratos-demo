package main

import (
	"demo/internal/conf"
	"demo/internal/pkg/task"
	"github.com/RichardKnop/machinery/v1"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
)

var (
	taskConf   conf.Task
	taskServer *machinery.Server
)

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

	if err := c.Value("task").Scan(&taskConf); err != nil {
		panic(err)
	}

	taskServer, err := task.NewTaskServer(&taskConf)
	if err != nil {
		panic(err)
	}

	taskServer.NewWorker("audit_worker", 0)
}
