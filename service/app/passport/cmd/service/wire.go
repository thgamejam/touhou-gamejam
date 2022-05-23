//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	"service/app/passport/internal/biz"
	"service/app/passport/internal/data"
	"service/app/passport/internal/server"
	"service/app/passport/internal/service"

	pkgConf "service/pkg/conf"
)

// initApp init kratos application.
func initApp(*pkgConf.Server, *pkgConf.Service, registry.Registrar, registry.Discovery, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
