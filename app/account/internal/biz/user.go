/*
 * @Autor: Ruic
 * @since: 2021-09-13 14:41:23
 * @LastEditTime: 2021-09-17 15:25:45
 */
package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Id       int64
	Username string
	Password string
}

type UserRepo interface {
	CreateUser(context.Context, *User) error
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/interface"))}
}

func (uc *UserUsecase) CreateUser(ctx context.Context, u *User) error {
	return uc.repo.CreateUser(ctx, u)
}
