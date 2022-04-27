package models

// 数据库，Seller 数据映射模型
type Seller struct {
	Id           uint64 `gorm:"primaryKey"`
	Username     string `gorm:"username"`
	Password     string `gorm:"password"`
	Status       uint   `gorm:"status"`
	CaptchaId    string `gorm:"captchaId"`
	CaptchaValue string `gorm:"captchaValue"`
}

// Seller 登录参数模型
type SellerLoginParam struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	CaptchaId    string `json:"captchaId"`
	CaptchaValue string `json:"captchaValue"`
}

// 数据库，Seller 数据映射模型
type Buyer struct {
	Id           uint64 `gorm:"primaryKey"`
	Username     string `gorm:"username"`
	Password     string `gorm:"password"`
	Status       uint   `gorm:"status"`
	CaptchaId    string `gorm:"captchaId"`
	CaptchaValue string `gorm:"captchaValue"`
}

// Buyer 登录参数模型
type BuyerLoginParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 用户信息传输模型
type UserInfo struct {
	Uid   uint64 `json:"uid"`
	Token string `json:"token"`
}
