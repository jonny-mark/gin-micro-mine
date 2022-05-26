/**
 * @author jiangshangfang
 * @date 2022/4/27 3:29 PM
 **/
package user

import (
	"github.com/go-playground/validator/v10"
	"time"
)

// TableName 表名
func (u *UsersModel) TableName() string {
	return "etc_users"
}

var User *UsersModel

type UsersModel struct {
	ID           uint64    `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	Token        string    `gorm:"index:token;column:token;type:varchar(50);not null;default:''" json:"token"`
	ClientType   uint      `gorm:"column:client_type;type:tinyint unsigned;not null;default:1" json:"client_type"`
	ThirdType    uint      `gorm:"column:third_type;type:int unsigned;not null;default:0" json:"third_type"`
	ThirdCode    string    `gorm:"column:third_code;type:varchar(255);not null;default:''" json:"third_code"`
	WxUnionid    string    `gorm:"unique;column:wx_unionid;type:varchar(100);not null;default:''" json:"wx_unionid"`        //微信 unionid
	MinaOpenid   string    `gorm:"index:m_oid;column:mina_openid;type:varchar(100);not null;default:''" json:"mina_openid"` //小程序 openid
	MpOpenid     string    `gorm:"index:mp_openid;column:mp_openid;type:varchar(100);not null;default:''" json:"mp_openid"` //公众服务号 openid
	HasCard      uint      `gorm:"column:has_card;type:tinyint unsigned;not null;default:0" json:"has_card"`              //是否有已激活的卡
	ApplyOrderSn string    `gorm:"column:apply_order_sn;type:varchar(50);not null;default:''" json:"apply_order_sn"`        //申办订单号
	Nickname     string    `gorm:"column:nickname;type:varchar(50);not null;default:''" json:"nickname"`                    // 昵称
	Gender       uint8     `gorm:"column:gender;type:tinyint unsigned;not null;default:3" json:"gender"`                    // 性别 1:男 2:女 3:未知
	Province     string    `gorm:"column:province;type:varchar(50);not null;default:''" json:"province"`                    // 省份
	City         string    `gorm:"column:city;type:varchar(50);not null;default:''" json:"city"`                            // 城市
	Country      string    `gorm:"column:country;type:varchar(50);not null;default:''" json:"country"`                      // 国家
	AvatarURL    string    `gorm:"column:avatar_url;type:varchar(191);not null;default:''" json:"avatar_url"`               // 头像
	Phone        string    `gorm:"index:phone;column:phone;type:varchar(50);not null;default:''" json:"phone"`              // 联系电话
	SessionKey   string    `gorm:"column:session_key;type:varchar(191);not null;default:''" json:"session_key"`             // 微信用户数据签名私钥
	IsTester     uint8     `gorm:"column:is_tester;type:tinyint unsigned;not null;default:0" json:"is_tester"`              // 是否测试人员
	Status       uint8     `gorm:"column:status;type:tinyint unsigned;not null;default:1" json:"status"`                    // 用户是否可用，1:可用，0:不可用
	MpStatus     uint8     `gorm:"column:mp_status;type:tinyint unsigned;not null;default:0" json:"mp_status"`              // 服务号关注状态，0：未关注 1：已关注 2：取消关注
	Age          int       `gorm:"column:age;type:int;default:null" json:"age"`
	Name         string    `gorm:"column:name;type:varchar(255);default:null" json:"name"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"-"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"-"`
}

// Validate the fields.
func (u *UsersModel) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
