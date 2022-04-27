/**
 * @author jiangshangfang
 * @date 2022/4/3 6:26 PM
 **/
package vehicle

import (
	"time"
	"github.com/go-playground/validator/v10"
)

// TableName 表名
func (u *VehiclesModel) TableName() string {
	return "etc_users_cards"
}

type VehiclesModel struct {
	ID               uint64     `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	Uid              uint       `gorm:"column:uid" json:"uid"`
	CardId           uint       `gorm:"column:card_id" json:"card_id"`                                                               // 卡种（对应 etc_cards.id）
	PlateNo          string     `gorm:"column:plate_no;type:varchar(32);not null;default:''" json:"plate_no"`                        // 车牌号（冗余）
	PlateColor       uint       `gorm:"column:plate_color;type:tinyint;not null;default:0" json:"plate_color"`                       // 车牌颜色：0、蓝色 1、黄色 2、黑色 3、白色 4、渐变绿色 5、黄绿双拼色 6、蓝白渐变色 7、临时牌照 9、未确定 11、绿色 12、红色
	CardNo           string     `gorm:"column:card_no" json:"card_no"`                                                               // ETC通行卡号
	ObuDeviceSn      string     `gorm:"column:obu_device_sn" json:"obu_device_sn"`                                                   // OBU设备电子标签号
	WarrantyService  int        `gorm:"column:warranty_service" json:"warranty_service"`                                             // 首次购买设备质保天数：0无服务，大于0有质保
	IssuerCode       string     `gorm:"column:issuer_code" json:"issuer_code"`                                                       // 卡方编码
	IssuerAccount    string     `gorm:"column:issuer_account" json:"issuer_account"`                                                 // 第三方系统的账号
	IssuerContractSn string     `gorm:"column:issuer_contract_sn" json:"issuer_contract_sn"`                                         // 卡方监管平台车辆签约流水号
	ApplyOrderSn     string     `gorm:"column:apply_order_sn" json:"apply_order_sn"`                                                 // 申请订单流水号
	IdCardInfoId     int        `gorm:"column:idcard_info_id;type:int unsigned;not null;default:0" json:"idcard_info_id"`            // 身份证信息 etc_users_idcards.id
	VehicleInfoId    int        `gorm:"column:vehicle_info_id;type:int unsigned;not null;default:0" json:"vehicle_info_id"`          // 行驶证信息 etc_users_vehicles.id
	EarnestMoney     float64    `gorm:"column:earnest_money;type:decimal(10,2) unsigned;not null;default:0.00" json:"earnest_money"` // 已付保证金
	SurplusMoney     float64    `gorm:"column:surplus_money;type:decimal(10,2);not null;default:0.00" json:"surplus_money"`          // 卡片余额（元）
	DebtAmount       float64    `gorm:"column:debt_amount;type:decimal(10,2);not null;default:0.00" json:"debt_amount"`              // 欠费总额（元）
	TollDebtAmount   float64    `gorm:"column:toll_debt_amount;type:decimal(10,2);not null;default:0.00" json:"toll_debt_amount"`    // 通行欠费（元）
	TollAmount       float64    `gorm:"column:toll_amount;type:decimal(10,2) unsigned;not null;default:0.00" json:"toll_amount"`     // 通行消费（元）
	TollTimes        int        `gorm:"column:toll_times;type:int unsigned;not null;default:0" json:"toll_times"`                    // 通行次数
	AllowReactivate  int        `gorm:"column:allow_reactivate;type:tinyint unsigned;not null;default:0" json:"allow_reactivate"`    // 是否允许重新激活
	FirstActivatedAt *time.Time `gorm:"column:first_activated_at;type:timestamp;default:null" json:"first_activated_at"`             // 首次激活时间，二次激活不更新
	ActivatedAt      *time.Time `gorm:"column:activated_at;type:timestamp;default:null" json:"activated_at"`                         // 激活时间
	SurplusUpdatedAt *time.Time `gorm:"column:surplus_updated_at;type:timestamp;default:null" json:"surplus_updated_at"`             // 卡片余额更新时间
	VoidAt           *time.Time `gorm:"column:void_at;type:timestamp;default:null" json:"void_at"`                                   // 发起注销时间
	EtcVoidAt        *time.Time `gorm:"column:etc_void_at;type:timestamp;default:null" json:"etc_void_at"`                           // 微行ETC注销时间
	IssuerVoidAt     *time.Time `gorm:"column:issuer_void_at;type:timestamp;default:null" json:"issuer_void_at"`                     // 发卡方注销时间
	FirstTollAt      *time.Time `gorm:"column:first_toll_at;type:timestamp;default:null" json:"first_toll_at"`                       // 首次通行时间
	IsDel            int        `gorm:"column:is_del;type:tinyint unsigned;not null;default:0" json:"is_del"`                        // 是否删除
	Status           int        `gorm:"column:status;type:tinyint unsigned;not null;default:0" json:"status"`                        // 卡片状态 0:正常 2:挂失 3:补办中 4:作废 8:注销中 9:已注销
	TollPayChannel   int        `gorm:"column:toll_pay_channel;type:tinyint unsigned;not null;default:1" json:"toll_pay_channel"`    // 通行扣费支付渠道 1、我方自己发起的扣费 2、米大师发起的扣费 3、招行代扣
	ServiceStatus    int        `gorm:"column:service_status;type:tinyint;not null;default:0" json:"service_status"`                 // 是否开启售后权限
	VoidStatus       int        `gorm:"column:void_status;type:tinyint;not null;default:0" json:"void_status"`                       // 是否允许用户注销 0:不允许 1:允许
	Manufacturer     int        `gorm:"column:manufacturer;type:tinyint unsigned;not null;default:0" json:"manufacturer"`            // 设备商类型（埃特斯1、金溢2、聚力3、万集4、成谷5）
	DeviceType       int        `gorm:"column:device_type;type:tinyint unsigned;not null;default:0" json:"device_type"`              // 设备类型[0-普通 1-可充电设备]
	BizInfoId        int        `gorm:"column:biz_info_id;type:int unsigned;not null;default:0" json:"biz_info_id"`                  // 营业执照信息 etc_users_biz_license.id
	ShowRefundBtn    int        `gorm:"column:show_refund_btn;type:tinyint unsigned;not null;default:0" json:"show_refund_btn"`      // 是否显示退款按钮
	ReviewOrderSn    string     `gorm:"column:review_order_sn;type:varchar(50)" json:"review_order_sn"`                              // 审核单号
	DebtVer          int        `gorm:"column:debt_ver;type:int;not null;default:0" json:"debt_ver"`                                 // 债务乐观锁
	SettleCycle      int        `gorm:"column:settle_cycle;type:tinyint;not null;default:1" json:"settle_cycle"`                     // 结算周期：1日结2周结3月结
	HasChangeSign    int        `gorm:"column:has_change_sign;type:tinyint(1);not null;default:0" json:"has_change_sign"`            // 是否已修改签约[0-未改签 1-已改签]
	IsFront          int        `gorm:"column:is_front;type:tinyint unsigned;not null;default:0" json:"is_front"`                    // 是否前装[0-默认 1-前装]
	CreatedAt        time.Time  `gorm:"column:created_at" json:"-"`
	UpdatedAt        time.Time  `gorm:"column:updated_at" json:"-"`

	VehiclesCards []VehiclesCardsModel `gorm:"foreignKey:VehicleId;associationForeignKey:ID" json:"vehicles_cards"`
	VehiclesObus  []VehiclesObusModel  `gorm:"foreignKey:VehicleId;associationForeignKey:ID" json:"vehicles_obus"`
}

// Validate the fields.
func (u *VehiclesModel) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
