package conf

import (
    "github.com/go-kratos/kratos/v2/config"
    "github.com/go-kratos/kratos/v2/config/file"
)

func NewPkgTestConfService() *PkgBootstrap {
    flagConfigPath := "../../app/template/configs"

    pkgConfig := config.New(
        config.WithSource(
            file.NewSource(flagConfigPath), // 获取本地的配置文件
        ),
    )
    defer pkgConfig.Close()
    if err := pkgConfig.Load(); err != nil {
        panic(err)
    }

    // 读取通用配置到结构体
    var pkgBootstrap PkgBootstrap
    if err := pkgConfig.Scan(&pkgBootstrap); err != nil {
        panic(err)
    }

    return &pkgBootstrap
}
