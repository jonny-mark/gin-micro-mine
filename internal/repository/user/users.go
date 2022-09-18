/**
 * @author jiangshangfang
 * @date 2022/4/27 4:36 PM
 **/
package user

import (
	"gin-micro-mine/pkg/storage/orm"
	"gorm.io/gorm"
	"gin/internal/model/user"
	"errors"
	"gin/internal/constant"
)

// 根据token获取正常的用户信息
func FindValidOneByToken(token string) (userInfo *user.UsersModel, err error) {
	err = orm.DB.Where(&user.UsersModel{Token: token, Status: uint8(constant.StatusNormal)}).Last(&userInfo).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	user.User = userInfo
	return
}
