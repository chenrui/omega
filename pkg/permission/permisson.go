/*
 * @Autor: Ruic
 * @since: 2021-09-15 16:08:03
 * @LastEditTime: 2021-09-17 09:53:16
 */
package permission

import (
	"omega/pkg/database/dbengine"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/google/wire"
)

type PermOptions struct {
	Model string
}

type Permission struct {
	*casbin.Enforcer
}

func NewPermission(db *dbengine.DBEngine, opts *PermOptions) (*Permission, error) {
	adapter, err := gormadapter.NewAdapterByDB(db.DB)
	if err != nil {
		return nil, err
	}
	m, err := model.NewModelFromString(opts.Model)
	if err != nil {
		return nil, err
	}
	enforce, err := casbin.NewEnforcer(m, adapter)
	if err != nil {
		return nil, err
	}
	return &Permission{enforce}, nil
}

var ProviderSet = wire.NewSet(NewPermission)
