package subscriber

import (
	"context"
	"go-api/component/appctx"
	"go-api/pubsub"
	"log"
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

func SendMessageAfterUserLikeRestaurant(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "SendMessageAfterUserLikeRestaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			log.Println("Sending message to restaurant")
			return nil
		}}
}
