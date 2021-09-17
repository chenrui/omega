/*
 * @Autor: Ruic
 * @since: 2021-09-14 14:03:19
 * @LastEditTime: 2021-09-15 10:13:50
 */
package dbengine

import (
	"time"

	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBEngine struct {
	*gorm.DB
}

type DBOptions struct {
	Driver          string
	Source          string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

// Open 构造引擎，打开数据库连接池
func NewDBEngine(optins *DBOptions) (*DBEngine, error) {
	db, err := gorm.Open(mysql.Open(optins.Source), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	sqldb, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqldb.SetConnMaxIdleTime(optins.ConnMaxLifetime)
	sqldb.SetMaxIdleConns(optins.MaxIdleConns)
	sqldb.SetMaxOpenConns(optins.MaxOpenConns)
	return &DBEngine{db}, nil
}

var ProviderSet = wire.NewSet(NewDBEngine)
