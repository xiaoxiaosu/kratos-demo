package biz

import (
	"context"
)

type User struct {
	Name string
}

type UserRepo interface {
	Save(context.Context, *User) (bool, error)
	Update(context.Context, *User) (bool, error)
	Get(context.Context, int64) (*User, error)
	ListAll(context.Context) ([]*User, error)
}

type UserUseCase struct {
	repo UserRepo
}

func NewUserUseCase(repo UserRepo) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (uc *UserUseCase) CreateUser(ctx context.Context, u *User) (bool, error) {
	return uc.repo.Save(ctx, u)
}

func (uc *UserUseCase) UpdateUser(ctx context.Context, m *User) (bool, error) {
	return uc.repo.Update(ctx, m)
}

func (uc *UserUseCase) GetUser(ctx context.Context, id int64) (*User, error) {
	return uc.repo.Get(ctx, id)
}

func (uc *UserUseCase) ListUser(ctx context.Context) ([]*User, error) {
	return uc.repo.ListAll(ctx)
}
