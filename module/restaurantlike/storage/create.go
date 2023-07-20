package restaurantlikestorage

import (
	"context"
	restaurantlikemodel "go-api/module/restaurantlike/model"
)

func (s *sqlStore) Create(context context.Context, data *restaurantlikemodel.RestaurantLikes) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
