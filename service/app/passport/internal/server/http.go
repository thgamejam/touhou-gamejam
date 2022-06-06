package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	nghttp "net/http"
	v1 "service/api/passport/v1"
	"service/app/passport/internal/conf"
	"service/app/passport/internal/service"
	pkgConf "service/pkg/conf"
	"service/pkg/jwt"
	"strings"
)

var LoginAuthRouters = []string{
	"/passport.v1.Passport/ChangePassword",
	"/passport.v1.Passport/Logout",
}

func LocalHttpRequestFilter() http.FilterFunc {
	return func(next nghttp.Handler) nghttp.Handler {
		return nghttp.HandlerFunc(func(w nghttp.ResponseWriter, req *nghttp.Request) {
			req.Header.Add("X-RemoteAddr", strings.Split(req.RemoteAddr, ":")[0])

			next.ServeHTTP(w, req)
		})
	}
}

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *pkgConf.Server, t *conf.Passport, service *service.PassportService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			validate.Validator(),

			selector.Server(jwt.LoginAuthMiddleware([]byte(t.VerifyEmailKey))).Match(jwt.ValidateLoginListMatcher(LoginAuthRouters)).Build(),
		),
		http.Filter(LocalHttpRequestFilter()),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterPassportHTTPServer(srv, service)
	return srv
}
