package service

import (
	"context"

	pb "service/api/user/v1"
)

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	return &pb.CreateUserReply{}, nil
}
func (s *UserService) GetUserByAccountID(ctx context.Context, req *pb.GetUserByAccountIDRequest) (*pb.GetUserByAccountIDReply, error) {
	return &pb.GetUserByAccountIDReply{}, nil
}
func (s *UserService) GetUsernameByEmail(ctx context.Context, req *pb.GetUserByEmailRequest) (*pb.GetUserReply, error) {
	return &pb.GetUserReply{}, nil
}
