package main

import (
	"demo/internal/conf"
	merchantBiz "demo/internal/merchant/biz"
	merchantData "demo/internal/merchant/data"
	merchantService "demo/internal/merchant/service"
	userBiz "demo/internal/user/biz"
	userData "demo/internal/user/data"
	userService "demo/internal/user/service"

	"demo/internal/pkg/data"
	"demo/internal/pkg/server"
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
)

func genApp(*conf.Server, *conf.Data) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, merchantService.ProviderSet, merchantBiz.ProviderSet, data.ProviderSet, merchantData.ProviderSet, userService.ProviderSet, userBiz.ProviderSet, userData.ProviderSet, newApp))
}
