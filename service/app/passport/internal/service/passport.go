package service

import (
	"context"
	"service/app/passport/internal/biz"

	pb "service/api/passport/v1"
)

// CreateAccount 预创建账户
func (s *PassportService) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*pb.CreateAccountReply, error) {
	err := s.uc.PrepareCreateAccount(ctx, biz.Account{
		Email:    req.Body.Email,
		Password: req.Body.Password,
		Hash:     req.Body.Hash,
	}, req.Body.Token)
	if err != nil {
		return nil, err
	}
	return &pb.CreateAccountReply{
		Ok: true,
	}, nil
}
func (s *PassportService) VerifyEmail(ctx context.Context, req *pb.VerifyEmailRequest) (*pb.VerifyEmailReply, error) {
	return &pb.VerifyEmailReply{}, nil
}
func (s *PassportService) GetPublicKey(ctx context.Context, req *pb.GetPublicKeyRequest) (*pb.GetPublicKeyReply, error) {
	return &pb.GetPublicKeyReply{}, nil
}
func (s *PassportService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	return &pb.LoginReply{}, nil
}
func (s *PassportService) ChangePassword(ctx context.Context, req *pb.ChangePasswordRequest) (*pb.ChangePasswordReply, error) {
	return &pb.ChangePasswordReply{}, nil
}
