/**
 * @author jiangshangfang
 * @date 2022/1/11 8:33 PM
 **/
package consul

type Option = func(*Registry)

// WithHealthCheck with registry health check option.
func WithHealthCheck(enable bool) Option {
	return func(o *Registry) {
		o.enableHealthCheck = enable
	}
}
