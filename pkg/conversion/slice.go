/**
 * @author jiangshangfang
 * @date 2022/1/26 5:16 PM
 **/
package conversion

import (
	"strings"
	"fmt"
)

//数组||切片格式化为字符串
func SliceToString(array []interface{}) string {
	return strings.Replace(strings.Trim(fmt.Sprint(array), "[]"), " ", ",", -1)
}