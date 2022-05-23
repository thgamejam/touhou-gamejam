package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "service/api/passport/v1"
	"service/app/passport/internal/biz"
)

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

func (s *PassportService) CreatePassport(ctx context.Context, req *pb.CreatePassportRequest) (*pb.CreatePassportReply, error) {
	return &pb.CreatePassportReply{}, nil
}
func (s *PassportService) DeletePassport(ctx context.Context, req *pb.DeletePassportRequest) (*pb.DeletePassportReply, error) {
	return &pb.DeletePassportReply{}, nil
}
func (s *PassportService) GetPassport(ctx context.Context, req *pb.GetPassportRequest) (*pb.GetPassportReply, error) {
	return &pb.GetPassportReply{}, nil
}
func (s *PassportService) UpdatePassport(ctx context.Context, req *pb.UpdatePassportRequest) (*pb.UpdatePassportReply, error) {
	return &pb.UpdatePassportReply{}, nil
}
func (s *PassportService) ListPassport(ctx context.Context, req *pb.ListPassportRequest) (*pb.ListPassportReply, error) {
	return &pb.ListPassportReply{}, nil
}
