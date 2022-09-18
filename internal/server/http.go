package server

import (
	"github.com/jonny-mark/gin-micro-mine/internal/routers"
	"github.com/jonny-mark/gin-micro-mine/pkg/app"
	"github.com/jonny-mark/gin-micro-mine/pkg/transport/http"
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
