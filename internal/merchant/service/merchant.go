package service

import (
	"context"
	pb "demo/api/merchant"
	"demo/internal/merchant/biz"
	"demo/internal/pkg/middleware/localize"
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

	merchants = append(merchants, &pb.Merchant{Name: "test1", Status: "blocked"})
	merchants = append(merchants, &pb.Merchant{Name: "test2", Status: "normal"})

	for _, v := range merchants {
		v.Status, _ = localize.Translate(ctx, "merchant_status", v.Status)
	}

	reply := &pb.ListMerchantReply{Data: merchants}

	return reply, nil
}
