package consul

import (
	"fmt"
	consulConfig "github.com/go-kratos/kratos/contrib/config/consul/v2"
	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/registry"
	consulAPI "github.com/hashicorp/consul/api"
	"service/pkg/conf"
)

type Consul struct {
	client *consulAPI.Client
	conf   *conf.Consul
}

func New(conf *conf.Consul) *Consul {
	cc := consulAPI.DefaultConfig()

	if conf != nil {
		fmt.Printf("consul address: %v, scheme: %v, datacenter: %v\n",
			conf.Address,
			conf.Scheme,
			conf.Datacenter,
		)

		cc.Address = conf.Address
		cc.Scheme = conf.Scheme
		cc.Datacenter = conf.Datacenter
	}

	client, err := consulAPI.NewClient(cc)
	if err != nil {
		panic(err)
	}

	return &Consul{
		client: client,
		conf:   conf,
	}
}

func (c *Consul) NewDiscovery() registry.Discovery {
	r := consul.New(c.client, consul.WithHealthCheck(false))
	return r
}

func (c *Consul) NewRegistrar() registry.Registrar {
	r := consul.New(c.client, consul.WithHealthCheck(false))
	return r
}

func (c *Consul) NewConfigSource() config.Source {
	cs, err := consulConfig.New(c.client, consulConfig.WithPath(c.conf.Path))
	//consul中需要标注文件后缀，kratos读取配置需要适配文件后缀
	//The file suffix needs to be marked, and kratos needs to adapt the file suffix to read the configuration.
	if err != nil {
		panic(err)
	}

	return cs
}
