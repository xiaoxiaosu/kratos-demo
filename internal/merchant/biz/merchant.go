package biz

import (
	"context"
)

type Merchant struct {
	Name string
}

type MerchantRepo interface {
	Save(context.Context, *Merchant) (bool, error)
	Update(context.Context, *Merchant) (bool, error)
	Get(context.Context, int64) (*Merchant, error)
	ListAll(context.Context) ([]*Merchant, error)
}

type MerchantUseCase struct {
	repo MerchantRepo
}

func NewMerchantUseCase(repo MerchantRepo) *MerchantUseCase {
	return &MerchantUseCase{repo: repo}
}

func (mc *MerchantUseCase) CreateMerchant(ctx context.Context, m *Merchant) (bool, error) {
	return mc.repo.Save(ctx, m)
}

func (mc *MerchantUseCase) UpdateMerchant(ctx context.Context, m *Merchant) (bool, error) {
	return mc.repo.Update(ctx, m)
}

func (mc *MerchantUseCase) GetMerchant(ctx context.Context, id int64) (*Merchant, error) {
	return mc.repo.Get(ctx, id)
}

func (mc *MerchantUseCase) ListMerchant(ctx context.Context) ([]*Merchant, error) {
	return mc.repo.ListAll(ctx)
}
