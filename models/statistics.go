package models

// SellerDataOverviewInfo Seller 数据总览统计传输模型
type SellerDataOverviewInfo struct {
	GoodsCount int     `json:"goodsCount"`
	OrderCount int     `json:"orderCount"`
	Amount     float64 `json:"amount"`
}

// SellerTodayOrderInfo Seller 今日订单数据统计传输模型
type SellerTodayOrderInfo struct {
	Data [5]int `json:"data"`
}

// SellerWeekOrderInfo Seller 本周数据总览统计传输模型
type SellerWeekOrderInfo struct {
	Orders [7]int     `json:"orders"`
	Amount [7]float64 `json:"amount"`
}
