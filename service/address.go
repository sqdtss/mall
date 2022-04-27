package service

import (
	"fmt"
	"mall/common"
	"mall/global"
	"mall/models"
)

type AddressService struct {
}

// Add Buyer 添加地址
func (*AddressService) Add(param models.BuyerAddressAddParam) int64 {
	address := models.Address{
		UserId:          param.UserId,
		Name:            param.Name,
		Mobile:          param.Mobile,
		PostalCode:      param.PostalCode,
		Province:        param.Province,
		City:            param.City,
		District:        param.District,
		DetailedAddress: param.DetailedAddress,
		IsDefault:       param.IsDefault,
		Created:         common.NowTime(),
	}
	if address.IsDefault == 1 {
		var addressId uint
		row := global.Db.Table("address").Select("id").
			Where("is_default = ? and user_id = ?", 1, address.UserId).Take(&addressId).RowsAffected
		if row > 0 {
			global.Db.Table("address").Where("id = ?", addressId).
				Update("is_default", 2)
			return global.Db.Create(&address).RowsAffected
		}
		return global.Db.Create(&address).RowsAffected
	}
	return global.Db.Create(&address).RowsAffected
}

// Delete Buyer 删除地址
func (*AddressService) Delete(id uint) int64 {
	return global.Db.Delete(&models.Address{}, id).RowsAffected
}

// Update Buyer 更新地址
func (*AddressService) Update(param models.BuyerAddressUpdateParam) int64 {
	address := models.Address{
		Id:              param.Id,
		UserId:          param.UserId,
		Name:            param.Name,
		Mobile:          param.Mobile,
		PostalCode:      param.PostalCode,
		Province:        param.Province,
		City:            param.City,
		District:        param.District,
		DetailedAddress: param.DetailedAddress,
		IsDefault:       param.IsDefault,
		Updated:         common.NowTime(),
	}
	if address.IsDefault == 1 {
		var addressId uint
		row := global.Db.Table("address").Select("id").
			Where("is_default = ? and user_id = ?", 1, address.UserId).Take(&addressId).RowsAffected
		fmt.Println(addressId)
		if row > 0 {
			global.Db.Table("address").Where("id = ?", addressId).
				Update("is_default", 2)
			return global.Db.Updates(&address).RowsAffected
		}
		return global.Db.Updates(&address).RowsAffected
	}
	return global.Db.Updates(&address).RowsAffected
}

// GetId Buyer 获取地址Id
func (*AddressService) GetId(uid string) uint64 {
	var id uint64
	global.Db.Table("address").Select("id").
		Where("is_default = ? and user_id = ?", 1, uid).Take(&id)
	return id
}

// GetInfo Buyer 获取更新信息
func (*AddressService) GetInfo(id uint) models.BuyerAddressUpdateInfo {
	var updateInfo models.BuyerAddressUpdateInfo
	global.Db.Table("address").First(&updateInfo, id)
	return updateInfo
}

// GetList Buyer 获取地址列表
func (*AddressService) GetList(uid string) []models.BuyerAddressList {
	var aList []models.BuyerAddressList
	global.Db.Table("address").Where("user_id = ?", uid).Order("is_default").Find(&aList)
	return aList
}
