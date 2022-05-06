package models

// BuyerCollectionAddParam Buyer 商品收藏添加参数模型
type BuyerCollectionAddParam struct {
	ProductId uint64 `form:"productId" binding:"required"`
	UserId    uint64 `form:"userId" binding:"required"`
}

// BuyerCollectionDeleteParam Buyer 商品收藏删除参数模型
type BuyerCollectionDeleteParam struct {
	UserId uint64 `form:"userId" binding:"required"`
}

// BuyerCollectionQueryParam Buyer 商品收藏查询参数模型
type BuyerCollectionQueryParam struct {
	UserId uint64 `form:"userId" binding:"required"`
}

// BuyerCollectionList Buyer 商品收藏列表传输模型
type BuyerCollectionList struct {
	Id        uint64  `json:"id"`
	MainImage string  `json:"mainImage"`
	Title     string  `json:"title"`
	Price     float64 `json:"price"`
}
