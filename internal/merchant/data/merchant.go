package data

import (
	"context"
	"demo/internal/pkg/data"

	"demo/internal/merchant/biz"
)

type MerchantRepo struct {
	data *data.Data
}

func NewMerchantRepo(data *data.Data) biz.MerchantRepo {
	return &MerchantRepo{
		data: data,
	}
}

func (m MerchantRepo) Save(ctx context.Context, merchant *biz.Merchant) (bool, error) {
	m.data.Db.Exec("")
	return false, nil
}

func (m MerchantRepo) Update(ctx context.Context, merchant *biz.Merchant) (bool, error) {
	m.data.Db.Exec("")
	return false, nil
}

func (m MerchantRepo) Get(ctx context.Context, i int64) (*biz.Merchant, error) {
	m.data.RedisCli.Get(ctx, "")
	return nil, nil
}

func (m MerchantRepo) ListAll(ctx context.Context) ([]*biz.Merchant, error) {
	m.data.RedisCli.HGetAll(ctx, "")
	return nil, nil
}
