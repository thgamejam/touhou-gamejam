package biz

import (
	"context"
	"os"

	"github.com/go-kratos/kratos/v2/log"
)

type FileInfo struct {
	ID        uint64   // 文件ID
	Name      string   // 文件名
	ExpiresAt int64    // 到期时间戳
	File      *os.File // 文件File
}

type FileEdgeRepo interface {
	PutFile(context.Context, *FileInfo) error
}

type FileEdgeUseCase struct {
	repo FileEdgeRepo
	log  *log.Helper
}

func NewFileEdgeUseCase(repo FileEdgeRepo, logger log.Logger) *FileEdgeUseCase {
	return &FileEdgeUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *FileEdgeUseCase) PutFile(ctx context.Context, f *FileInfo) error {
	return uc.repo.PutFile(ctx, f)
}
