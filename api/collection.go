package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mall/models"
	"mall/response"
	"mall/service"
)

var collection service.CollectionService

// BuyerAddCollection Buyer 添加收藏的商品
func BuyerAddCollection(c *gin.Context) {
	var param models.BuyerCollectionAddParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	if count := collection.Add(param); count > 0 {
		response.Success("收藏成功", count, c)
		return
	}
	response.Failed("已收藏", c)
}

// BuyerDeleteCollection Buyer 删除收藏的商品
func BuyerDeleteCollection(c *gin.Context) {
	var param models.BuyerCollectionDeleteParam
	if err := c.ShouldBind(&param); err != nil {
		fmt.Println(err)
		response.Failed("请求参数无效", c)
		return
	}
	if count := collection.Delete(param); count > 0 {
		response.Success("清除成功", count, c)
		return
	}
	response.Failed("清除失败", c)
}

// BuyerGetCollectionList Buyer 获取收藏的商品列表
func BuyerGetCollectionList(c *gin.Context) {
	var param models.BuyerCollectionQueryParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	info := collection.GetList(param)
	response.Success("查询成功", info, c)
}
