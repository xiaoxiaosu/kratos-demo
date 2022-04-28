package service

import (
	"demo/api/user"
	"demo/internal/pkg"
	"demo/internal/user/biz"
	"demo/internal/user/data"
)

type UserService struct {
	user.UnimplementedUserServer

	uc *biz.UserUseCase
}

func NewUserService(d *pkg.Data) *UserService {
	UserRepo := data.NewUserRepo(d)
	UserUseCase := biz.NewUserUseCase(UserRepo)

	return &UserService{
		uc: UserUseCase,
	}
}
