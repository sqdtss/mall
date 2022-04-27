package service

import (
	"mall/common"
	"mall/global"
	"mall/models"
	"strconv"
	"strings"
)

var cart CartService
var address AddressService

type SellerOrderService struct {
}

type BuyerOrderService struct {
}

// Delete Seller 删除订单
func (*SellerOrderService) Delete(param models.SellerOrderDeleteParam) int64 {
	return global.Db.Delete(&models.Order{}, param.Id).RowsAffected
}

// Update Seller 更新订单
func (*SellerOrderService) Update(param models.SellerOrderUpdateParam) int64 {
	order := models.Order{
		Id:      param.Id,
		Status:  param.Status,
		Updated: common.NowTime(),
	}
	return global.Db.Model(&order).Updates(order).RowsAffected
}

// GetList Seller 获取订单列表
func (*SellerOrderService) GetList(param models.SellerOrderListParam) ([]models.SellerOrderList, int64) {
	var orderList []models.SellerOrderList
	query := &models.Order{
		Id:     param.Id,
		Status: param.Status,
	}
	rows := common.RestPage(param.Page, "order", query, &orderList, &[]models.Order{})
	return orderList, rows
}

// GetDetail Seller 获取订单详情
func (*SellerOrderService) GetDetail(param models.SellerOrderDetailParam) (od models.SellerOrderDetail) {
	var order models.Order
	var address models.Address
	var productItem []models.SellerProductItem

	// 查询订单信息与地址信息
	global.Db.First(&order, param.Id)
	global.Db.First(&address, order.AddressId)

	// 查询订单中包含的商品信息
	idList := strings.Split(order.ProductItem, ",")
	var productIdList []uint64
	for _, id := range idList {
		pid, _ := strconv.Atoi(id)
		if pid != 0 {
			productIdList = append(productIdList, uint64(pid))
		}
	}
	global.Db.Table("product").Find(&productItem, productIdList)
	orderDetail := models.SellerOrderDetail{
		Id:              order.Id,
		Created:         order.Created,
		NickName:        order.NickName,
		Status:          order.Status,
		TotalPrice:      order.TotalPrice,
		Name:            address.Name,
		Mobile:          address.Mobile,
		PostalCode:      address.PostalCode,
		Province:        address.Province,
		City:            address.City,
		District:        address.District,
		DetailedAddress: address.DetailedAddress,
		ProductItem:     productItem,
	}
	return orderDetail
}

// Submit Buyer 提交订单
func (*BuyerOrderService) Submit(param models.BuyerOrderSubmitParam) int64 {
	info := cart.GetInfo(models.BuyerCartQueryParam{UserId: param.UserId})
	var pids []string
	for _, item := range info.CartItem {
		pids = append(pids, strconv.Itoa(int(item.Id)))
	}
	pidsItem := strings.Join(pids, ",")
	order := models.Order{
		ProductItem: pidsItem,
		TotalPrice:  info.TotalPrice,
		Status:      param.Status,
		AddressId:   address.GetId(param.UserId),
		UserId:      param.UserId,
		NickName:    param.NickName,
		Created:     common.NowTime(),
	}
	return global.Db.Create(&order).RowsAffected
}

// GetList Buyer 获取订单列表
func (*BuyerOrderService) GetList(param models.BuyerOrderQueryParam) []models.BuyerOrderListInfo {
	// 查询满足特定条件的订单
	var orderList []models.Order
	query := &models.Order{
		UserId: param.UserId,
		Status: param.Status,
	}
	global.Db.Table("order").Where(query).Find(&orderList)
	var orderInfoList []models.BuyerOrderListInfo
	for _, order := range orderList {
		var orderInfo models.BuyerOrderListInfo
		var productItem []models.BuyerProductItem
		orderInfo.Id = order.Id
		orderInfo.Status = order.Status
		orderInfo.TotalPrice = order.TotalPrice

		// 查询商品信息
		sList := strings.Split(order.ProductItem, ",")
		var idList []uint
		for _, pid := range sList {
			id, _ := strconv.Atoi(pid)
			if id != 0 {
				i := uint(id)
				idList = append(idList, i)
			}
		}
		global.Db.Table("product").Find(&productItem, idList)
		orderInfo.ProductItem = productItem
		orderInfoList = append(orderInfoList, orderInfo)
	}
	return orderInfoList
}
