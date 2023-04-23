package main

import (
	"com.gchat/connect/initialize"
	"com.gchat/connect/server"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var finishChan = make(chan struct{})

func main() {
	//初始化
	initialize.Init()

	//启动websocket服务
	s := server.DefaultServer()

	go s.RunWsServer()

	go s.RunHttpServer()

	//优雅退出
	exitChan := make(chan os.Signal)
	signal.Notify(exitChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)

	sig := <-exitChan
	switch sig {
	case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
		s.Close(finishChan)
	case syscall.SIGUSR1:
		fmt.Println("user1", sig)
	case syscall.SIGUSR2:
		fmt.Println("user2", sig)
	default:
		fmt.Println("other", sig)
	}
}
