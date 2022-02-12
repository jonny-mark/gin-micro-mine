/**
 * @author jiangshangfang
 * @date 2022/2/9 5:39 PM
 **/
package utils

import (
	"testing"
	"fmt"
	"reflect"
)

func TestCompare(t *testing.T) {
	s1 := []string{"h", "e", "l", "l", "o"}
	s2 := []string{"h", "e", "l", "l", "o"}
	if Strings(s1, s2) {
		fmt.Println("equal..")
	} else {
		fmt.Println("not equal!")
	}
}

var (
	slice1 = []string{"foo", "bar", "h", "e", "l", "l", "o"}
	slice2 = []string{"foo", "bar", "h", "e", "l", "l", "oooo"}
)
//1961168	       597.1 ns/op
func BenchmarkStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Strings(slice1, slice2)
	}

}

//1000000	      1030 ns/op
func BenchmarkRelfectDeepEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reflect.DeepEqual(slice1, slice2)
	}
}

//62924170	        19.57 ns/op
//最佳选择
func BenchmarkStringSliceEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringSliceEqual(slice1, slice2)
	}
}