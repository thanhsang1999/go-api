package subscriber

import (
	"context"
	"go-api/component/appctx"
	restaurantstorage "go-api/module/restaurant/storage"
	"go-api/pubsub"
)

//func IncreaseLikeCountAfterUserLikeRestaurant(appCtx appctx.AppContext, ctx context.Context) {
//	c, _ := appCtx.GetPubSub().Subscribe(ctx, common.TopicLikeRestaurant)
//	store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
//	go func() {
//		defer common.AppRecover()
//		for {
//			msg := <-c
//			likeData := msg.Data().(HasRestaurantId)
//			_ = store.UpdateIncreaseLike(ctx, likeData.GetRestaurantId())
//		}
//	}()
//}

func IncreaseLikeCountAfterUserLikeRestaurant(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "IncreaseLikeCountAfterUserLikeRestaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
			likeData := message.Data().(HasRestaurantId)
			return store.UpdateIncreaseLike(ctx, likeData.GetRestaurantId())
		}}
}
