package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mall/api"
	"mall/global"
	"mall/middleware"
)

func Router() {

	engine := gin.Default()

	// 开启跨域
	engine.Use(middleware.Cors())

	// 静态资源请求映射
	engine.Static("/image", global.Config.Upload.SavePath)

	// 卖家后端API
	seller := engine.Group("/seller")
	{
		// 用户登录API
		seller.GET("/captcha", api.GetCaptcha)
		seller.POST("/login", api.SellerLogin)
		seller.POST("/register", api.SellerRegister)

		// 开启JWT认证,以下接口需要认证成功才能访问
		seller.Use(middleware.JwtAuth())

		// 文件上传API
		seller.POST("/upload", api.FileUpload)

		// 类目管理API
		seller.POST("/category", api.SellerCreateCategory)
		seller.DELETE("/category", api.SellerDeleteCategory)
		seller.PUT("/category", api.SellerUpdateCategory)
		seller.GET("/category/list", api.SellerGetCategoryList)
		seller.GET("/category/option", api.SellerGetCategoryOption)

		// 商品管理API
		seller.POST("/product", api.SellerCreateProduct)
		seller.DELETE("/product", api.SellerDeleteProduct)
		seller.PUT("/product", api.SellerUpdateProduct)
		seller.PUT("/product/status", api.SellerUpdateProductStatus)
		seller.GET("/product/info", api.SellerGetProductInfo)
		seller.GET("/product/list", api.SellerGetProductList)

		// 订单管理API
		seller.DELETE("/order", api.SellerDeleteOrder)
		seller.PUT("/order", api.SellerUpdateOrder)
		seller.GET("/order/list", api.SellerGetOrderList)
		seller.GET("/order/detail", api.SellerGetOrderDetail)

		// 数据统计API
		seller.GET("/data/overview/info", api.SellerGetDataOverviewInfo)
		seller.GET("/today/order/data/info", api.SellerGetTodayOrderDataInfo)
		seller.GET("/week/data/info", api.SellerGetWeekDataInfo)
	}

	// 买家后端API
	buyer := engine.Group("/buyer")
	{
		// 用户登录API
		buyer.GET("/captcha", api.GetCaptcha)
		buyer.POST("/login", api.BuyerLogin)
		buyer.POST("/register", api.BuyerRegister)

		// 开启JWT认证,以下接口需要认证成功才能访问
		//buyer.Use(middleware.JwtAuth())

		// 分类API
		buyer.GET("/category/option", api.BuyerGetCategoryOption)

		// 商品API
		buyer.GET("/product/list", api.BuyerGetProductList)
		buyer.GET("/product/search", api.BuyerGetProductSearchList)
		buyer.GET("/product/detail", api.BuyerGetProductDetail)

		// 商品订单API
		buyer.POST("/order/create", api.BuyerCreateOrder)
		buyer.GET("/order/list", api.BuyerGetOrderList)

		// 购物车API
		buyer.POST("/cart/add", api.BuyerAddCart)
		buyer.DELETE("/cart/delete", api.BuyerDeleteCart)
		buyer.DELETE("/cart/clear", api.BuyerClearCart)
		buyer.GET("/cart/info", api.BuyerGetCartInfo)

		// 收货地址API
		buyer.POST("/address/add", api.BuyerAddAddress)
		buyer.DELETE("/address/delete", api.BuyerDeleteAddress)
		buyer.PUT("/address/update", api.BuyerUpdateAddress)
		buyer.GET("/address/info", api.BuyerGetAddressUpdateInfo)
		buyer.GET("/address/list", api.BuyerGetAddressList)

		// 商品收藏API
		buyer.POST("/collection/add", api.BuyerAddCollection)
		buyer.DELETE("/collection/delete", api.BuyerDeleteCollection)
		buyer.GET("/collection/list", api.BuyerGetCollectionList)

		// 商品浏览记录API
		buyer.POST("/browse/save", api.BuyerSaveBrowseRecord)
		buyer.DELETE("/browse/delete", api.BuyerDeleteBrowseRecord)
		buyer.GET("/browse/list", api.BuyerGetBrowseRecordList)
	}

	// 启动、监听端口
	post := fmt.Sprintf(":%s", global.Config.Server.Port)
	if err := engine.Run(post); err != nil {
		fmt.Printf("server start error: %s", err)
	}
}
