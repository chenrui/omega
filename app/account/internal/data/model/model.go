/*
 * @Autor: Ruic
 * @since: 2021-09-15 09:37:50
 * @LastEditTime: 2021-09-17 15:14:31
 */
package model

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primarykey"`
	Name      string `gorm:"size:64"`
	Password  string `gorm:"size:64"`
	Phone     string `gorm:"size:20"`
	RoleID    uint
	Role      Role
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Role struct {
	ID          uint   `gorm:"primarykey"`
	Name        string `gorm:"size:32"`
	Description string `gorm:"size:64"`
	Policies    []Policy
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Policy struct {
	ID        uint `gorm:"primarykey"`
	RoleID    uint
	Scope     string     `gorm:"size:64"` //作用域 sys/org/area/dept
	Resources []Resource `gorm:"many2many:policy_resource_relationship;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Resource struct {
	ID        uint   `gorm:"primarykey"`
	Appname   string `gorm:"size:64"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
