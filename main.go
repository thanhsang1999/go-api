package main

import (
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
	"go-api/common"
	"go-api/component/appctx"
	"go-api/middleware"
	"go-api/pubsub/pubsublocal"
	"go-api/subscriber"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
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
	router.StaticFile("/demo-socket", "./assets/demo.html")
	router.Use(middleware.Recover(appCtx))
	RouterV1(appCtx, router.Group("/v1"))
	socketIO(router, appCtx)
	router.StaticFS("/public", http.Dir("./assets"))
	_ = router.Run(":8000")

}
func socketIO(engine *gin.Engine, appCtx appctx.AppContext) {

	serverSocket := socketio.NewServer(&engineio.Options{Transports: []transport.Transport{websocket.Default}})

	serverSocket.OnConnect("/", func(s socketio.Conn) error {
		log.Println("connected:", s.ID())
		return nil
	})

	serverSocket.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		log.Println("notice:", msg)
		s.Emit("reply", "have "+msg)
	})

	serverSocket.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		return "recv " + msg
	})

	serverSocket.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})

	serverSocket.OnError("/", func(s socketio.Conn, e error) {
		log.Println("meet error:", e)
	})

	serverSocket.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("closed", reason)
	})

	go func() {
		defer common.AppRecover()
		if err := serverSocket.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()

	engine.GET("/socket.io/*any", gin.WrapH(serverSocket))
	engine.POST("/socket.io/*any", gin.WrapH(serverSocket))

}
