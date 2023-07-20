package restaurantlikebiz

import (
	"context"
	restaurantlikemodel "go-api/module/restaurantlike/model"
)

type RestaurantDisLikeStore interface {
	Delete(context context.Context, condition map[string]interface{}) error
}

type createRestaurantDisLikeBiz struct {
	store RestaurantDisLikeStore
}

func NewCreateRestaurantDisLikeBiz(store RestaurantDisLikeStore) *createRestaurantDisLikeBiz {
	return &createRestaurantDisLikeBiz{store}
}

func (c *createRestaurantDisLikeBiz) DislikeRestaurant(ctx context.Context, data *restaurantlikemodel.RestaurantLikes) error {
	err := c.store.Delete(ctx, map[string]interface{}{
		"restaurant_id": data.RestaurantId,
		"user_id":       data.UserId,
	})
	if err != nil {
		return err
	}
	return nil
}
