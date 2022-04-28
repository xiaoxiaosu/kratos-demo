package service

import (
	"demo/api/merchant"
	"demo/internal/merchant/biz"
	"demo/internal/merchant/data"
	"demo/internal/pkg"
)

type MerchantService struct {
	merchant.UnimplementedMerchantServer

	mc *biz.MerchantUseCase
	bc *biz.BusinessLineUseCase
}

func NewMerchantService(d *pkg.Data) *MerchantService {
	merchantRepo := data.NewMerchantRepo(d)
	merchantUseCase := biz.NewMerchantUseCase(merchantRepo)

	businessLineRepo := data.NewBusinessLineRepo(d)
	businessLineUseCase := biz.NewBusinessLineUseCase(businessLineRepo)

	return &MerchantService{
		mc: merchantUseCase,
		bc: businessLineUseCase,
	}
}
