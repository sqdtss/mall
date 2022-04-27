package service

import (
	"mall/global"
	"mall/models"
	"strconv"
	"strings"
)

type CollectionService struct {
}

// Add Buyer 添加收藏的商品
func (*CollectionService) Add(param models.BuyerCollectionAddParam) int64 {
	key := strings.Join([]string{"user", param.UserId, "collect"}, ":")
	return global.Rdb.SAdd(ctx, key, param.ProductId).Val()
}

// Delete Buyer 删除收藏的商品
func (*CollectionService) Delete(param models.BuyerCollectionDeleteParam) int64 {
	key := strings.Join([]string{"user", param.UserId, "collect"}, ":")
	pids := global.Rdb.SMembers(ctx, key).Val()
	return global.Rdb.SRem(ctx, key, pids).Val()
}

// GetList Buyer 获取收藏的商品列表
func (*CollectionService) GetList(param models.BuyerCollectionQueryParam) []models.BuyerCollectionList {
	var items []models.BuyerCollectionList
	key := strings.Join([]string{"user", param.UserId, "collect"}, ":")
	pids := global.Rdb.SMembers(ctx, key).Val()
	var productIds []uint
	for _, pid := range pids {
		id, _ := strconv.Atoi(pid)
		productIds = append(productIds, uint(id))
	}
	if len(productIds) > 0 {
		global.Db.Table("product").Find(&items, productIds)
		return items
	}
	return items
}
