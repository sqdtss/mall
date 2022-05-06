package api

import (
	"github.com/gin-gonic/gin"
	"mall/models"
	"mall/response"
	"mall/service"
)

var browse service.BrowseService

// BuyerAddBrowseRecord Buyer 添加浏览记录
func BuyerAddBrowseRecord(c *gin.Context) {
	var param models.BuyerBrowseSaveParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	if count := browse.Add(param); count > 0 {
		response.Success("保存成功", count, c)
		return
	}
	response.Failed("保存失败", c)
}

// BuyerDeleteBrowseRecord Buyer 删除浏览记录
func BuyerDeleteBrowseRecord(c *gin.Context) {
	var param models.BuyerBrowseDeleteParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	if count := browse.Delete(param); count > 0 {
		response.Success("删除成功", count, c)
		return
	}
	response.Failed("删除失败", c)
}

// BuyerGetBrowseRecordList Buyer 获取浏览记录列表
func BuyerGetBrowseRecordList(c *gin.Context) {
	var param models.BuyerBrowseQueryParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	browseRecordList := browse.GetList(param)
	response.Success("查询成功", browseRecordList, c)
}
