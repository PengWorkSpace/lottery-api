package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"lottery-api/internal/server"
	"lottery-api/internal/service"
)

func main() {

	//启动服务
	svc := service.New()
	//开启8080端口
	srv := server.New(":8080", svc)
	//接受信号量，退出程序
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	defer svc.Close()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("ottery server shutdown:", err)
	}

	log.Println("lottery server exiting")
}
