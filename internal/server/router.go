package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"lottery-api/internal/service"
)

var svc *service.Svc

//启动gin服务器，初始化路由
func New(port string, s *service.Svc) *http.Server {
	svc = s
	router := gin.Default()
	initRouter(router)
	srv := &http.Server{Addr: port, Handler: router}
	serveAsync(srv)
	return srv
}

func initRouter(e *gin.Engine) {
	g := e.Group("lottery")
	g.Use(accessJsMiddleware())
	gin.SetMode(gin.DebugMode)
	//发送验证码
	g.POST("/generatecode", generatePhoneCode)
	//用户报名注册
	g.POST("/join", join)
	//抽奖
	g.GET("/draw", rateLimitMiddleware(), drawPrize)
	//获取中奖记录信息
	g.GET("/list", prizeList)
	//导出中奖记录信息
	g.GET("/export", drawRecords)
	//获取用户的征文记录信息
	g.GET("/users/articles", userArticles)
}

func serveAsync(daemon *http.Server) {
	go func() {
		if err := daemon.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
}
