package service

import (
	"CloudRestaurant/model"
	"CloudRestaurant/dao"
)

type GoodService struct {
}

func NewGoodService() *GoodService {
	return &GoodService{}
}

/**
 * 获取商家的食品列表
 */
func (gs *GoodService) GetFoods(shop_id int64) []model.Goods {
	goodDao := dao.NewGoodDao()
	goods, err := goodDao.QueryFoods(shop_id)
	if err != nil {
		return nil
	}
	return goods
}
