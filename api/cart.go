package api

import (
	"github.com/gin-gonic/gin"
	"mall/models"
	"mall/response"
	"mall/service"
)

var cart service.CartService

// BuyerAddCart Buyer 添加购物车
func BuyerAddCart(c *gin.Context) {
	var param models.BuyerCartAddParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	if cart.Add(param) {
		response.Success("添加成功", 1, c)
		return
	}
	response.Failed("已添加", c)
}

// BuyerDeleteCart Buyer 删除购物车中的商品
func BuyerDeleteCart(c *gin.Context) {
	var param models.BuyerCartDeleteParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	if count := cart.Delete(param); count > 0 {
		response.Success("删除成功", count, c)
		return
	}
	response.Failed("删除失败", c)
}

// BuyerClearCart Buyer 清除购物车中的商品
func BuyerClearCart(c *gin.Context) {
	var param models.BuyerCartClearParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	if count := cart.Clear(param); count > 0 {
		response.Success("清除成功", count, c)
		return
	}
	response.Failed("清除失败", c)
}

// BuyerGetCartInfo Buyer 获取购物车中信息
func BuyerGetCartInfo(c *gin.Context) {
	var param models.BuyerCartQueryParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	info := cart.GetInfo(param)
	if len(info.CartItem) == 0 {
		response.Success("购物车竟然是空的", info, c)
	}
	response.Success("查询成功", info, c)
}
