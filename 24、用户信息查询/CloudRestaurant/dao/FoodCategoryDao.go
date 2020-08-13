package dao

import (
	"CloudRestaurant/tool"
	"CloudRestaurant/model"
)

type FoodCategoryDao struct {
	*tool.Orm
}

//实例化Dao对象
func NewFoodCategoryDao() *FoodCategoryDao {
	return &FoodCategoryDao{tool.DbEngine}
}

//从数据库中查询所有的食品种类，并返回
func (fcd *FoodCategoryDao) QueryCategories() ([]model.FoodCategory, error) {
	var categories []model.FoodCategory
	if err := fcd.Engine.Find(&categories); err != nil {
		return nil, err
	}
	return categories, nil
}
