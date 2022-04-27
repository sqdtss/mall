package api

import (
	"github.com/gin-gonic/gin"
	"mall/models"
	"mall/response"
	"mall/service"
)

var sellerCategory service.SellerCategoryService
var buyerCategory service.BuyerCategoryService

// SellerCreateCategory Seller 创建类目
func SellerCreateCategory(c *gin.Context) {
	var param models.SellerCategoryCreateParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	if count := sellerCategory.Create(param); count > 0 {
		response.Success("创建成功", count, c)
		return
	}
	response.Failed("创建失败", c)
}

// SellerDeleteCategory Seller 删除类目
func SellerDeleteCategory(c *gin.Context) {
	var param models.SellerCategoryDeleteParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	if count := sellerCategory.Delete(param); count > 0 {
		response.Success("删除成功", count, c)
		return
	}
	response.Failed("删除失败", c)
}

// SellerUpdateCategory Seller 更新类目
func SellerUpdateCategory(c *gin.Context) {
	var param models.SellerCategoryUpdateParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	if count := sellerCategory.Update(param); count > 0 {
		response.Success("更新成功", count, c)
		return
	}
	response.Failed("更新失败", c)
}

// SellerGetCategoryList Seller 获取类目列表
func SellerGetCategoryList(c *gin.Context) {
	var param models.SellerCategoryQueryParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	cateList, rows := sellerCategory.GetList(param)
	response.SuccessPage("查询成功", cateList, rows, c)
}

// SellerGetCategoryOption Seller 获取类目选项
func SellerGetCategoryOption(c *gin.Context) {
	option := sellerCategory.GetOption()
	response.Success("查询成功", option, c)
}

// BuyerGetCategoryOption Buyer 获取类目选项
func BuyerGetCategoryOption(c *gin.Context) {
	option := buyerCategory.GetOption()
	response.Success("查询成功", option, c)
}
