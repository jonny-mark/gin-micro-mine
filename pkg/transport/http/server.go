package http

import (
	"context"
	"errors"
	"github.com/jonny-mark/gin-micro-mine/pkg/log"
	"github.com/jonny-mark/gin-micro-mine/pkg/utils"
	"net"
	"net/http"
	"net/url"
	"time"
)

type Server struct {
	http.Server
	lis          net.Listener
	network      string
	address      string
	readTimeout  time.Duration
	writeTimeout time.Duration
	log          log.Logger
	endpoint     *url.URL
}

func NewServer(opts ...ServerOption) *Server {
	srv := &Server{
		network:      "tcp",
		address:      ":8080",
		readTimeout:  time.Second,
		writeTimeout: time.Second,
		log:          log.GetLogger(),
	}

	// apply options
	for _, o := range opts {
		o(srv)
	}

	srv.Server = http.Server{
		ReadTimeout:  srv.readTimeout,
		WriteTimeout: srv.writeTimeout,
		Handler:      srv,
	}
	return srv
}

// ServeHTTP should write reply headers and data to the ResponseWriter and then return.
func (s *Server) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	s.ServeHTTP(resp, req)
}

// Start start a server
func (s *Server) Start(ctx context.Context) error {
	lis, err := net.Listen(s.network, s.address)
	if err != nil {
		return err
	}
	s.lis = lis
	s.log.Infof("[HTTP] server is listening on: %s", lis.Addr().String())
	if err := s.Serve(lis); !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

// Stop stop server
func (s *Server) Stop(ctx context.Context) error {
	s.log.Info("[HTTP] server is stopping")
	return s.Shutdown(ctx)
}

func (s *Server) Endpoint() (*url.URL, error) {
	addr, err := utils.Extract(s.address, s.lis)
	if err != nil {
		return nil, err
	}
	s.endpoint = &url.URL{Scheme: "http", Host: addr}
	return s.endpoint, nil
}
