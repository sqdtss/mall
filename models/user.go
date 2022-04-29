package models

// Seller 数据库，Seller 数据映射模型
type Seller struct {
	Id       uint64 `gorm:"primaryKey"`
	Username string `gorm:"username"`
	Password string `gorm:"password"`
	Status   uint   `gorm:"status"`
	Created  string `gorm:"created"`
	Updated  string `gorm:"updated"`
}

// SellerLoginParam Seller 登录参数模型
type SellerLoginParam struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	CaptchaId    string `json:"captchaId"`
	CaptchaValue string `json:"captchaValue"`
}

// SellerRegisterParam Seller 注册参数模型
type SellerRegisterParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Status   uint   `json:"status"`
}

// Buyer 数据库，Buyer 数据映射模型
type Buyer struct {
	Id       uint64 `gorm:"primaryKey"`
	Username string `gorm:"username"`
	Password string `gorm:"password"`
	Status   uint   `gorm:"status"`
	Created  string `gorm:"created"`
	Updated  string `gorm:"updated"`
}

// BuyerLoginParam Buyer 登录参数模型
type BuyerLoginParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// BuyerRegisterParam Buyer 注册参数模型
type BuyerRegisterParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Status   uint   `json:"status"`
}

// UserInfo 用户信息传输模型
type UserInfo struct {
	Uid   uint64 `json:"uid"`
	Token string `json:"token"`
}
