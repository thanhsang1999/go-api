package restaurantstorage

import (
	"context"
	"go-api/common"
	restaurantmodel "go-api/module/restaurant/model"
	"gorm.io/gorm"
)

func (s *sqlStore) UpdateIncreaseLike(context context.Context, idRestaurant int) error {
	if err := s.db.Table(restaurantmodel.RestaurantUpdate{}.TableName()).
		Where("id = ?", idRestaurant).
		Update("like_count", gorm.Expr("like_count + ?", 1)).
		Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
func (s *sqlStore) UpdateDecreaseLike(context context.Context, idRestaurant int) error {
	if err := s.db.Table(restaurantmodel.RestaurantUpdate{}.TableName()).
		Where("id = ?", idRestaurant).
		Update("like_count", gorm.Expr("like_count - ?", 1)).
		Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
