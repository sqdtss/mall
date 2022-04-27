package service

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"mall/common"
	"mall/global"
	"mall/models"
	"reflect"
	"strconv"
)

type SellerProductService struct {
}

type BuyerProductService struct {
}

// Create Seller 创建商品
func (*SellerProductService) Create(param models.SellerProductCreateParam) int64 {
	product := models.Product{
		CategoryId:        param.CategoryId,
		Title:             param.Title,
		Description:       param.Description,
		Price:             param.Price,
		Amount:            param.Amount,
		MainImage:         param.MainImage,
		Delivery:          param.Delivery,
		Assurance:         param.Assurance,
		Name:              param.Name,
		Weight:            param.Weight,
		Brand:             param.Brand,
		Origin:            param.Origin,
		ShelfLife:         param.ShelfLife,
		NetWeight:         param.NetWeight,
		UseWay:            param.UseWay,
		PackingWay:        param.PackingWay,
		StorageConditions: param.StorageConditions,
		DetailImage:       param.DetailImage,
		Status:            param.Status,
		Created:           common.NowTime(),
	}
	rows := global.Db.Create(&product).RowsAffected
	records := global.Db.First(&product, product.Id).RowsAffected
	if records > 0 {
		id := strconv.FormatUint(product.Id, 10)
		result, err := global.Es.Index().Index("product").Id(id).BodyJson(product).Do(ctx)
		if err != nil {
			fmt.Println(err)
		}
		return result.PrimaryTerm
	}
	return rows
}

// Delete Seller 删除商品
func (*SellerProductService) Delete(param models.SellerProductDeleteParam) int64 {
	rows := global.Db.Delete(&models.Product{}, param.Id).RowsAffected
	if rows > 0 {
		id := strconv.FormatUint(param.Id, 10)
		_, err := global.Es.Delete().Index("product").Id(id).Do(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	return rows
}

// Update Seller 更新商品
func (*SellerProductService) Update(param models.SellerProductUpdateParam) int64 {
	product := models.Product{
		Id:                param.Id,
		CategoryId:        param.CategoryId,
		Title:             param.Title,
		Description:       param.Description,
		Price:             param.Price,
		Amount:            param.Amount,
		MainImage:         param.MainImage,
		Delivery:          param.Delivery,
		Assurance:         param.Assurance,
		Name:              param.Name,
		Weight:            param.Weight,
		Brand:             param.Brand,
		Origin:            param.Origin,
		ShelfLife:         param.ShelfLife,
		NetWeight:         param.NetWeight,
		UseWay:            param.UseWay,
		PackingWay:        param.PackingWay,
		StorageConditions: param.StorageConditions,
		DetailImage:       param.DetailImage,
		Status:            param.Status,
		Updated:           common.NowTime(),
	}
	rows := global.Db.Model(&product).Updates(product).RowsAffected
	records := global.Db.First(&product, product.Id).RowsAffected
	if records > 0 {
		id := strconv.FormatUint(param.Id, 10)
		_, err := global.Es.Update().Index("product").Id(id).Doc(product).Do(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	return rows
}

// UpdateStatus Seller 更新商品状态
func (*SellerProductService) UpdateStatus(param models.SellerProductStatusUpdateParam) int64 {
	product := models.Product{
		Id:     param.Id,
		Status: param.Status,
	}
	rows := global.Db.Model(&product).Update("status", product.Status).RowsAffected
	records := global.Db.First(&product, product.Id).RowsAffected
	if records > 0 {
		id := strconv.FormatUint(param.Id, 10)
		_, err := global.Es.Update().Index("product").Id(id).Doc(product).Do(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	return rows
}

// GetInfo Seller 获取商品信息
func (*SellerProductService) GetInfo(param models.SellerProductInfoParam) models.SellerProductInfo {
	var product models.SellerProductInfo
	global.Db.Table("product").First(&product, param.Id)
	return product
}

// GetList Seller 获取商品列表
func (*SellerProductService) GetList(param models.SellerProductListParam) ([]models.SellerProductList, int64) {
	query := &models.Product{
		Id:         param.Id,
		CategoryId: param.CategoryId,
		Title:      param.Title,
		Status:     param.Status,
	}
	var productList []models.SellerProductList
	rows := common.RestPage(param.Page, "product", query, &productList, &[]models.Product{})
	return productList, rows
}

// GetList Buyer 获取商品列表
func (*BuyerProductService) GetList(param models.BuyerProductListParam) []models.BuyerProductList {
	var productList []models.BuyerProductList
	if param.CategoryId != 0 {
		category := models.Category{
			ParentId: param.CategoryId,
		}
		var categoryIds []uint64
		global.Db.Table("category").Select("id").Where(category).Find(&categoryIds)
		if len(categoryIds) > 0 {
			global.Db.Table("product").Where("status = 1 and category_id in ?", categoryIds).Find(&productList)
		}
		return productList
	}
	global.Db.Table("product").Where("status = 1").Find(&productList)
	return productList
}

// GetSearchList Buyer 搜索商品
func (*BuyerProductService) GetSearchList(param models.BuyerProductSearchParam) []models.BuyerProductList {
	var productList []models.BuyerProductList
	phraseQuery := elastic.NewMatchPhraseQuery("title", param.KeyWord)
	searchResult, err := global.Es.Search().Index("product").Query(phraseQuery).Do(ctx)
	if err != nil {
		fmt.Println(err)
	}
	for _, item := range searchResult.Each(reflect.TypeOf(models.BuyerProductList{})) {
		t := item.(models.BuyerProductList)
		productList = append(productList, t)
	}
	return productList
}

// GetDetail Buyer 获取商品详情
func (*BuyerProductService) GetDetail(param models.BuyerProductDetailParam) models.BuyerProductDetail {
	var productDetail models.BuyerProductDetail
	global.Db.Table("product").First(&productDetail, param.ProductId)
	return productDetail
}
