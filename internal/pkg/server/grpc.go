package server

import (
	merchantProto "common/api/merchant"
	userProto "common/api/user"
	"common/internal/conf"
	merchantService "common/internal/merchant/service"
	userService "common/internal/user/service"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

func NewGRPCServer(c *conf.Server, merchant *merchantService.MerchantService, user *userService.UserService) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}

	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}

	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}

	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}

	srv := grpc.NewServer(opts...)
	merchantProto.RegisterMerchantServer(srv, merchant)
	userProto.RegisterUserServer(srv, user)
	return srv
}
