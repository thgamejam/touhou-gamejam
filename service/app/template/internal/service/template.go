package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "service/api/template/v1"
	"service/app/template/internal/biz"
)

type TemplateService struct {
	pb.UnimplementedTemplateServer

	uc  *biz.TemplateUseCase
	log *log.Helper
}

func NewTemplateService(
	uc *biz.TemplateUseCase,
	logger log.Logger,
) *TemplateService {
	return &TemplateService{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

func (s *TemplateService) CreateTemplate(ctx context.Context, req *pb.CreateTemplateRequest) (*pb.CreateTemplateReply, error) {
	return &pb.CreateTemplateReply{}, nil
}
func (s *TemplateService) DeleteTemplate(ctx context.Context, req *pb.DeleteTemplateRequest) (*pb.DeleteTemplateReply, error) {
	return &pb.DeleteTemplateReply{}, nil
}
func (s *TemplateService) GetTemplate(ctx context.Context, req *pb.GetTemplateRequest) (*pb.GetTemplateReply, error) {
	return &pb.GetTemplateReply{}, nil
}
func (s *TemplateService) UpdateTemplate(ctx context.Context, req *pb.UpdateTemplateRequest) (*pb.UpdateTemplateReply, error) {
	return &pb.UpdateTemplateReply{}, nil
}
func (s *TemplateService) ListTemplate(ctx context.Context, req *pb.ListTemplateRequest) (*pb.ListTemplateReply, error) {
	return &pb.ListTemplateReply{}, nil
}
