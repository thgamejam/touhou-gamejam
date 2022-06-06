package service

import (
	"context"
	"errors"
	pb "service/api/passport/v1"
	"service/app/passport/internal/biz"
	"service/pkg/jwt"
)

//Logout 登出请求
func (s *PassportService) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutReply, error) {
	loginToken, ok := jwt.FromLoginTokenContext(ctx)
	if !ok {
		return nil, errors.New("TokenNotFound")
	}
	err := s.uc.Logout(ctx, loginToken.UserID, loginToken.UUID)
	if err != nil {
		return nil, err
	}
	return &pb.LogoutReply{}, nil
}

// RenewalToken 续签Token
func (s *PassportService) RenewalToken(ctx context.Context, req *pb.RenewalTokenRequest) (*pb.RenewalTokenReply, error) {
	token, err := s.uc.RenewalToken(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.RenewalTokenReply{
		Token: token,
	}, nil
}

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

// ChangePassword 修改密码
func (s *PassportService) ChangePassword(ctx context.Context, req *pb.ChangePasswordRequest) (*pb.ChangePasswordReply, error) {
	loginToken, ok := jwt.FromLoginTokenContext(ctx)
	if !ok {
		return nil, errors.New("TokenNotFound")
	}
	newToken, err := s.uc.ChangePassword(ctx, loginToken.UserID, req.Body.NewPassword, req.Body.Hash)
	if err != nil {
		return nil, err
	}
	return &pb.ChangePasswordReply{
		Ok:    true,
		Token: newToken,
	}, nil
}
