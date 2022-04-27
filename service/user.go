package service

import (
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

// BuyerLogin Buyer 用户登录信息验证
func (*BuyerService) BuyerLogin(param models.BuyerLoginParam) uint64 {
	var buyer models.Buyer
	global.Db.Where("username = ? and password = ?", param.Username, param.Password).First(&buyer)
	return buyer.Id
}
