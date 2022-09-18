/**
 * @author jiangshangfang
 * @date 2021/10/23 7:32 PM
 **/
package server

import (
	"gin/internal/routers"
	"gin-micro-mine/pkg/app"
	"gin-micro-mine/pkg/transport/http"
)

func NewHTTPServer(cfg *app.ServerConfig) *http.Server {
	r := routers.NewRouter()
	srv := http.NewServer(
		http.WithAddress(cfg.Addr),
		http.WithReadTimeout(cfg.ReadTimeout),
		http.WithWriteTimeout(cfg.WriteTimeout),
	)
	srv.Handler = r

	return srv
}
