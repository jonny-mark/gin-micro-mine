/**
 * @author jiangshangfang
 * @date 2021/12/12 8:38 PM
 **/
package app

import (
	"testing"
	"path/filepath"
	"github.com/davecgh/go-spew/spew"
)

func TestLoad(t *testing.T) {
	c := filepath.Join("sss/bbb","dev")
	println(c)
	spew.Dump(c)
}