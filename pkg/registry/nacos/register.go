package nacos

import "github.com/jonny-mark/gin-micro-mine/pkg/registry"

var (
	_ registry.Registry  = (*Registry)(nil)
	_ registry.Discovery = (*Registry)(nil)
)
