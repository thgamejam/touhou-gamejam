package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	accountV1 "service/api/account/v1"
	userV1 "service/api/user/v1"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewPassportRepo,
	NewAccountServiceClient,
	NewUserServiceClient,
	// TODO 数据客户端构建函数
)

// Data .
type Data struct {
	// TODO 封装的数据客户端
	accountClient accountV1.AccountClient
	userClient    userV1.UserClient
}

// NewData .
func NewData(
	// TODO 需要的数据客户端
	accountClient accountV1.AccountClient,
	userClient userV1.UserClient,
	logger log.Logger,
) (*Data, func(), error) {
	data := &Data{
		accountClient: accountClient,
		userClient:    userClient,
	}

	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return data, cleanup, nil
}

func NewUserServiceClient(r registry.Discovery) userV1.UserClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///thjam.user.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	return userV1.NewUserClient(conn)
}

func NewAccountServiceClient(r registry.Discovery) accountV1.AccountClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///thjam.account.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	c := accountV1.NewAccountClient(conn)
	return c
}
