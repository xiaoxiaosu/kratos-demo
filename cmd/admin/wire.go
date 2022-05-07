package main

import (
	"common/internal/conf"
	merchantBiz "common/internal/merchant/biz"
	merchantData "common/internal/merchant/data"
	merchantService "common/internal/merchant/service"
	userBiz "common/internal/user/biz"
	userData "common/internal/user/data"
	userService "common/internal/user/service"

	"common/internal/pkg/data"
	"common/internal/pkg/server"
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
)

func genApp(*conf.Server, *conf.Data) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, merchantService.ProviderSet, merchantBiz.ProviderSet, data.ProviderSet, merchantData.ProviderSet, userService.ProviderSet, userBiz.ProviderSet, userData.ProviderSet, newApp))
}
