package main

import (
	"github.com/gin-gonic/gin"
	"go-api/component/appctx"
	"go-api/middleware"
	"go-api/module/restaurant/transport/ginrestaurant"
	ginrestaurantlike "go-api/module/restaurantlike/transport/gin"
)

func RouterV1(appCtx appctx.AppContext, v1 *gin.RouterGroup) {
	restaurantRouter := v1.Group("/restaurants", middleware.Authenticate(appCtx))
	restaurantRouter.POST("", ginrestaurant.CreateRestaurant(appCtx))
	restaurantRouter.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))
	restaurantRouter.GET("", ginrestaurant.ListRestaurant(appCtx))
	restaurantRouter.POST("/:id/like", ginrestaurantlike.LikeRestaurant(appCtx))
	restaurantRouter.DELETE("/:id/like", ginrestaurantlike.DislikeRestaurant(appCtx))
}
