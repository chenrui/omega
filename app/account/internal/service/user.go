/*
 * @Autor: Ruic
 * @since: 2021-09-13 14:41:23
 * @LastEditTime: 2021-09-17 15:30:04
 */
package service

import (
	"context"
	v1 "omega/api/account/service/v1"
	"omega/app/account/internal/biz"
)

func (s *AccountService) CreateUser(ctx context.Context, req *v1.CreateUserReq) (*v1.CreateUserReply, error) {
	u := biz.User{
		Username: req.GetUsername(),
		Password: req.GetPassword(),
	}
	s.uc.CreateUser(ctx, &u)
	return nil, nil
}

func (s *AccountService) GetUser(ctx context.Context, req *v1.GetUserReq) (*v1.GetUserReply, error) {
	return nil, nil
}

func (s *AccountService) VerifyPassword(ctx context.Context, req *v1.VerifyPasswordReq) (*v1.VerifyPasswordReply, error) {
	return nil, nil
}
