package main

import (
	"github.com/gin-gonic/gin"
	"go-api/component/appctx"
	"go-api/middleware"
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
	router.Use(middleware.Recover(appCtx))
	RouterV1(appCtx, router.Group("/v1"))
	_ = router.Run()

}
