//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	"service/app/account/internal/biz"
	"service/app/account/internal/data"
	"service/app/account/internal/server"
	"service/app/account/internal/service"

	"service/app/account/internal/conf"
	pkgConf "service/pkg/conf"
)

// initApp init kratos application.
func initApp(*pkgConf.Server, *pkgConf.Service, *conf.Account, registry.Registrar, registry.Discovery, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
