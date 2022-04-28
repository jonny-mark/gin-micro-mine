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
	Uid                 uint       `gorm:"index:uid;column:uid;type:int unsigned;not null;default:0" json:"uid"`                             // 用户id
	VehicleId           int        `gorm:"index:vehicle_id;column:vehicle_id;type:int;not null;default:0" json:"vehicleId"`                  // etc_users_cards表的主键id
	CardNo              string     `gorm:"index:card_no;column:card_no;type:varchar(20);not null;default:''" json:"cardNo"`                  // 卡号
	CardStatus          int8       `gorm:"column:card_status;type:tinyint;not null;default:0" json:"cardStatus"`                             // 卡状态[0-待发行 10-发行中 11-已发行 20-注销中 21-已注销 22-换卡注销]
	CardBlacklistStatus int8       `gorm:"column:card_blacklist_status;type:tinyint;not null;default:0" json:"cardBlacklistStatus"`          // 卡片黑名单状态[0-未拉黑 1-黑名单]
	BusinessSource      int8       `gorm:"column:business_source;type:tinyint;not null;default:0" json:"businessSource"`                     // 业务来源[0-未知 1-系统 2-申办 3-补办业务 4-维修换货 5-退货退款 6-注销 7-二次激活]
	BusinessType        int8       `gorm:"column:business_type;type:tinyint;not null;default:0" json:"businessType"`                         // 业务类型[0-未知 1-新办 2-更换 3-激活 20-注销 21-挂失]
	CardReleasedAt      *time.Time `gorm:"index:card_released_at;column:card_released_at;type:timestamp;default:null" json:"cardReleasedAt"` // 卡片发行时间
	CardEnableTime      *time.Time `gorm:"column:card_enable_time;type:timestamp;default:null" json:"cardEnableTime"`                        // 卡片启用时间
	CardExpireTime      *time.Time `gorm:"index:card_expire_time;column:card_expire_time;type:timestamp;default:null" json:"cardExpireTime"` // 卡片失效时间
	CardRevokedAt       *time.Time `gorm:"index:card_revoked_at;column:card_revoked_at;type:timestamp;default:null" json:"cardRevokedAt"`    // 卡注销时间
	CreatedAt           time.Time  `gorm:"column:created_at" json:"-"`
	UpdatedAt           time.Time  `gorm:"column:updated_at" json:"-"`
}

type VehiclesCardsVehicleModel struct {
	VehiclesCardsModel
	Vehicle VehiclesModel `gorm:"foreignKey:ID;associationForeignKey:VehicleId" json:"vehicle"`
}
