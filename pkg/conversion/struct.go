/**
 * @author jiangshangfang
 * @date 2022/1/26 11:38 AM
 **/
package conversion

import (
	"github.com/fatih/structs"
)

// 结构体转化为map
func StructToMap(obj interface{}) map[string]interface{} {
	return structs.Map(obj)
}
