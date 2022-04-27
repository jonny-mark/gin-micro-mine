/**
 * @author jiangshangfang
 * @date 2022/1/26 5:45 PM
 **/
package directory

import (
	"testing"
)

func TestPathExists(t *testing.T) {
	path := string("../../pkg/directory/")
	file := path + "directory.go"
	println(PathExists(path))
	println(FileExist(file))
}

func TestFileMove(t *testing.T) {
	a := string("../../pkg/directory/README.md")
	b := string("../../pkg/conversion/ccc.md")
	println(FileMove(b, a).Error())
}
