/**
 * @author jiangshangfang
 * @date 2022/4/27 4:36 PM
 **/
package user

import (
	"gin/pkg/storage/orm"
	"gorm.io/gorm"
	"gin/internal/model/user"
	userConstant "gin/internal/constant/model/user"
	"errors"
)

// 根据token获取正常的用户信息
func FindValidOneByToken(token string) (userInfo *user.UsersModel, err error) {
	err = orm.DB.Where(&user.UsersModel{Token: token, Status: uint8(userConstant.STATUS_NORMAL)}).Last(&userInfo).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return
}
