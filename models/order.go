package models

// Order 数据库，订单数据映射模型
type Order struct {
	Id          uint64  `gorm:"primaryKey"`
	ProductItem string  `gorm:"product_item"`
	TotalPrice  float64 `gorm:"total_price"`
	Status      string  `gorm:"status"`
	AddressId   uint64  `gorm:"address_id"`
	UserId      string  `gorm:"user_id"`
	NickName    string  `gorm:"nick_name"`
	Created     string  `gorm:"created"`
	Updated     string  `gorm:"updated"`
}

// SellerOrderDeleteParam Seller 订单删除参数模型
type SellerOrderDeleteParam struct {
	Id uint64 `form:"id"`
}

// SellerOrderUpdateParam Seller 订单更新参数模型
type SellerOrderUpdateParam struct {
	Id     uint64 `json:"id"`
	Status string `json:"status"`
}

// SellerOrderListParam Seller 订单列表查询参数模型
type SellerOrderListParam struct {
	Page   Page
	Id     uint64 `form:"id"`
	Status string `form:"status"`
}

// SellerOrderDetailParam Seller 订单列表查询参数模型
type SellerOrderDetailParam struct {
	Id uint64 `form:"id"`
}

// SellerOrderList Seller 订单列表传输模型
type SellerOrderList struct {
	Id         uint64  `json:"id"`
	NickName   string  `json:"nickName"`
	Status     string  `json:"status"`
	TotalPrice float64 `json:"totalPrice"`
	Created    string  `json:"created"`
}

// SellerOrderDetail Seller 订单详情传输模型
type SellerOrderDetail struct {

	// 订单信息
	Id         uint64  `json:"id"`
	Created    string  `json:"created"`
	NickName   string  `json:"nickName"`
	Status     string  `json:"status"`
	TotalPrice float64 `json:"totalPrice"`

	// 地址信息
	Name            string `json:"name"`
	Mobile          string `json:"mobile"`
	PostalCode      int    `json:"postalCode"`
	Province        string `json:"province"`
	City            string `json:"city"`
	District        string `json:"district"`
	DetailedAddress string `json:"detailedAddress"`

	// 商品列表信息
	ProductItem []SellerProductItem `json:"productItem"`
}

// BuyerOrderSubmitParam Buyer 订单提交参数模型
type BuyerOrderSubmitParam struct {
	UserId   string `form:"userId"    json:"userId"`
	NickName string `form:"nickName"  json:"nickName"`
	Status   string `form:"status"    json:"status"`
}

// BuyerOrderQueryParam Buyer 订单查询参数模型
type BuyerOrderQueryParam struct {
	UserId string `form:"userId"    json:"userId"`
	Status string `form:"status"    json:"status"`
}

// BuyerOrderListInfo Buyer 订单列表传输模型
type BuyerOrderListInfo struct {
	Id          uint64             `json:"id"`
	Status      string             `json:"status"`
	TotalPrice  float64            `json:"totalPrice"`
	ProductItem []BuyerProductItem `json:"productItem"`
}
