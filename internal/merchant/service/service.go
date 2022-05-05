package service

import (
	"demo/api/merchant"
	"demo/internal/merchant/biz"
	"github.com/google/wire"
)

type MerchantService struct {
	merchant.UnimplementedMerchantServer

	mc *biz.MerchantUseCase
	bc *biz.BusinessLineUseCase
}

var ProviderSet = wire.NewSet(NewMerchantService)

func NewMerchantService(mc *biz.MerchantUseCase, bc *biz.BusinessLineUseCase) *MerchantService {

	return &MerchantService{
		mc: mc,
		bc: bc,
	}
}
