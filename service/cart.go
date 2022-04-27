package service

import (
	"mall/global"
	"mall/models"
	"strconv"
	"strings"
)

type CartService struct {
}

// Add Buyer 添加购物车
func (*CartService) Add(param models.BuyerCartAddParam) bool {
	key := strings.Join([]string{"user", param.UserId, "cart"}, ":")
	pid := strconv.Itoa(int(param.ProductId))
	return global.Rdb.HSetNX(ctx, key, pid, param.ProductCount).Val()
}

// Delete Buyer 删除购物车中的某项商品
func (*CartService) Delete(param models.BuyerCartDeleteParam) int64 {
	key := strings.Join([]string{"user", param.UserId, "cart"}, ":")
	return global.Rdb.HDel(ctx, key, param.ProductId).Val()
}

// Clear Buyer 清楚购物车中的商品
func (*CartService) Clear(param models.BuyerCartClearParam) int64 {
	key := strings.Join([]string{"user", param.UserId, "cart"}, ":")
	pidsAndCounts := global.Rdb.HGetAll(ctx, key).Val()
	var rows int64
	for id, _ := range pidsAndCounts {
		rows += global.Rdb.HDel(ctx, key, id).Val()
	}
	return rows
}

// GetInfo Buyer 获取购物车信息
func (*CartService) GetInfo(param models.BuyerCartQueryParam) models.BuyerCartInfo {
	var cartInfo models.BuyerCartInfo
	key := strings.Join([]string{"user", param.UserId, "cart"}, ":")
	productIdsAndCounts := global.Rdb.HGetAll(ctx, key).Val()
	productIds := make([]uint, 0)
	idsAndCounts := make(map[uint64]int, 0)
	for pid, pcount := range productIdsAndCounts {
		id, _ := strconv.Atoi(pid)
		count, _ := strconv.Atoi(pcount)
		productIds = append(productIds, uint(id))
		idsAndCounts[uint64(id)] = count
	}
	if len(productIds) > 0 {
		global.Db.Table("product").Find(&cartInfo.CartItem, productIds)
		for index, item := range cartInfo.CartItem {
			cartInfo.CartItem[index].Count = idsAndCounts[item.Id]
			cartInfo.TotalPrice = cartInfo.TotalPrice + item.Price*float64(idsAndCounts[item.Id])
		}
		return cartInfo
	}
	return cartInfo
}
