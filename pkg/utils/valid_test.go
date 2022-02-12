package utils

import (
	"testing"
)

// IsZero 检查是否是零值
func TestIsZero(t *testing.T) {
	a := 0
	b := []int{1, 2, 0}
	c := struct{}{}
	var d = []int{1, 2, 3}
	println(IsZero(a))
	println(IsZero(b))
	println(IsZero(c))
	println(IsZero(d))
}
