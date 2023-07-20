package restaurantlikestorage

import (
	"context"
	restaurantlikemodel "go-api/module/restaurantlike/model"
)

func (s *sqlStore) Delete(context context.Context, condition map[string]interface{}) error {
	if err := s.db.
		Where(condition).
		Delete(&restaurantlikemodel.RestaurantLikes{}).
		Error; err != nil {
		return err
	}
	return nil
}
