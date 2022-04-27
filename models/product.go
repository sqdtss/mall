package models

// 数据库，商品数据映射模型
type Product struct {
	Id                uint64  `gorm:"primaryKey"         json:"id"`
	CategoryId        uint64  `gorm:"category_id"        json:"categoryId"`
	Title             string  `gorm:"title"              json:"title"`
	Description       string  `gorm:"description"        json:"description"`
	Price             float64 `gorm:"price"              json:"price"`
	Amount            uint    `gorm:"amount"             json:"amount"`
	Sales             uint    `gorm:"sales"              json:"sales"`
	MainImage         string  `gorm:"main_image"         json:"mainImage"`
	Delivery          string  `gorm:"delivery"           json:"delivery"`
	Assurance         string  `gorm:"assurance"          json:"assurance"`
	Name              string  `gorm:"name"               json:"name"`
	Weight            float32 `gorm:"weight"             json:"weight"`
	Brand             string  `gorm:"brand"              json:"brand"`
	Origin            string  `gorm:"origin"             json:"origin"`
	ShelfLife         uint    `gorm:"shelf_life"         json:"shelfLife"`
	NetWeight         float32 `gorm:"net_weight"         json:"netWeight"`
	UseWay            string  `gorm:"use_way"            json:"useWay"`
	PackingWay        string  `gorm:"packing_way"        json:"packingWay"`
	StorageConditions string  `gorm:"storage_conditions" json:"storageConditions"`
	DetailImage       string  `gorm:"detail_image"       json:"detailImage"`
	Status            uint    `gorm:"status"             json:"status"`
	Created           string  `gorm:"created"            json:"created"`
	Updated           string  `gorm:"updated"            json:"updated"`
}

// Seller 商品创建参数模型
type SellerProductCreateParam struct {
	CategoryId        uint64  `json:"categoryId"         binding:"required,gt=0"`
	Title             string  `json:"title"              binding:"required"`
	Description       string  `json:"description"`
	Price             float64 `json:"price"              binding:"required,gt=0"`
	Amount            uint    `json:"amount"             binding:"required,gt=0"`
	MainImage         string  `json:"mainImage"          binding:"required"`
	Delivery          string  `json:"delivery"           binding:"required"`
	Assurance         string  `json:"assurance"          binding:"required"`
	Name              string  `json:"name"               binding:"required"`
	Weight            float32 `json:"weight"             binding:"required"`
	Brand             string  `json:"brand"`
	Origin            string  `json:"origin"`
	ShelfLife         uint    `json:"shelfLife"          binding:"required,gt=0"`
	NetWeight         float32 `json:"netWeight"          binding:"required,gt=0"`
	UseWay            string  `json:"useWay"`
	PackingWay        string  `json:"packingWay"`
	StorageConditions string  `json:"storageConditions"`
	DetailImage       string  `json:"detailImage"        binding:"required"`
	Status            uint    `json:"status"             binding:"required,oneof=1 2"`
}

// Seller 商品删除参数模型
type SellerProductDeleteParam struct {
	Id uint64 `form:"id" binding:"required,gt=0"`
}

// Seller 商品更新参数模型
type SellerProductUpdateParam struct {
	Id                uint64  `json:"id"                 binding:"required,gt=0"`
	CategoryId        uint64  `json:"categoryId"         binding:"required,gt=0"`
	Title             string  `json:"title"              binding:"required"`
	Description       string  `json:"description"`
	Price             float64 `json:"price"              binding:"required,gt=0"`
	Amount            uint    `json:"amount"             binding:"required,gt=0"`
	MainImage         string  `json:"mainImage"          binding:"required"`
	Delivery          string  `json:"delivery"           binding:"required"`
	Assurance         string  `json:"assurance"          binding:"required"`
	Name              string  `json:"name"               binding:"required"`
	Weight            float32 `json:"weight"             binding:"required"`
	Brand             string  `json:"brand"`
	Origin            string  `json:"origin"`
	ShelfLife         uint    `json:"shelfLife"          binding:"required,gt=0"`
	NetWeight         float32 `json:"netWeight"          binding:"required,gt=0"`
	UseWay            string  `json:"useWay"`
	PackingWay        string  `json:"packingWay"`
	StorageConditions string  `json:"storageConditions"`
	DetailImage       string  `json:"detailImage"        binding:"required"`
	Status            uint    `json:"status"             binding:"omitempty,oneof=1 2"`
}

// Seller 商品状态更新参数模型
type SellerProductStatusUpdateParam struct {
	Id     uint64 `json:"id"     binding:"required,gt=0"`
	Status uint   `json:"status" binding:"required,oneof=1 2"`
}

// Seller 商品信息查询参数模型
type SellerProductInfoParam struct {
	Id uint64 `form:"id" binding:"required,gt=0"`
}

// Seller 商品列表查询参数模型
type SellerProductListParam struct {
	Page       Page
	Id         uint64 `form:"id"         binding:"omitempty,gt=0"`
	CategoryId uint64 `form:"categoryId" binding:"omitempty,gt=0"`
	Title      string `form:"title"      binding:"omitempty"`
	Status     uint   `form:"status"     binding:"omitempty,oneof=1 2"`
}

// Seller 商品信息传输模型
type SellerProductInfo struct {
	Id                uint64  `json:"id"`
	CategoryId        uint64  `json:"categoryId"`
	Title             string  `json:"title"`
	Description       string  `json:"description"`
	Price             float64 `json:"price"`
	Amount            uint    `json:"amount"`
	MainImage         string  `json:"mainImage"`
	Delivery          string  `json:"delivery"`
	Assurance         string  `json:"assurance"`
	Name              string  `json:"name"`
	Weight            float32 `json:"weight"`
	Brand             string  `json:"brand"`
	Origin            string  `json:"origin"`
	ShelfLife         uint    `json:"shelfLife"`
	NetWeight         float32 `json:"netWeight"`
	UseWay            string  `json:"useWay"`
	PackingWay        string  `json:"packingWay"`
	StorageConditions string  `json:"storageConditions"`
	DetailImage       string  `json:"detailImage"`
}

// Seller 商品项传输模型
type SellerProductItem struct {
	Id        uint64  `json:"id"`
	Title     string  `json:"title"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	MainImage string  `json:"mainImage"`
}

// Seller 商品列表传输模型
type SellerProductList struct {
	Id        uint64  `json:"id"`
	Title     string  `json:"title"`
	Price     float64 `json:"price"`
	Amount    int     `json:"amount"`
	MainImage string  `json:"mainImage"`
	Sales     int     `json:"sales"`
	Status    int     `json:"status"`
	Created   string  `json:"created"`
}

// Buyer 商品列表参数模型
type BuyerProductListParam struct {
	CategoryId uint64 `form:"categoryId"`
}

// Buyer 商品搜索参数模型
type BuyerProductSearchParam struct {
	KeyWord string `form:"keyWord"`
}

// Buyer 商品详情参数模型
type BuyerProductDetailParam struct {
	ProductId uint64 `form:"productId"`
}

// Buyer 商品列表传输模型
type BuyerProductList struct {
	Id          uint64  `json:"id"`
	MainImage   string  `json:"mainImage"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Sales       int     `json:"sales"`
}

// Buyer 商品项传输模型
type BuyerProductItem struct {
	Id        uint64  `json:"id"`
	Title     string  `json:"title"`
	Price     float64 `json:"price"`
	MainImage string  `json:"mainImage"`
}

// Buyer 商品信息传输模型
type BuyerProductDetail struct {
	Id                uint64  `json:"id"`
	Title             string  `json:"title"`
	Price             float64 `json:"price"`
	Amount            uint    `json:"amount"`
	Sales             uint    `json:"sales"`
	MainImage         string  `json:"mainImage"`
	Delivery          string  `json:"delivery"`
	Assurance         string  `json:"assurance"`
	Name              string  `json:"name"`
	Brand             string  `json:"brand"`
	Origin            string  `json:"origin"`
	ShelfLife         uint    `json:"shelfLife"`
	NetWeight         float32 `json:"netWeight"`
	UseWay            string  `json:"useWay"`
	PackingWay        string  `json:"packingWay"`
	StorageConditions string  `json:"storageConditions"`
	DetailImage       string  `json:"detailImage"`
}
