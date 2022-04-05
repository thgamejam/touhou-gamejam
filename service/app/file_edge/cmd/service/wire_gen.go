// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"service/app/file_edge/internal/biz"
	conf2 "service/app/file_edge/internal/conf"
	"service/app/file_edge/internal/data"
	"service/app/file_edge/internal/server"
	"service/app/file_edge/internal/service"
	"service/pkg/cache"
	"service/pkg/conf"
	"service/pkg/object_storage"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confService *conf.Service, edgeFile *conf2.EdgeFile, registrar registry.Registrar, discovery registry.Discovery, logger log.Logger) (*kratos.App, func(), error) {
	cacheCache, err := cache.NewCache(confService)
	if err != nil {
		return nil, nil, err
	}
	objectStorage, err := object_storage.NewObjectStorage(confService)
	if err != nil {
		return nil, nil, err
	}
	dataData, cleanup, err := data.NewData(cacheCache, objectStorage, logger)
	if err != nil {
		return nil, nil, err
	}
	fileEdgeRepo := data.NewFileEdgeRepo(dataData, logger)
	fileEdgeUseCase := biz.NewFileEdgeUseCase(fileEdgeRepo, logger)
	fileEdgeService := service.NewFileEdgeService(fileEdgeUseCase, logger)
	httpServer := server.NewHTTPServer(confServer, confService, fileEdgeService, logger)
	grpcServer := server.NewGRPCServer(confServer, fileEdgeService, logger)
	app := newApp(logger, registrar, httpServer, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
