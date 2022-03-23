// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"service/app/user/internal/biz"
	"service/app/user/internal/data"
	"service/app/user/internal/server"
	"service/app/user/internal/service"
	"service/pkg/cache"
	"service/pkg/conf"
	"service/pkg/database"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confService *conf.Service, registrar registry.Registrar, discovery registry.Discovery, logger log.Logger) (*kratos.App, func(), error) {
	db, err := database.NewDataBase(confService)
	if err != nil {
		return nil, nil, err
	}
	cacheCache, err := cache.NewCache(confService)
	if err != nil {
		return nil, nil, err
	}
	dataData, cleanup, err := data.NewData(db, cacheCache, logger)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	userUseCase := biz.NewUserUseCase(userRepo, logger)
	userService := service.NewUserService(userUseCase, logger)
	httpServer := server.NewHTTPServer(confServer, userService, logger)
	grpcServer := server.NewGRPCServer(confServer, userService, logger)
	app := newApp(logger, registrar, httpServer, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
