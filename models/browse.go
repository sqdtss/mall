package models

// BuyerBrowseSaveParam Buyer 浏览记录保存参数模型
type BuyerBrowseSaveParam struct {
	ProductId uint64 `form:"productId" binding:"required"`
	UserId    string `form:"userId" binding:"required"`
}

// BuyerBrowseDeleteParam Buyer 浏览记录删除参数模型
type BuyerBrowseDeleteParam struct {
	ProductIds []string `form:"productIds" binding:"required"`
	UserId     string   `form:"userId" binding:"required"`
}

// BuyerBrowseQueryParam Buyer 浏览记录查询参数模型
type BuyerBrowseQueryParam struct {
	UserId string `form:"userId" binding:"required"`
}

// BuyerBrowseItem Buyer 浏览记录列表商品项传输模型
type BuyerBrowseItem struct {
	Id        uint64  `json:"id"`
	MainImage string  `json:"mainImage"`
	Title     string  `json:"title"`
	Price     float64 `json:"price"`
}

// BuyerBrowseRecordList Buyer 浏览记录列表传输模型
type BuyerBrowseRecordList struct {
	BrowseTime string          `json:"browseTime"`
	BrowseItem BuyerBrowseItem `json:"browseItem"`
}
