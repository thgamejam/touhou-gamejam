package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	kratosHTTP "github.com/go-kratos/kratos/v2/transport/http"

	//"service/app/file_edge/internal/conf"
	"service/app/file_edge/internal/service"

	pkgConf "service/pkg/conf"
	pkgHTTP "service/pkg/http"
)

const (
	PrefixURL = "/v1/upload/web/"
	URL       = "/v1/upload/web/{token}"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(cSrv *pkgConf.Server, cSce *pkgConf.Service, service *service.FileEdgeService, logger log.Logger) *kratosHTTP.Server {
	var opts = []kratosHTTP.ServerOption{
		kratosHTTP.Middleware(
			recovery.Recovery(),
			validate.Validator(),
		),
	}
	if cSrv.Http.Network != "" {
		opts = append(opts, kratosHTTP.Network(cSrv.Http.Network))
	}
	if cSrv.Http.Addr != "" {
		opts = append(opts, kratosHTTP.Address(cSrv.Http.Addr))
	}
	if cSrv.Http.Timeout != nil {
		opts = append(opts, kratosHTTP.Timeout(cSrv.Http.Timeout.AsDuration()))
	}
	srv := kratosHTTP.NewServer(opts...)
	//v1.RegisterFileEdgeHTTPServer(srv, service)
	RegisterUploadFileHTTPServer(cSrv, cSce, srv, service)
	return srv
}

func RegisterUploadFileHTTPServer(cSrv *pkgConf.Server, cSce *pkgConf.Service, srv *kratosHTTP.Server, service *service.FileEdgeService) {
	route := srv.Route("/")
	if cSrv.Http.RequiredBodySize == 0 {
		panic("server.http.required_body_size cannot be zero")
	}
	if cSrv.Http.MaxBodySize == 0 {
		panic("server.http.max_body_size cannot be zero")
	}
	route.PUT(URL, uploadFile(cSce, service), pkgHTTP.MaxBytesFilter(int64(cSrv.Http.MaxBodySize)+int64(cSrv.Http.RequiredBodySize)))
}
