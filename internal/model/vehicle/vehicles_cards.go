/**
 * @author jiangshangfang
 * @date 2022/4/3 6:31 PM
 **/
package vehicle

import (
	"time"
)

// TableName 表名
func (u *VehiclesCardsModel) TableName() string {
	return "etc_vehicles_cards"
}

type VehiclesCardsModel struct {
	ID                  uint64     `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	Uid                 int        `gorm:"column:uid" json:"uid"`                                     // 用户id
	VehicleId           uint       `gorm:"column:vehicle_id" json:"vehicle_id"`                       // etc_users_cards表的主键id
	CardNo              string     `gorm:"column:card_no" json:"card_no"`                             // 卡号
	CardStatus          int        `gorm:"column:card_status" json:"card_status"`                     // 卡状态[0-待发行 10-发行中 11-已发行 20-注销中 21-已注销 22-换卡注销]
	CardBlacklistStatus int        `gorm:"column:card_blacklist_status" json:"card_blacklist_status"` // 卡片黑名单状态[0-未拉黑 1-黑名单]
	BusinessSource      int        `gorm:"column:business_source" json:"business_source"`             // 业务来源[0-未知 1-系统 2-申办 3-补办业务 4-维修换货 5-退货退款 6-注销 7-二次激活]
	BusinessType        int        `gorm:"column:business_type" json:"business_type"`                 // 业务类型[0-未知 1-新办 2-更换 3-激活 20-注销 21-挂失]
	CardReleasedAt      *time.Time `gorm:"column:card_released_at" json:"card_released_at"`           // 卡片发行时间
	CardEnableTime      *time.Time `gorm:"column:card_enable_time" json:"card_enable_time"`           // 卡片启用时间
	CardExpireTime      *time.Time `gorm:"column:card_expire_time" json:"card_expire_time"`           // 卡片失效时间
	CardRevokedAt       *time.Time `gorm:"column:card_revoked_at" json:"card_revoked_at"`             // 卡注销时间
	CreatedAt           time.Time  `gorm:"column:created_at" json:"-"`
	UpdatedAt           time.Time  `gorm:"column:updated_at" json:"-"`
}

type VehiclesCardsVehicleModel struct {
	VehiclesCardsModel
	Vehicle VehiclesModel `gorm:"foreignKey:ID;associationForeignKey:VehicleId" json:"vehicle"`
}
