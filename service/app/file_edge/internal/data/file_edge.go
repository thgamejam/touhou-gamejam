package data

import (
	"context"
	"strconv"

	"github.com/go-kratos/kratos/v2/log"
	"service/app/file_edge/internal/biz"
)

type fileEdgeRepo struct {
	data *Data
	log  *log.Helper
}

// NewFileEdgeRepo .
func NewFileEdgeRepo(data *Data, logger log.Logger) biz.FileEdgeRepo {
	return &fileEdgeRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

var fileEdgeCacheKey = func(id uint64) string {
	return "edge_file_" + strconv.FormatUint(id, 10)
}

type FileInfoCache struct {
	ID uint64 // 文件ID
}

func (r *fileEdgeRepo) PutFile(ctx context.Context, info *biz.FileInfo) error {
	// TODO implement me
	panic("implement me")
}
