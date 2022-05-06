package api

import (
	"github.com/gin-gonic/gin"
	"mall/models"
	"mall/response"
	"mall/service"
)

var address service.AddressService

// BuyerAddAddress Buyer 添加地址
func BuyerAddAddress(c *gin.Context) {
	var param models.BuyerAddressAddParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	if count := address.Add(param); count > 0 {
		response.Success("添加成功", count, c)
		return
	}
	response.Failed("添加失败", c)
}

// BuyerDeleteAddress Buyer 删除地址
func BuyerDeleteAddress(c *gin.Context) {
	var key models.BuyerAddressDeleteParam
	if err := c.ShouldBind(&key); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	if count := address.Delete(key); count > 0 {
		response.Success("删除成功", count, c)
		return
	}
	response.Failed("删除失败", c)
}

// BuyerUpdateAddress Buyer 更新地址
func BuyerUpdateAddress(c *gin.Context) {
	var param models.BuyerAddressUpdateParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	count := address.Update(param)
	if count > 0 {
		response.Success("更新成功", count, c)
		return
	}
	response.Failed("更新失败", c)
}

// BuyerGetAddressUpdateInfo Buyer 获取地址更新信息
func BuyerGetAddressUpdateInfo(c *gin.Context) {
	var key models.BuyerAddressInfoParam
	if err := c.ShouldBind(&key); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	info := address.GetInfo(key.AddressId)
	response.Success("查询成功", info, c)
}

// BuyerGetAddressList Buyer 获取地址列表
func BuyerGetAddressList(c *gin.Context) {
	var param models.BuyerAddressListParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	addressList := address.GetList(param.UserId)
	response.Success("查询成功", addressList, c)
}
