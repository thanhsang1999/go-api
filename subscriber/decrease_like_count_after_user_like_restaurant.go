package subscriber

import (
	"context"
	"go-api/component/appctx"
	restaurantstorage "go-api/module/restaurant/storage"
	"go-api/pubsub"
)

type HasRestaurantId interface {
	GetRestaurantId() int
}

//	func DecreaseLikeCountAfterUserLikeRestaurant(appCtx appctx.AppContext, ctx context.Context) {
//		c, _ := appCtx.GetPubSub().Subscribe(ctx, common.TopicDisLikeRestaurant)
//		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
//		go func() {
//			defer common.AppRecover()
//			for {
//				msg := <-c
//				likeData := msg.Data().(HasRestaurantId)
//				_ = store.UpdateDecreaseLike(ctx, likeData.GetRestaurantId())
//			}
//		}()
//	}
func DecreaseLikeCountAfterUserLikeRestaurant(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "DecreaseLikeCountAfterUserLikeRestaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
			likeData := message.Data().(HasRestaurantId)
			return store.UpdateDecreaseLike(ctx, likeData.GetRestaurantId())
		}}
}
