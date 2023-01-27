package service

import (
	"CloudRestaurant/dao"
	"CloudRestaurant/model"
	"strconv"
)

type ShopService struct {
}

// ShopList 查询商铺列表数据
func (shopService *ShopService) ShopList(long, lat string) []model.Shop {
	longitude, err := strconv.ParseFloat(long, 10)
	if err != nil {
		return nil
	}
	latitude, err := strconv.ParseFloat(lat, 10)
	if err != nil {
		return nil
	}

	shopDao := dao.NewShopDao()
	return shopDao.QueryShops(longitude, latitude)
}
