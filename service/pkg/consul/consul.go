package consul

import (
	"fmt"
	"github.com/go-kratos/kratos/contrib/config/consul/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/registry"
	"service/pkg/conf"

	consulRegistry "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	consulAPI "github.com/hashicorp/consul/api"
)

type Consul struct {
	client *consulAPI.Client
	conf   *conf.Consul
}

func New(conf *conf.Consul) *Consul {
	consulConfig := consulAPI.DefaultConfig()

	if conf != nil {
		fmt.Printf("consul address: %v, scheme: %v, datacenter: %v\n",
			conf.Address,
			conf.Scheme,
			conf.Datacenter,
		)

		consulConfig.Address = conf.Address
		consulConfig.Scheme = conf.Scheme
		consulConfig.Datacenter = conf.Datacenter
	}

	client, err := consulAPI.NewClient(consulConfig)
	if err != nil {
		panic(err)
	}

	return &Consul{
		client: client,
		conf:   conf,
	}
}

func (c *Consul) Registrar() registry.Registrar {
	r := consulRegistry.New(c.client, consulRegistry.WithHealthCheck(false))
	return r
}

func (c *Consul) NewConfigSource() config.Source {
	cs, err := consul.New(c.client, consul.WithPath(c.conf.Path))
	//consul中需要标注文件后缀，kratos读取配置需要适配文件后缀
	//The file suffix needs to be marked, and kratos needs to adapt the file suffix to read the configuration.
	if err != nil {
		panic(err)
	}

	return cs
}
