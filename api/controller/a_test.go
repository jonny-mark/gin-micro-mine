/**
 * @author jiangshangfang
 * @date 2021/8/8 5:28 PM
 **/
package controller

import (
	"gin/common/response"
	"testing"
	"time"
)

func TestAa(t *testing.T) {
	c := time.Tick(3)
	d := time.Tick(3)
	now := time.Now()
	t.Log(c, d,now)
	response.OkWithData("创建成功")
	t.Log("222")
}
