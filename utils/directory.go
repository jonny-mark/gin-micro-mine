/**
 * @author jiangshangfang
 * @date 2021/8/8 3:41 PM
 **/
package utils

import "os"

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}