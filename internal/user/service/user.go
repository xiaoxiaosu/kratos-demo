package service

import (
	"context"
	pb "demo/api/user"
	"demo/internal/user/biz"
	"net/http"
)

func (u *UserService) CreateMerchant(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	// 定义返回值
	reply := &pb.CreateUserReply{
		Code:    0,
		Message: "succ",
	}

	// 调用biz层执行业务逻辑
	_, err := u.uc.CreateUser(ctx, &biz.User{Name: in.Name})
	if err != nil {
		reply.Code = http.StatusInternalServerError
		reply.Message = err.Error()
	}

	return reply, nil
}

func (m *UserService) ListUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserReply, error) {

	return nil, nil
}
