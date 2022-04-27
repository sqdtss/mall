package api

import (
	"github.com/gin-gonic/gin"
	"mall/response"
	"mall/service"
)

var sellerStatistics service.SellerStatisticsService

// SellerGetDataOverviewInfo Seller 获取数据总览统计信息
func SellerGetDataOverviewInfo(c *gin.Context) {
	overviewInfo := sellerStatistics.GetDataOverviewInfo()
	response.Success("查询成功", overviewInfo, c)
}

// SellerGetTodayOrderDataInfo Seller 获取今日订单数据统计信息
func SellerGetTodayOrderDataInfo(c *gin.Context) {
	todayInfo := sellerStatistics.GetTodayDataInfo()
	response.Success("查询成功", todayInfo, c)
}

// SellerGetWeekDataInfo Seller 获取本周数据统计信息
func SellerGetWeekDataInfo(c *gin.Context) {
	weekInfo := sellerStatistics.GetWeekDataInfo()
	response.Success("查询成功", weekInfo, c)
}
