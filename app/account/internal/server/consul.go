/*
 * @Autor: Ruic
 * @since: 2021-09-13 15:31:34
 * @LastEditTime: 2021-09-13 15:37:14
 */

package server

import (
	"omega/app/account/internal/conf"

	consul "github.com/go-kratos/consul/registry"
	"github.com/go-kratos/kratos/v2/registry"
	consulAPI "github.com/hashicorp/consul/api"
)

func NewRegistrar(conf *conf.Registry) registry.Registrar {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(true))
	return r
}
