package restaurantlikebiz

import (
	"context"
	restaurantlikemodel "go-api/module/restaurantlike/model"
)

type RestaurantLikeStore interface {
	Create(context context.Context, data *restaurantlikemodel.RestaurantLikes) error
}

type createRestaurantLikeBiz struct {
	store RestaurantLikeStore
}

func NewCreateRestaurantLikeBiz(store RestaurantLikeStore) *createRestaurantLikeBiz {
	return &createRestaurantLikeBiz{store}
}

func (c *createRestaurantLikeBiz) CreateRestaurantLike(ctx context.Context, data *restaurantlikemodel.RestaurantLikes) error {
	err := c.store.Create(ctx, data)
	if err != nil {
		return err
	}
	return nil
}
