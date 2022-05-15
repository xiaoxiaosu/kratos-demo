package server

import (
	merchantProto "demo/api/merchant"
	userProto "demo/api/user"
	"demo/internal/conf"
	merchantService "demo/internal/merchant/service"
	"demo/internal/pkg/middleware/localize"
	userService "demo/internal/user/service"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func NewHTTPServer(c *conf.Server, merchant *merchantService.MerchantService, user *userService.UserService) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			localize.I18N(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}

	srv := http.NewServer(opts...)
	merchantProto.RegisterMerchantHTTPServer(srv, merchant)
	userProto.RegisterUserHTTPServer(srv, user)

	return srv
}
