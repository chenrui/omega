/*
 * @Autor: Ruic
 * @since: 2021-09-13 14:41:23
 * @LastEditTime: 2021-09-17 15:17:31
 */
package data

import (
	"context"
	"omega/app/account/internal/biz"
	"omega/app/account/internal/data/model"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/user")),
	}
}

func (r *userRepo) CreateUser(ctx context.Context, u *biz.User) error {
	user := model.User{
		Name:     u.Username,
		Password: u.Password,
	}
	return r.data.db.Create(&user).Error
}

func (r *userRepo) UpdateGreeter(ctx context.Context, u *biz.User) error {
	return nil
}
