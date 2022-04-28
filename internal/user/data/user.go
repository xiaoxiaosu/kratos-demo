package data

import (
	"context"
	"demo/internal/pkg"

	"demo/internal/user/biz"
)

type UserRepo struct {
	data *pkg.Data
}

func NewUserRepo(data *pkg.Data) biz.UserRepo {
	return &UserRepo{
		data: data,
	}
}

func (m *UserRepo) Save(ctx context.Context, user *biz.User) (bool, error) {
	m.data.Db.Exec("")
	return false, nil
}

func (m *UserRepo) Update(ctx context.Context, user *biz.User) (bool, error) {
	m.data.Db.Exec("")
	return false, nil
}

func (m *UserRepo) Get(ctx context.Context, i int64) (*biz.User, error) {
	m.data.RedisCli.Get(ctx, "")
	return nil, nil
}

func (m *UserRepo) ListAll(ctx context.Context) ([]*biz.User, error) {
	m.data.RedisCli.HGetAll(ctx, "")
	return nil, nil
}
