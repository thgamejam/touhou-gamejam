package main

import (
	"flag"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"os"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"service/app/passport/internal/conf"

	pkgConf "service/pkg/conf"
	pkgConsul "service/pkg/consul"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string = "thjam.passport.service"
	// Version is the version of the compiled software.
	Version string
	// flagConfigPath is the config flag.
	flagConfigPath string
	// cloudConfigFile 服务注册，配置中心的配置文件
	cloudConfigFile string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagConfigPath, "conf", "../../configs", "config path, eg: -conf config.yaml")
	flag.StringVar(&cloudConfigFile, "cloud_conf", "../../configs/cloud.yaml", "config path, eg: -cloud_conf cloud.yaml")
}

func newApp(logger log.Logger, rr registry.Registrar, hs *http.Server, gs *grpc.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			hs,
			gs,
		),
		kratos.Registrar(rr),
	)
}

func main() {
	flag.Parse()
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace_id", tracing.TraceID(),
		"span_id", tracing.SpanID(),
	)

	cloudConfig := config.New(
		config.WithSource(
			file.NewSource(cloudConfigFile), // 获取本地的配置文件
		),
	)
	defer cloudConfig.Close()
	// 必须进行一次合并
	if err := cloudConfig.Load(); err != nil {
		panic(err)
	}

	var consulConfig pkgConf.CloudBootstrap
	if err := cloudConfig.Scan(&consulConfig); err != nil {
		panic(err)
	}

	consulUtil := pkgConsul.New(consulConfig.Consul)

	pkgConfig := config.New(
		config.WithSource(
			file.NewSource(flagConfigPath), // 获取本地的配置文件
			consulUtil.NewConfigSource(),   // 获取配置中心的配置文件
		),
	)
	defer pkgConfig.Close()
	if err := pkgConfig.Load(); err != nil {
		panic(err)
	}

	// 读取通用配置到结构体
	var pkgBootstrap pkgConf.PkgBootstrap
	if err := pkgConfig.Scan(&pkgBootstrap); err != nil {
		panic(err)
	}

	// 读取本服务配置到结构体
	var bc conf.Bootstrap
	if err := pkgConfig.Scan(&bc); err != nil {
		panic(err)
	}

	// 服务注册
	rr := consulUtil.NewRegistrar()
	// 服务发现
	rd := consulUtil.NewDiscovery()

	app, cleanup, err := initApp(pkgBootstrap.Server, pkgBootstrap.Service, bc.Passport, rr, rd, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
