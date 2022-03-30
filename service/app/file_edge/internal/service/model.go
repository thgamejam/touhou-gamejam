package service

import (
	"os"
)

// UploadFileRequest 上传文件请求
type UploadFileRequest struct {
	// 以下数据为通过 http server 处理后获得
	Name string   `json:"name"` // 文件名
	Path string   `json:"path"` // 文件路径
	File *os.File `json:"-"`    // 文件File

	// 以下数据为 http URL 中 JWT 的数据
	ID        uint64 `json:"id"`                  // 文件ID
	MD5       string `json:"md5"`                 // MD5Hash值
	SHA1      string `json:"sha1"`                // SHA1Hash值
	SliceID   uint32 `json:"slice_id,omitempty"`  // 分片ID
	SliceLen  uint32 `json:"slice_len,omitempty"` // 分片长度
	ExpiresAt int64  `json:"exp"`                 // 到期时间戳
}

// UploadFileReply 上传文件回复
type UploadFileReply struct {
	Ok bool `json:"ok,omitempty"`
}
