package restaurantlikebiz

import (
	"context"
	"go-api/common"
	restaurantlikemodel "go-api/module/restaurantlike/model"
	"go-api/pubsub"
)

type RestaurantLikeStore interface {
	Create(context context.Context, data *restaurantlikemodel.RestaurantLikes) error
}

type createRestaurantLikeBiz struct {
	store RestaurantLikeStore
	ps    pubsub.Pubsub
}

func NewCreateRestaurantLikeBiz(store RestaurantLikeStore, ps pubsub.Pubsub) *createRestaurantLikeBiz {
	return &createRestaurantLikeBiz{store, ps}
}

func (c *createRestaurantLikeBiz) CreateRestaurantLike(ctx context.Context, data *restaurantlikemodel.RestaurantLikes) error {
	err := c.store.Create(ctx, data)
	_ = c.ps.Publish(ctx, common.TopicLikeRestaurant, pubsub.NewMessage(data))
	if err != nil {
		return err
	}
	return nil
}
