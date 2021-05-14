package service

import (
	"CloudRestaurant/dao"
	"CloudRestaurant/model"
)

type FoodCategoryService struct {
}

/**
 * 获取美食类别
 */
func (fcs *FoodCategoryService) Categories() ([]model.FoodCategory, error) {
	//数据库操作层
	foodCategoryDao := dao.NewFoodCategoryDao()
	return foodCategoryDao.QueryCategories()
}
