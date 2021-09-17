/*
 * @Autor: Ruic
 * @since: 2021-09-13 14:41:23
 * @LastEditTime: 2021-09-15 09:25:48
 */
package service

import (
	"omega/app/account/internal/biz"

	v1 "omega/api/account/service/v1"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewAccountService)

type AccountService struct {
	v1.UnimplementedAccountServer

	log *log.Helper
	uc  *biz.UserUsecase
}

func NewAccountService(uc *biz.UserUsecase, logger log.Logger) *AccountService {
	return &AccountService{
		log: log.NewHelper(log.With(logger, "module", "service/server-service")),
		uc:  uc,
	}
}
