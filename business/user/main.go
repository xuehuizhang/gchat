package main

import (
	"com.gchat.business/user/initialize"
	"com.gchat.business/user/router"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var exitChan = make(chan struct{})

func main() {
	initialize.Init()

	r := gin.Default()

	//注册路由
	router.Init(r)

	server := &http.Server{
		Addr:              ":9099",
		Handler:           r,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
	}

	//优雅退出
	go gentleExit(server)

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("服务", err)
	}
	<-exitChan
}

func gentleExit(server *http.Server) {
	c := make(chan os.Signal, 1)

	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	go func() {
		sig := <-c
		switch sig {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			ServerExit(server)
		case syscall.SIGUSR1:
			fmt.Println("user1", sig)
		case syscall.SIGUSR2:
			fmt.Println("user2", sig)
		default:
			fmt.Println("other", sig)
		}
	}()
}

func ServerExit(server *http.Server) {
	fmt.Println("开始退出....", time.Now())
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*10)
	defer cancelFunc()

	err := server.Shutdown(ctx)
	if err != nil {
		if err.Error() == "context deadline exceeded" {
			fmt.Println("超时退出")
			close(exitChan)
		}
		return
	}
	fmt.Println("退出成功....", time.Now())
	close(exitChan)
}
