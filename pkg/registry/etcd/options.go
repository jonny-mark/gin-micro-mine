/**
 * @author jiangshangfang
 * @date 2022/1/11 8:33 PM
 **/
package etcd

import (
	"context"
	"time"
)

// Option is etcd registry option.
type Option func(o *options)

type options struct {
	ctx       context.Context
	namespace string
	ttl       time.Duration
	maxRetry  int
}

// WithContext with registry context.
func WithContext(ctx context.Context) Option {
	return func(o *options) { o.ctx = ctx }
}

// WithNamespace with registry namespace.
func WithNamespace(ns string) Option {
	return func(o *options) { o.namespace = ns }
}

// WithRegisterTTL with register ttl.
func WithRegisterTTL(ttl time.Duration) Option {
	return func(o *options) { o.ttl = ttl }
}

// WithMaxRetry set max retry times.
func WithMaxRetry(num int) Option {
	return func(o *options) { o.maxRetry = num }
}
