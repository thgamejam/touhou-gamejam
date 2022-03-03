package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "service/api/account/v1"
	"service/app/account/internal/biz"
)

type AccountService struct {
	pb.UnimplementedAccountServer

	uc  *biz.AccountUseCase
	log *log.Helper
}

func NewAccountService(
	uc *biz.AccountUseCase,
	logger log.Logger,
) *AccountService {
	return &AccountService{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

func (s *AccountService) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*pb.CreateAccountReply, error) {
	return &pb.CreateAccountReply{}, nil
}
func (s *AccountService) DeleteAccount(ctx context.Context, req *pb.DeleteAccountRequest) (*pb.DeleteAccountReply, error) {
	return &pb.DeleteAccountReply{}, nil
}
func (s *AccountService) GetAccount(ctx context.Context, req *pb.GetAccountRequest) (*pb.GetAccountReply, error) {
	return &pb.GetAccountReply{}, nil
}
func (s *AccountService) UpdateAccount(ctx context.Context, req *pb.UpdateAccountRequest) (*pb.UpdateAccountReply, error) {
	return &pb.UpdateAccountReply{}, nil
}
func (s *AccountService) ListAccount(ctx context.Context, req *pb.ListAccountRequest) (*pb.ListAccountReply, error) {
	return &pb.ListAccountReply{}, nil
}
