package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	pb "service/api/passport/v1"
	"service/app/passport/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewPassportService)

type PassportService struct {
	pb.UnimplementedPassportServer

	uc  *biz.PassportUseCase
	log *log.Helper
}

func NewPassportService(
	uc *biz.PassportUseCase,
	logger log.Logger,
) *PassportService {
	return &PassportService{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}
