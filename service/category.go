package service

import (
	"mall/common"
	"mall/global"
	"mall/models"
)

type SellerCategoryService struct {
}

type BuyerCategoryService struct {
}

// Create Seller 创建类目
func (*SellerCategoryService) Create(param models.SellerCategoryCreateParam) uint64 {
	var category models.Category
	result := global.Db.Where("name = ?", param.Name).First(&category)
	if result.RowsAffected > 0 {
		return category.Id
	}
	category = models.Category{
		Name:     param.Name,
		ParentId: param.ParentId,
		Level:    param.Level,
		Sort:     param.Sort,
		Created:  common.NowTime(),
	}
	global.Db.Create(&category)
	return category.Id
}

// Delete Seller 删除类目
func (*SellerCategoryService) Delete(param models.SellerCategoryDeleteParam) int64 {
	var pid2, pid3 models.Category
	global.Db.Where("parent_id = ?", param.Id).First(&pid2)
	global.Db.Where("parent_id = ?", pid2.Id).First(&pid3)
	return global.Db.Delete(&models.Category{}, []uint64{param.Id, pid2.Id, pid3.Id}).RowsAffected
}

// Update Seller 更新类目
func (*SellerCategoryService) Update(param models.SellerCategoryUpdateParam) int64 {
	category := models.Category{
		Id:      param.Id,
		Name:    param.Name,
		Sort:    param.Sort,
		Updated: common.NowTime(),
	}
	return global.Db.Model(&category).Updates(category).RowsAffected
}

// GetList Seller 获取类目列表
func (*SellerCategoryService) GetList(param models.SellerCategoryQueryParam) ([]models.SellerCategoryList, int64) {
	var categoryList []models.SellerCategoryList
	query := &models.Category{
		Id:       param.Id,
		Name:     param.Name,
		Level:    param.Level,
		ParentId: param.ParentId,
	}
	rows := common.RestPage(param.Page, "category", query, &categoryList, &[]models.Category{})
	return categoryList, rows
}

// GetOption Seller 获取类目选项
func (*SellerCategoryService) GetOption() (option []models.SellerCategoryOption) {
	var selectList []models.SellerCategoryList
	global.Db.Table("category").Find(&selectList)
	return getTreeOptions(1, selectList)
}

// GetOption Buyer 获取类目选项
func (*BuyerCategoryService) GetOption() []models.BuyerCategoryOption {
	var category []models.Category
	var optionList []models.BuyerCategoryOption
	global.Db.Table("category").Where("parent_id = ?", 1).Find(&category)
	for _, c := range category {
		optionList = append(optionList, models.BuyerCategoryOption{Id: c.Id, Text: c.Name})
	}
	return optionList
}

// Seller 获取树形结构的选项
func getTreeOptions(id uint64, cateList []models.SellerCategoryList) (option []models.SellerCategoryOption) {
	var optionList []models.SellerCategoryOption
	for _, opt := range cateList {
		if opt.ParentId == id && (opt.Level == 1 || opt.Level == 2) {
			option := models.SellerCategoryOption{
				Value:    opt.Id,
				Label:    opt.Name,
				Children: getTreeOptions(opt.Id, cateList),
			}
			if opt.Level == 2 {
				option.Children = nil
			}
			optionList = append(optionList, option)
		}
	}
	return optionList
}
