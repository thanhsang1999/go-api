package main

import (
	"github.com/gin-gonic/gin"
	"go-api/component/appctx"
	"go-api/middleware"
	"go-api/pubsub/pubsublocal"
	"go-api/subscriber"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	//SQl_DSN=root:admin@tcp(127.0.0.1:3307)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local
	dsn := os.Getenv("SQL_DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug()
	ps := pubsublocal.NewPubSub()
	appCtx := appctx.NewAppContext(db, ps)
	_ = subscriber.NewEngine(appCtx).Start()
	router := gin.Default()
	router.Use(middleware.Recover(appCtx))
	RouterV1(appCtx, router.Group("/v1"))
	_ = router.Run()

}
