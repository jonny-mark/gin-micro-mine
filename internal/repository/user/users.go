package user

import (
	"errors"
	"github.com/jonny-mark/gin-micro-mine/internal/constant"
	"github.com/jonny-mark/gin-micro-mine/internal/model/user"
	"github.com/jonny-mark/gin-micro-mine/pkg/storage/orm"
	"gorm.io/gorm"
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
