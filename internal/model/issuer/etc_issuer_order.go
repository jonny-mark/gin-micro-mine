/**
 * @author jiangshangfang
 * @date 2022/1/17 3:07 PM
 **/
package issuer

import (
	"github.com/go-playground/validator/v10"
	"time"
)

// TableName 表名
func (u *IssuerOrderModel) TableName() string {
	return "etc_issuer_order"
}

type IssuerOrderModel struct {
	Id               int    `gorm:"id" json:"id"`                                 // id
	OrderSn          string `gorm:"order_sn" json:"order_sn"`                     // 业务订单号
	PushOrderSn      string `gorm:"push_order_sn" json:"push_order_sn"`           // 推送订单号
	PlateNo          string `gorm:"plate_no" json:"plate_no"`                     // 车牌
	PlateColor       int    `gorm:"plate_color" json:"plate_color"`               // 车牌颜色
	Issuer           string `gorm:"issuer" json:"issuer"`                         // 发卡方定义（不作为选择Service的前提，只允许做更新校验使用
	Auditor          int    `gorm:"auditor" json:"auditor"`                       // 审核方：1审核平台，2发卡方
	VehicleBelong    int    `gorm:"vehicle_belong" json:"vehicle_belong"`         // 车辆归属：1个人车，2单位车
	IsTruck          int    `gorm:"is_truck" json:"is_truck"`                     // 是否货车：0否，1货车
	Uid              int    `gorm:"uid" json:"uid"`                               // 用户UID
	Status           int    `gorm:"status" json:"status"`                         // 状态：0待推送，10已推送，20审核，30审核失败，99已取消
	ThirdOrderSn     string `gorm:"third_order_sn" json:"third_order_sn"`         // 外部卡方签约相关：卡方签约号、绑定签约流水号、签约账号等
	IssuerOrderSn    string `gorm:"issuer_order_sn" json:"issuer_order_sn"`       // 外部卡方订单号
	IssuerAccount    string `gorm:"issuer_account" json:"issuer_account"`         // 外部卡方客户ID
	IssuerVehicleId  string `gorm:"issuer_vehicle_id" json:"issuer_vehicle_id"`   // 外部卡方车辆ID
	PlateId          string `gorm:"plate_id" json:"plate_id"`                     // 外部卡方车辆ID
	CardNo           string `gorm:"card_no" json:"card_no"`                       // 卡号
	ObuDeviceSn      string `gorm:"obu_device_sn" json:"obu_device_sn"`           // OBU号
	ActivateType     int    `gorm:"activate_type" json:"activate_type"`           // 激活类型1/2
	TruckType        int    `gorm:"truck_type" json:"truck_type"`                 // 货车车种：0普通，1运政车
	VehicleTruckType int    `gorm:"vehicle_truck_type" json:"vehicle_truck_type"` // 货车类型：0非货车  1普通货车 2牵引车集装箱
	VehicleType      int    `gorm:"vehicle_type" json:"vehicle_type"`             // 车型(国标)：0我记录，1一型客车  2二型客车 3三型客车 4四型客车 11一型货车 12二型货车 13三型货车 14四型货车 15五型货车 16六型货车 21一型专项作业车 22二型专项作业车 23三型专项作业车 24四型专项作业车 25五型专项作业车 26六型专项作业车
	VehicleSubType   int    `gorm:"vehicle_sub_type" json:"vehicle_sub_type"`     // 细分车型[101-普通客车 102-出租车 201-普通货车 202-牵引车 301-普通专项车]
	BusinessScope    int    `gorm:"business_scope" json:"business_scope"`         // 牵引车经营范围 0-无经营范围 1 -非货物专用运输车辆 2-专用集装箱运输车辆 3-混用集装箱运输车辆
	UseCharacterType int    `gorm:"use_character_type" json:"use_character_type"` // 使用性质:1 -营运 2-非营运 3-普通货运 4-专用集装箱 5-混用集装箱
	InstallType      int    `gorm:"install_type" json:"install_type"`             // 安装方式 [0-默认 1-前装]
	ApplyAt          time.Time `gorm:"apply_at" json:"apply_at"`                     // 申请时间
	AuditedAt        time.Time `gorm:"audited_at" json:"audited_at"`                 // 审核时间
	ShippedAt        time.Time `gorm:"shipped_at" json:"shipped_at"`                 // 发货时间
	ReceivedAt       time.Time `gorm:"received_at" json:"received_at"`               // 签收时间
	CanceledAt       time.Time `gorm:"canceled_at" json:"canceled_at"`               // 取消时间
	CreatedAt        time.Time `gorm:"created_at" json:"created_at"`                 // 创建时间
	UpdatedAt        time.Time `gorm:"updated_at" json:"updated_at"`                 // 更新时间
}

// Validate the fields.
func (u *IssuerOrderModel) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

// Token represents a JSON web token.
type Token struct {
	Token string `json:"token"`
}