package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"service/app/user/internal/biz"

	pb "service/api/user/v1"
)

type UserService struct {
	pb.UnimplementedUserServer

	uc  *biz.UserUseCase
	log *log.Helper
}

func NewUserService(
	uc *biz.UserUseCase,
	logger log.Logger,
) *UserService {
	return &UserService{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	return &pb.CreateUserReply{}, nil
}
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	return &pb.DeleteUserReply{}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	return &pb.GetUserReply{}, nil
}
func (s *UserService) GetUserByUsername(ctx context.Context, req *pb.GetUserByUsernameRequest) (*pb.GetUserReply, error) {
	return &pb.GetUserReply{}, nil
}
func (s *UserService) GetUsernameByEmail(ctx context.Context, req *pb.GetUsernameByEmailRequest) (*pb.GetUsernameReply, error) {
	return &pb.GetUsernameReply{}, nil
}
func (s *UserService) GetUsernameByPhone(ctx context.Context, req *pb.GetUsernameByPhoneRequest) (*pb.GetUsernameReply, error) {
	return &pb.GetUsernameReply{}, nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	return &pb.UpdateUserReply{}, nil
}
func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	return &pb.ListUserReply{}, nil
}
func (s *UserService) VerifyPassword(ctx context.Context, req *pb.VerifyPasswordRequest) (*pb.VerifyPasswordReply, error) {
	return &pb.VerifyPasswordReply{}, nil
}
