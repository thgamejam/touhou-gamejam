//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	"service/app/file_edge/internal/biz"
	"service/app/file_edge/internal/conf"
	"service/app/file_edge/internal/data"
	"service/app/file_edge/internal/server"
	"service/app/file_edge/internal/service"

	pkgConf "service/pkg/conf"
)

// initApp init kratos application.
func initApp(*pkgConf.Server, *pkgConf.Service, *conf.EdgeFile, registry.Registrar, registry.Discovery, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
