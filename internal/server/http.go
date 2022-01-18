/**
 * @author jiangshangfang
 * @date 2021/10/23 7:32 PM
 **/
package server

import (
	"gin/pkg/transport/http"
	"gin/internal/router"
)

func NewHTTPServer(cfg *http.Http) *http.Server {
	r := router.NewRouter()
	srv := http.NewServer(
		http.WithAddress(cfg.Addr),
		http.WithReadTimeout(cfg.ReadTimeout),
		http.WithWriteTimeout(cfg.WriteTimeout),
	)
	srv.Handler = r

	return srv
}
