package data

import (
	"context"
	"demo/internal/merchant/biz"
	"demo/internal/pkg/data"
)

type BusinessLineRepo struct {
	data *data.Data
}

func NewBusinessLineRepo(data *data.Data) biz.BusinessLineRepo {
	return &BusinessLineRepo{
		data: data,
	}
}

func (b BusinessLineRepo) Save(ctx context.Context, line *biz.BusinessLine) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (b BusinessLineRepo) Update(ctx context.Context, line *biz.BusinessLine) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (b BusinessLineRepo) Get(ctx context.Context, i int64) (*biz.BusinessLine, error) {
	//TODO implement me
	panic("implement me")
}

func (b BusinessLineRepo) ListAll(ctx context.Context) ([]*biz.BusinessLine, error) {
	//TODO implement me
	panic("implement me")
}
