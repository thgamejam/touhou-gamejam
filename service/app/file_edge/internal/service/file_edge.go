package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	pb "service/api/file_edge/v1"
	"service/app/file_edge/internal/biz"
)

type FileEdgeService struct {
	pb.UnimplementedFileEdgeServer

	uc  *biz.FileEdgeUseCase
	log *log.Helper
}

func NewFileEdgeService(
	uc *biz.FileEdgeUseCase,
	logger log.Logger,
) *FileEdgeService {
	return &FileEdgeService{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

func (s *FileEdgeService) UploadFile(ctx context.Context, req *UploadFileRequest) (*UploadFileReply, error) {
	return &UploadFileReply{}, nil
}
