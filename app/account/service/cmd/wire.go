// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"omega/app/account/internal/biz"
	"omega/app/account/internal/conf"
	"omega/app/account/internal/data"
	"omega/app/account/internal/server"
	"omega/app/account/internal/service"
	"omega/pkg/database/dbengine"
	"omega/pkg/permission"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Registry, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(dbengine.ProviderSet, permission.ProviderSet, server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
