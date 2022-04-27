/**
 * @author jiangshangfang
 * @date 2021/10/23 10:17 PM
 **/
package transport

import (
	"context"
	"net/url"
)

type Server interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}

type Endpoint interface {
	Endpoint() (*url.URL, error)
}
