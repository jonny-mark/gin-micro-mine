/**
 * @author jiangshangfang
 * @date 2021/10/23 10:17 PM
 **/
package transport

import "context"

type Server interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}
