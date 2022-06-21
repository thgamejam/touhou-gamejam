//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	"service/app/user/internal/biz"
	"service/app/user/internal/conf"
	"service/app/user/internal/data"
	"service/app/user/internal/server"
	"service/app/user/internal/service"

	pkgConf "service/pkg/conf"
)

// initApp init kratos application.
func initApp(*pkgConf.Server, *pkgConf.Service, *conf.User, registry.Registrar, registry.Discovery, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
