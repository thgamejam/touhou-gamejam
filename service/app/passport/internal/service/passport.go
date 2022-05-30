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

// VerifyEmail 验证邮箱并返回登录token
func (s *PassportService) VerifyEmail(ctx context.Context, req *pb.VerifyEmailRequest) (*pb.VerifyEmailReply, error) {
	token, err := s.uc.CreatAccount(ctx, req.Body.Sid, req.Body.Key)
	if err != nil {
		return nil, err
	}
	return &pb.VerifyEmailReply{
		Token: token,
	}, nil
}

// GetPublicKey 获取公钥和哈希值
func (s *PassportService) GetPublicKey(ctx context.Context, req *pb.GetPublicKeyRequest) (*pb.GetPublicKeyReply, error) {
	k, h, err := s.uc.GetKey(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.GetPublicKeyReply{
		Key:  k,
		Hash: h,
	}, nil
}

// Login 登录
func (s *PassportService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	ok, token, err := s.uc.Login(ctx, req.Body.Email, req.Body.Password, req.Body.Hash)
	if err != nil {
		return nil, err
	}

	return &pb.LoginReply{
		Ok:    ok,
		Token: token,
	}, nil
}

func (s *PassportService) ChangePassword(ctx context.Context, req *pb.ChangePasswordRequest) (*pb.ChangePasswordReply, error) {
	return &pb.ChangePasswordReply{}, nil
}
