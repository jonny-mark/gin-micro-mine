/**
 * @author jiangshangfang
 * @date 2022/4/27 4:40 PM
 **/
package constant

import "time"

const (
	StatusNormal = 0

	//cache前缀, 规则：业务+模块+{ID}
	PrefixUserBaseCacheKey =  "gin:user:base:%d"

	//默认user的cache过期时间
	DefaultExpireTime = time.Hour * 24
)