package api

import (
	"github.com/gin-gonic/gin"
	"mall/models"
	"mall/response"
	"mall/service"
)

var sellerOrder service.SellerOrderService
var buyerOrder service.BuyerOrderService

// SellerDeleteOrder Seller 删除订单
func SellerDeleteOrder(c *gin.Context) {
	var param models.SellerOrderDeleteParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	if count := sellerOrder.Delete(param); count > 0 {
		response.Success("删除成功", count, c)
		return
	}
	response.Failed("删除失败", c)
}

// SellerUpdateOrder Seller 更新订单
func SellerUpdateOrder(c *gin.Context) {
	var param models.SellerOrderUpdateParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	if count := sellerOrder.Update(param); count > 0 {
		response.Success("更新成功", count, c)
		return
	}
	response.Failed("更新失败", c)
}

// SellerGetOrderList Seller 获取订单列表
func SellerGetOrderList(c *gin.Context) {
	var param models.SellerOrderListParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	productList, rows := sellerOrder.GetList(param)
	response.SuccessPage("查询成功", productList, rows, c)
}

// SellerGetOrderDetail Seller 获取订单详情
func SellerGetOrderDetail(c *gin.Context) {
	var param models.SellerOrderDetailParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	productDetail := sellerOrder.GetDetail(param)
	response.Success("查询成功", productDetail, c)
}

// BuyerCreateOrder Buyer 提交订单
func BuyerCreateOrder(c *gin.Context) {
	var param models.BuyerOrderSubmitParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	if count := buyerOrder.Submit(param); count > 0 {
		response.Success("提交成功", count, c)
		return
	}
	response.Failed("提交失败", c)
}

// BuyerGetOrderList Buyer 获取订单列表信息
func BuyerGetOrderList(c *gin.Context) {
	var param models.BuyerOrderQueryParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	orderList := buyerOrder.GetList(param)
	response.Success("查询成功", orderList, c)
}
