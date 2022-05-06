package models

// BuyerCartAddParam Buyer 购物车添加参数模型
type BuyerCartAddParam struct {
	ProductId    uint   `form:"productId"`
	ProductCount uint   `form:"productCount"`
	UserId       uint64 `form:"userId"`
}

// BuyerCartDeleteParam Buyer 购物车删除参数模型
type BuyerCartDeleteParam struct {
	ProductId string `form:"productId"`
	UserId    uint64 `form:"userId"`
}

// BuyerCartClearParam Buyer 购物车清除参数模型
type BuyerCartClearParam struct {
	UserId uint64 `form:"userId"`
}

// BuyerCartQueryParam Buyer 购物车信息查询参数模型
type BuyerCartQueryParam struct {
	ProductId uint   `form:"productId"`
	UserId    uint64 `form:"userId"`
}

// BuyerCartItem Buyer 购物车商品项传输模型
type BuyerCartItem struct {
	Id        uint64  `gorm:"primaryKey" json:"id"`
	MainImage string  `gorm:"image_url"  json:"mainImage"`
	Title     string  `gorm:"title"      json:"title"`
	Price     float64 `gorm:"price"      json:"price"`
	Count     int     `gorm:"count"      json:"count"`
}

// BuyerCartInfo 购物车信息传输模型
type BuyerCartInfo struct {
	CartItem   []BuyerCartItem `json:"cartItem"`
	TotalPrice float64         `json:"totalPrice"`
}
