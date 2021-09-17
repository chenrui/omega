/*
 * @Autor: Ruic
 * @since: 2021-09-13 14:41:23
 * @LastEditTime: 2021-09-17 15:10:49
 */
package data

import (
	"omega/app/account/internal/conf"
	"omega/app/account/internal/data/model"
	"omega/pkg/database/dbengine"
	"omega/pkg/permission"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

type Data struct {
	log  *log.Helper
	db   *dbengine.DBEngine
	perm *permission.Permission
}

func NewDBOptions(cfg *conf.Data) *dbengine.DBOptions {
	return &dbengine.DBOptions{
		Driver:          cfg.Database.GetDriver(),
		Source:          cfg.Database.GetSource(),
		MaxIdleConns:    int(cfg.Database.GetMaxIdleConns()),
		MaxOpenConns:    int(cfg.Database.GetMaxOpenConns()),
		ConnMaxLifetime: cfg.Database.GetMaxLifetime().AsDuration(),
	}
}

func NewPermOptions(cfg *conf.Data) *permission.PermOptions {
	return &permission.PermOptions{
		Model: cfg.GetPermissionModel(),
	}
}

func initData(db *dbengine.DBEngine) {
	// resource := &model.Resource{Appname: "account"}
	// db.Create(resource)
	// roles := []model.Role{{Name: "root"}, {Name: "admin"}, {Name: "employee"}}
	// db.Create(&roles)
	// policies := []model.Policy{{RoleID: 2, Scope: account.SCOPE_ORG}}
	// db.Create(&policies)
}

func NewData(db *dbengine.DBEngine, perm *permission.Permission, logger log.Logger) (*Data, error) {
	db.AutoMigrate(&model.User{}, &model.Role{}, &model.Policy{}, &model.Resource{})
	initData(db)
	return &Data{
		log:  log.NewHelper(log.With(logger, "module", "data")),
		db:   db,
		perm: perm,
	}, nil
}

var ProviderSet = wire.NewSet(NewDBOptions, NewPermOptions, NewData, NewUserRepo)
