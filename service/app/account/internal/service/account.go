package service

import (
    "context"
    "github.com/go-kratos/kratos/v2/log"
    "service/app/account/internal/biz"

    pb "service/api/account/v1"
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

func (s *AccountService) CreateEMailAccount(ctx context.Context, req *pb.CreateEMailAccountReq) (*pb.CreateEMailAccountReply, error) {
    return &pb.CreateEMailAccountReply{}, nil
}
func (s *AccountService) GetAccount(ctx context.Context, req *pb.GetAccountReq) (*pb.GetAccountReply, error) {
    return &pb.GetAccountReply{}, nil
}
func (s *AccountService) VerifyPassword(ctx context.Context, req *pb.VerifyPasswordReq) (*pb.VerifyPasswordReply, error) {
    return &pb.VerifyPasswordReply{}, nil
}
func (s *AccountService) SavePassword(ctx context.Context, req *pb.SavePasswordReq) (*pb.SavePasswordReply, error) {
    return &pb.SavePasswordReply{}, nil
}
func (s *AccountService) GetKey(ctx context.Context, req *pb.GetKeyReq) (*pb.GetKeyReply, error) {
    return &pb.GetKeyReply{}, nil
}
