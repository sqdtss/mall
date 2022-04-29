package models

// Category 数据库,类目数据映射模型
type Category struct {
	Id       uint64 `gorm:"primaryKey"`
	Name     string `gorm:"name"`
	ParentId uint64 `gorm:"parent_id"`
	Level    uint   `gorm:"level"`
	Sort     uint   `gorm:"sort"`
	Created  string `gorm:"created"`
	Updated  string `gorm:"updated"`
}

// SellerCategoryCreateParam Seller 类目创建参数模型
type SellerCategoryCreateParam struct {
	Name     string `json:"name"     binding:"required"`
	ParentId uint64 `json:"parentId" binding:"required,gt=0"`
	Level    uint   `json:"level"    binding:"required,oneof=1 2 3"`
	Sort     uint   `json:"sort"     binding:"required,gt=0"`
}

// SellerCategoryDeleteParam Seller 类目删除参数模型
type SellerCategoryDeleteParam struct {
	Id uint64 `form:"id" binding:"required,gt=0"`
}

// SellerCategoryUpdateParam Seller 类目更新参数模型
type SellerCategoryUpdateParam struct {
	Id   uint64 `json:"id"       binding:"required,gt=0"`
	Name string `json:"name"     binding:"required"`
	Sort uint   `json:"sort"     binding:"required,gt=0"`
}

// SellerCategoryQueryParam Seller 类目查询参数模型
type SellerCategoryQueryParam struct {
	Page     Page
	Id       uint64 `form:"id"       binding:"omitempty,gt=0"`
	Name     string `form:"name"     binding:"omitempty"`
	ParentId uint64 `form:"parentId" binding:"omitempty,gt=0"`
	Level    uint   `form:"level"    binding:"omitempty,oneof=1 2 3"`
}

// SellerCategoryList Seller 类目列表传输模型
type SellerCategoryList struct {
	Id       uint64 `json:"id"`
	Name     string `json:"name"`
	ParentId uint64 `json:"parentId"`
	Level    uint   `json:"level"`
	Sort     uint   `json:"sort"`
}

// SellerCategoryOption Seller 类目选项传输模型
type SellerCategoryOption struct {
	Value    uint64                 `json:"value"`
	Label    string                 `json:"label"`
	Children []SellerCategoryOption `json:"children"`
}

// BuyerCategoryOption Buyer 类目选项传输模型
type BuyerCategoryOption struct {
	Id   uint64 `json:"id"`
	Text string `json:"text"`
}
