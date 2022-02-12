/**
 * @author jiangshangfang
 * @date 2022/2/9 5:38 PM
 **/
package utils

//比较的封装
type Interface interface {
	// len方法返回集合中的元素个数
	Len() int
	// item方法根据索引返回集合的值
	Item(idx int) interface{}
}

type IntSlice []int

func (p IntSlice) Len() int                 { return len(p) }
func (p IntSlice) Item(idx int) interface{} { return p[idx] }
func Ints(a, b []int) bool {
	return Compare(IntSlice(a), IntSlice(b))
}

type StringSlice []string

func (p StringSlice) Len() int                 { return len(p) }
func (p StringSlice) Item(idx int) interface{} { return p[idx] }
func Strings(a, b []string) bool {
	return Compare(StringSlice(a), StringSlice(b))
}

// compare 循环遍历比较两个集合
// 先比较两个集合的长度是否相等
// 再循环遍历每一个元素进行比较
func Compare(a, b Interface) bool {
	if a.Len() != b.Len() {
		return false
	}
	if (a == nil) != (b == nil) {
		return false
	}
	for i := 0; i < a.Len(); i++ {
		if a.Item(i) != b.Item(i) {
			return false
		}
	}
	return true
}