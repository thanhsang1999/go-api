package restaurantlikebiz

import (
	"context"
	"go-api/common"
	restaurantlikemodel "go-api/module/restaurantlike/model"
	"go-api/pubsub"
)

type RestaurantDisLikeStore interface {
	Delete(context context.Context, condition map[string]interface{}) error
}

type createRestaurantDisLikeBiz struct {
	store RestaurantDisLikeStore
	ps    pubsub.Pubsub
}

func NewCreateRestaurantDisLikeBiz(store RestaurantDisLikeStore, ps pubsub.Pubsub) *createRestaurantDisLikeBiz {
	return &createRestaurantDisLikeBiz{store, ps}
}

func (c *createRestaurantDisLikeBiz) DislikeRestaurant(ctx context.Context, data *restaurantlikemodel.RestaurantLikes) error {
	err := c.store.Delete(ctx, map[string]interface{}{
		"restaurant_id": data.RestaurantId,
		"user_id":       data.UserId,
	})
	_ = c.ps.Publish(ctx, common.TopicDisLikeRestaurant, pubsub.NewMessage(data))
	if err != nil {
		return err
	}
	return nil
}
