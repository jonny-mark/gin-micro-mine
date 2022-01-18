/**
 * @author jiangshangfang
 * @date 2021/10/24 5:41 PM
 **/
package http

import "time"

// ServerOption is HTTP server option
type ServerOption func(*Server)

// WithAddress with server address.
func WithAddress(addr string) ServerOption {
	return func(s *Server) {
		s.address = addr
	}
}

// WithReadTimeout with read timeout.
func WithReadTimeout(timeout time.Duration) ServerOption {
	return func(s *Server) {
		s.readTimeout = timeout
	}
}

// WithWriteTimeout with write timeout.
func WithWriteTimeout(timeout time.Duration) ServerOption {
	return func(s *Server) {
		s.writeTimeout = timeout
	}
}
