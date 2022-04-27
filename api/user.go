package api

import (
	"github.com/gin-gonic/gin"
	"mall/common"
	"mall/models"
	"mall/response"
	"mall/service"
)

var seller service.SellerService
var buyer service.BuyerService

// GetCaptcha 获取验证码
func GetCaptcha(c *gin.Context) {
	id, b64s, _ := common.GenerateCaptcha()
	data := map[string]interface{}{"captchaId": id, "captchaImg": b64s}
	response.Success("操作成功", data, c)
}

// SellerLogin 用户登录
func SellerLogin(c *gin.Context) {
	var param models.SellerLoginParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}

	// 检查验证码
	if !common.VerifyCaptcha(param.CaptchaId, param.CaptchaValue) {
		response.Failed("验证码错误", c)
		return
	}
	// 生成token
	uid := seller.SellerLogin(param)
	if uid > 0 {
		token, _ := common.GenerateToken(param.Username)
		userInfo := models.UserInfo{
			Uid:   uid,
			Token: token,
		}
		response.Success("登录成功", userInfo, c)
		return
	}
	response.Failed("用户名或密码错误", c)
}

// BuyerLogin 用户登录
func BuyerLogin(c *gin.Context) {
	var param models.BuyerLoginParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}

	// 生成token
	uid := buyer.BuyerLogin(param)
	if uid > 0 {
		token, _ := common.GenerateToken(param.Username)
		userInfo := models.UserInfo{
			Uid:   uid,
			Token: token,
		}
		response.Success("登录成功", userInfo, c)
		return
	}
	response.Failed("用户名或密码错误", c)
}
