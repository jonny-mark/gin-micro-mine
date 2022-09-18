/**
 * @author jiangshangfang
 * @date 2022/4/3 6:26 PM
 **/
package vehicle

import (
	"gin/internal/model/vehicle"
	"gin-micro-mine/pkg/storage/orm"
	"github.com/go-errors/errors"
	"gorm.io/gorm"
)

//UsersCardsModel ...
type UsersCardsModel struct{}

// 通过plateNo、plateColor找到最后一条未被删除的用户数据
func FindValidOneByPlate(plateNo string, plateColor int8) (vehicleInfo *vehicle.VehiclesModel, err error) {
	err = orm.DB.Where(&vehicle.VehiclesModel{PlateNo: plateNo, PlateColor: plateColor, IsDel: 0}).Last(&vehicleInfo).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return
}

// 通过cardNo找到最后一条未被删除的用户数据
func FindValidOneByCardNo(cardNo string) (vehicleInfo *vehicle.VehiclesModel, err error) {
	err = orm.DB.Where(&vehicle.VehiclesModel{CardNo: cardNo, IsDel: 0}).Last(&vehicleInfo).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return
}

// 通过obuNo找到最后一条未被删除的用户数据
func FindValidOneByObuNo(obuNo string) (vehicleInfo *vehicle.VehiclesModel, err error) {
	err = orm.DB.Where(&vehicle.VehiclesModel{ObuDeviceSn: obuNo, IsDel: 0}).Last(&vehicleInfo).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return
}

// 通过uid、plateNo、plateColor找到最后一条未被删除的用户数据
func FindValidOneByUidAndPlate(uid uint, plateNo string, plateColor int8) (vehicleInfo *vehicle.VehiclesModel, err error) {
	err = orm.DB.Where(&vehicle.VehiclesModel{Uid: uid, PlateNo: plateNo, PlateColor: plateColor, IsDel: 0}).Last(&vehicleInfo).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return
}

// 通过uid、vehicleId找到最后一条未被删除的用户数据
func FindValidOneByUidAndVehicleId(uid uint, vehicleId uint) (vehicleInfo *vehicle.VehiclesModel, err error) {
	err = orm.DB.Where("uid = ? and id = ? and id_del = ?", uid, vehicleId, 0).Last(&vehicleInfo).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return
}

// 通过vehicleId找到最后一条用户数据
func FindOneByVehicleId(vehicleId uint) (vehicleInfo *vehicle.VehiclesModel, err error) {
	err = orm.DB.Where("id = ?", vehicleId).Last(&vehicleInfo).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return
}

// 通过uid、cardNo找到最后一条用户数据
func FindOneByUidAndCardNo(uid uint, cardNo string) (vehicleInfo *vehicle.VehiclesModel, err error) {
	err = orm.DB.Where(&vehicle.VehiclesModel{Uid: uid, CardNo: cardNo}).Last(&vehicleInfo).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return
}

// 通过uid、cardNo找到最后一条未被删除的用户数据
func FindValidOneByUidAndCardNo(uid uint, cardNo string) (vehicleInfo *vehicle.VehiclesModel, err error) {
	err = orm.DB.Where(&vehicle.VehiclesModel{Uid: uid, CardNo: cardNo, IsDel: 0}).Last(&vehicleInfo).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return
}

// 通过uid、cardId找到最后一条未被删除的用户数据
func FindValidOneByUidAndCardId(uid, cardId uint) (vehicleInfo *vehicle.VehiclesModel, err error) {
	err = orm.DB.Where(&vehicle.VehiclesModel{Uid: uid, CardId: cardId, IsDel: 0}).Last(&vehicleInfo).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return
}

// 通过uid、cardId找到所有未被删除的用户数据
func FindValidAllByUidAndCardId(uid, cardId uint) (vehiclesInfo []vehicle.VehiclesModel, err error) {
	err = orm.DB.Where(&vehicle.VehiclesModel{Uid: uid, CardId: cardId, IsDel: 0}).Find(&vehiclesInfo).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return
}

// 通过uid、plateNo、plateColor、cardNo找到最后一条未被删除的用户数据然后更新
func ModifyValidOneByUserAndPlate(uid uint, plateNo string, plateColor int8, cardNo string, vehicleInfo *vehicle.VehiclesModel) (err error) {
	err = orm.DB.Where(&vehicle.VehiclesModel{Uid: uid, PlateNo: plateNo, PlateColor: plateColor, CardNo: cardNo, IsDel: 0}).Updates(vehicleInfo).Error
	return
}

// 通过users_cards的id查找名下关联的有效的所以卡和签
func FindAssociatedVehicleCardsByValidPlate(uid uint, plateNo string, plateColor int8) (vehicleInfo *vehicle.VehiclesModel, err error) {
	vehicleInfo, err = FindValidOneByUidAndPlate(uid, plateNo, plateColor)
	if err != nil {
		return nil, err
	}
	orm.DB.Preload("VehiclesCards").Preload("VehiclesObus").Find(vehicleInfo)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return
}

// 通过users_cards的id查找名下有效的车辆发行的所以卡
func FindVehicleCardsByValidPlate(uid uint, plateNo string, plateColor int8) (cardsInfo []vehicle.VehiclesCardsModel, err error) {
	var vehicleInfo *vehicle.VehiclesModel
	vehicleInfo, err = FindValidOneByUidAndPlate(uid, plateNo, plateColor)
	if err != nil {
		return nil, err
	}
	orm.DB.Model(vehicleInfo).Association("VehiclesCards").Find(&cardsInfo)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return
}

// 通过users_cards的id查找名下有效的车辆发行的所以签
func FindVehicleObusByValidPlate(uid uint, plateNo string, plateColor uint) (obusInfo []vehicle.VehiclesObusModel, err error) {
	var vehicleInfo *vehicle.VehiclesModel
	vehicleInfo, err = FindValidOneByUidAndPlate(uid, plateNo, plateColor)
	if err != nil {
		return nil, err
	}
	orm.DB.Model(vehicleInfo).Association("VehiclesObus").Find(&obusInfo)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return
}