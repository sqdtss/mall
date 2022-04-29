package service

import (
	"mall/common"
	"mall/global"
	"mall/models"
)

type SellerService struct {
}

type BuyerService struct {
}

// SellerLogin Seller 用户登录信息验证
func (*SellerService) SellerLogin(param models.SellerLoginParam) uint64 {
	var seller models.Seller
	global.Db.Where("username = ? and password = ?", param.Username, param.Password).First(&seller)
	return seller.Id
}

// SellerRegister Seller 注册
func (*SellerService) SellerRegister(param models.SellerRegisterParam) uint64 {
	var seller models.Seller
	result := global.Db.Where("username = ?", param.Username).First(&seller)
	if result.RowsAffected > 0 {
		return seller.Id
	}
	seller = models.Seller{
		Username: param.Username,
		Password: param.Password,
		Status:   param.Status,
		Created:  common.NowTime(),
	}
	global.Db.Create(&seller)
	return seller.Id
}

// BuyerLogin Buyer 用户登录信息验证
func (*BuyerService) BuyerLogin(param models.BuyerLoginParam) uint64 {
	var buyer models.Buyer
	global.Db.Where("username = ? and password = ?", param.Username, param.Password).First(&buyer)
	return buyer.Id
}

// BuyerRegister Buyer 注册
func (*BuyerService) BuyerRegister(param models.BuyerRegisterParam) uint64 {
	var buyer models.Buyer
	result := global.Db.Where("username = ?", param.Username).First(&buyer)
	if result.RowsAffected > 0 {
		return buyer.Id
	}
	buyer = models.Buyer{
		Username: param.Username,
		Password: param.Password,
		Status:   param.Status,
		Created:  common.NowTime(),
	}
	global.Db.Create(&buyer)
	return buyer.Id
}
