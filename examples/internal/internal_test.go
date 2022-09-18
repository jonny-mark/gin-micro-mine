/**
 * @author jiangshangfang
 * @date 2022/4/25 4:05 PM
 **/
package internal

import (
	"gin/internal/repository/vehicle"
	"github.com/davecgh/go-spew/spew"
	"github.com/jonny-mark/gin-micro-mine/pkg/config"
	"github.com/jonny-mark/gin-micro-mine/pkg/load/nacos"
	"github.com/jonny-mark/gin-micro-mine/pkg/storage/orm"
	"github.com/spf13/pflag"
	"testing"
)

var (
	cfgDir = pflag.StringP("config dir", "c", "config", "config directory.")
	env    = pflag.StringP("env name", "e", "", "env var name.")
)

func TestVehicle(t *testing.T) {
	//初始化数据库
	pflag.Parse()
	config.New(*cfgDir, config.WithEnv(*env))

	//初始化nacos配置
	nacos.Init()

	orm.Init()
	//usersCard, err := vehicle.FindValidOneByUidAndCardId(uint(1991826963), uint(24))
	//if err != nil {
	//	t.Logf("err:%+v", err)
	//}
	//fmt.Print(usersCard)

	vehicleCards, err := vehicle.FindVehicleObusByValidPlate(uint(1991826963), "苏ZDEAAB", uint(0))
	if err != nil {
		t.Logf("err:%+v", err)
	}
	spew.Dump(vehicleCards)
}
