package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mall/models"
	"mall/response"
	"mall/service"
)

var sellerProduct service.SellerProductService
var buyerProduct service.BuyerProductService

// SellerCreateProduct Seller 创建商品
func SellerCreateProduct(c *gin.Context) {
	var param models.SellerProductCreateParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	if count := sellerProduct.Create(param); count > 0 {
		response.Success("创建成功", count, c)
		return
	}
	response.Failed("创建失败", c)
}

// SellerDeleteProduct Seller 删除商品
func SellerDeleteProduct(c *gin.Context) {
	var param models.SellerProductDeleteParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	if count := sellerProduct.Delete(param); count > 0 {
		response.Success("删除成功", count, c)
		return
	}
	response.Failed("删除失败", c)
}

// SellerUpdateProduct Seller 更新商品
func SellerUpdateProduct(c *gin.Context) {
	var param models.SellerProductUpdateParam
	if err := c.ShouldBind(&param); err != nil {
		fmt.Println(err)
		response.Failed("请求参数无效", c)
		return
	}
	if count := sellerProduct.Update(param); count > 0 {
		response.Success("更新成功", count, c)
		return
	}
	response.Failed("更新失败", c)
}

// SellerUpdateProductStatus Seller 更新商品状态
func SellerUpdateProductStatus(c *gin.Context) {
	var param models.SellerProductStatusUpdateParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	if count := sellerProduct.UpdateStatus(param); count > 0 {
		response.Success("更新成功", count, c)
		return
	}
	response.Failed("更新失败", c)
}

// SellerGetProductInfo Seller 获取商品信息
func SellerGetProductInfo(c *gin.Context) {
	var param models.SellerProductInfoParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	productInfo := sellerProduct.GetInfo(param)
	response.Success("查询成功", productInfo, c)
}

// SellerGetProductList Seller 获取商品列表
func SellerGetProductList(c *gin.Context) {
	var param models.SellerProductListParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	productList, rows := sellerProduct.GetList(param)
	response.SuccessPage("查询成功", productList, rows, c)
}

// BuyerGetProductList Buyer 获取商品列表
func BuyerGetProductList(c *gin.Context) {
	var param models.BuyerProductListParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	productList := buyerProduct.GetList(param)
	response.Success("查询成功", productList, c)
}

// BuyerGetProductList Buyer 获取商品搜索列表
func BuyerGetProductSearchList(c *gin.Context) {
	var param models.BuyerProductSearchParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	productList := buyerProduct.GetSearchList(param)
	response.Success("搜索成功", productList, c)
}

// BuyerGetProductDetail Buyer 获取商品详情
func BuyerGetProductDetail(c *gin.Context) {
	var param models.BuyerProductDetailParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	productDetail := buyerProduct.GetDetail(param)
	response.Success("查询成功", productDetail, c)
}
