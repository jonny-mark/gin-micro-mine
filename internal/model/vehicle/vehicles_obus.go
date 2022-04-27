/**
 * @author jiangshangfang
 * @date 2022/4/3 6:31 PM
 **/
package vehicle

import (
	"time"
)

// TableName 表名
func (u *VehiclesObusModel) TableName() string {
	return "etc_vehicles_obus"
}

type VehiclesObusModel struct {
	ID                 uint64     `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	Uid                int        `gorm:"column:uid" json:"uid"`                                   // 用户id
	VehicleId          uint       `gorm:"column:vehicle_id" json:"vehicle_id"`                     // users_cards表的主键id
	ObuDeviceSn        string     `gorm:"column:obu_device_sn" json:"obu_device_sn"`               // OBU设备电子标签号
	ObuStatus          int        `gorm:"column:obu_status" json:"obu_status"`                     // OBU状态[0-待发行 10-发行中 11-已发行 12-已激活 20-注销中 21-已注销 22-换签注销]
	ObuBlacklistStatus int        `gorm:"column:obu_blacklist_status" json:"obu_blacklist_status"` // OBU黑名单状态[0-未拉黑 1-黑名单]
	Manufacturer       int        `gorm:"column:manufacturer" json:"manufacturer"`                 // 设备厂商[0-未知 1-埃特斯 2-金溢 3-聚力 4-万集 5-成谷 6-云星宇]
	DeviceType         int        `gorm:"column:device_type" json:"device_type"`                   // 设备类型[0-普通 1-可充电设备 2-单片式OBU]
	BusinessSource     int        `gorm:"column:business_source" json:"business_source"`           // 业务来源[0-未知 1-系统 2-申办 3-补办业务 4-维修换货 5-退货退款 6-注销 7-二次激活]
	BusinessType       int        `gorm:"column:business_type" json:"business_type"`               // 业务类型[0-未知 1-新办 2-更换 3-激活 20-注销 21-挂失]
	ObuReleasedAt      *time.Time `gorm:"column:obu_released_at" json:"obu_released_at"`           // OBU发行时间
	FirstActivatedAt   *time.Time `gorm:"column:first_activated_at" json:"first_activated_at"`     // OBU首次时间
	ActivatedAt        *time.Time `gorm:"column:activated_at" json:"activated_at"`                 // OBU最近激活时间
	ObuEnableTime      *time.Time `gorm:"column:obu_enable_time" json:"obu_enable_time"`           // OBU启用时间
	ObuExpireTime      *time.Time `gorm:"column:obu_expire_time" json:"obu_expire_time"`           // OBU失效时间
	WarrantyStartAt    *time.Time `gorm:"column:warranty_start_at" json:"warranty_start_at"`       // 质保开始时间
	WarrantyEndAt      *time.Time `gorm:"column:warranty_end_at" json:"warranty_end_at"`           // 质保结束时间
	ObuRevokedAt       *time.Time `gorm:"column:obu_revoked_at" json:"obu_revoked_at"`             // OBU注销时间
	CreatedAt          time.Time  `gorm:"column:created_at" json:"-"`
	UpdatedAt          time.Time  `gorm:"column:updated_at" json:"-"`
}

type VehiclesObusVehicleModel struct {
	VehiclesObusModel
	Vehicle VehiclesModel `gorm:"foreignKey:ID;associationForeignKey:VehicleId" json:"vehicle"`
}
