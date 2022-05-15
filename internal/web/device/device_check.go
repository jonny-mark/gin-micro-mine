/**
 * @author jiangshangfang
 * @date 2022/4/16 6:13 PM
 **/
package device

import (
	"github.com/gin-gonic/gin"
	"time"
)

type CheckCreate struct {
	PlateNo            string        `json:"plate_no" form:"plate_no" binding:"required"`
	PlateColor         *uint         `json:"plate_color" form:"plate_color" binding:"required"`
	BusinessType       int           `json:"business_type" form:"business_type" binding:"required"`
	BusinessSn         string        `json:"business_sn" form:"business_sn"`
	CardNo             string        `json:"card_no" form:"card_no"`
	CardType           int           `json:"card_type" form:"card_type"`
	ObuDeviceSn        string        `json:"obu_device_sn" form:"obu_device_sn"`
	DeviceHealthStatus int           `json:"device_health_status" form:"device_health_status"`
	Manufacturer       int           `json:"manufacturer" form:"manufacturer"`
	DisassemblyStatus  string        `json:"disassembly_status" form:"disassembly_status"`
	CardEnableTime     time.Duration `json:"card_enable_time" form:"card_enable_time"`
	CardExpireTime     time.Duration `json:"card_expire_time" form:"card_expire_time"`
	UserPlateNo        string        `json:"user_plate_no" form:"user_plate_no"`
	UserPlateColor     *uint         `json:"user_plate_color" form:"user_plate_color"`
	Electricity        int           `json:"electricity" form:"electricity"`
}

// 创建一条设备检测
func Create(c *gin.Context) {
	//var d CheckCreate
	//err := c.Bind(d)
}
