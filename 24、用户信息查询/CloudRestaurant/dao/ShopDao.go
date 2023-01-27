package dao

import (
	"CloudRestaurant/model"
	"CloudRestaurant/tool"
)

type ShopDao struct {
	*tool.Orm
}

func NewShopDao() *ShopDao {
	return &ShopDao{tool.DbEngine}
}

const DEFAULT_RANGE = 5

// 116.31   39.21
// 115.31 - 117.31   38.21- 40.21

/**
 * 操作数据库查询商铺数据列表
 */
func (shopDao *ShopDao) QueryShops(longitude, latitude float64, keyword string) []model.Shop {
	var shops []model.Shop

	if keyword == "" {
		err := shopDao.Engine.Where(" longitude > ? and longitude < ? and latitude > ? and latitude < ?  and status = 1 ", longitude-DEFAULT_RANGE, longitude+DEFAULT_RANGE, latitude-DEFAULT_RANGE, latitude+DEFAULT_RANGE).Find(&shops)
		if err != nil {
			return nil
		}
	} else {
		err := shopDao.Engine.Where(" longitude > ? and longitude < ? and latitude > ? and latitude < ? and name like ？ and status = 1", longitude-DEFAULT_RANGE, longitude+DEFAULT_RANGE, latitude-DEFAULT_RANGE, latitude+DEFAULT_RANGE, keyword).Find(&shops)
		if err != nil {
			return nil
		}
	}
	return shops
}
