package service

import (
	pb "common/api/merchant"
	"common/internal/merchant/biz"
	"context"
	"net/http"
)

func (m *MerchantService) CreateMerchant(ctx context.Context, in *pb.CreateMerchantRequest) (*pb.CreateMerchantReply, error) {

	// 定义返回值
	reply := &pb.CreateMerchantReply{
		Code:    0,
		Message: "succ",
	}

	// 调用biz层执行业务逻辑
	_, err := m.mc.CreateMerchant(ctx, &biz.Merchant{Name: in.Name})
	if err != nil {
		reply.Code = http.StatusInternalServerError
		reply.Message = err.Error()
	}

	return reply, nil
}

func (m *MerchantService) ListMerchant(ctx context.Context, in *pb.ListMerchantRequest) (*pb.ListMerchantReply, error) {
	merchants := []*pb.Merchant{}

	merchants = append(merchants, &pb.Merchant{Name: "test1"})
	merchants = append(merchants, &pb.Merchant{Name: "test2"})

	reply := &pb.ListMerchantReply{Merchant: merchants}

	return reply, nil
}
