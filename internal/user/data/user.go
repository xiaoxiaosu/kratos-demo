package data

import (
	"common/internal/pkg/data"
	"context"

	"common/internal/user/biz"
)

type UserRepo struct {
	data *data.Data
}

func NewUserRepo(data *data.Data) biz.UserRepo {
	return &UserRepo{
		data: data,
	}
}

func (u *UserRepo) Save(ctx context.Context, user *biz.User) (bool, error) {
	u.data.Db.Exec("")
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
