package main

import (
	"github.com/gin-gonic/gin"
	"go-api/component/appctx"
	"go-api/module/restaurant/transport/ginrestaurant"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	dsn := os.Getenv("SQL_DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug()
	appCtx := appctx.NewAppContext(db)

	router := gin.Default()
	v1 := router.Group("/v1")
	restaurantRouter := v1.Group("/restaurants")
	restaurantRouter.POST("", ginrestaurant.CreateRestaurant(appCtx))
	restaurantRouter.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))
	restaurantRouter.GET("", ginrestaurant.ListRestaurant(appCtx))
	router.Run() //

}
