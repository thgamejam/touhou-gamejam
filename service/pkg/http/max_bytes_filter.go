package http

import "net/http"

type MaxBytesHandler struct {
	http.Handler
	maxBytesSize int64 // http Body 最大大小 单位Byte
}

func (h MaxBytesHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	req.Body = http.MaxBytesReader(w, req.Body, h.maxBytesSize)
	h.Handler.ServeHTTP(w, req)
}

// MaxBytesFilter HTTP最大字节过滤器
// 返回一个限制 http Body 最大大小的 kratos.http.FilterFunc
// maxBytesSize 单位kb
func MaxBytesFilter(maxBytesSize int64) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return &MaxBytesHandler{
			Handler:      h,
			maxBytesSize: maxBytesSize << 10,
		}
	}
}
