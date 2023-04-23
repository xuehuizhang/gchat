package server

import (
	"com.gchat/connect/api"
	"com.gchat/connect/conn"
	"com.gchat/connect/handler"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type ChatServer struct {
	wsServer   *http.Server
	WsHandler  handler.WebsocketHandler
	httpServer *http.Server
}

func DefaultServer() *ChatServer {
	return &ChatServer{
		WsHandler: handler.WebsocketHandler{},
	}
}

func (c *ChatServer) RunWsServer() {

	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":12345",
		Handler: mux,
	}
	c.wsServer = server

	mux.HandleFunc("/chat", c.WsHandler.Handler)

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("ws server 启动失败")
		return
	}
}

func (c *ChatServer) RunTcpServer() {
	//todo
}

func (c *ChatServer) RunHttpServer() {
	r := gin.Default()

	api.InitRouter(r)

	server := &http.Server{
		Addr:    ":9091",
		Handler: r,
	}
	c.httpServer = server
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("服务暂停")
	}
}

func (c *ChatServer) RunAll() {
	//todo
}

func (c *ChatServer) Close(finishChan chan struct{}) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*10)
	defer cancelFunc()

	//关闭websocket服务
	if c.wsServer != nil {
		c.wsServer.Shutdown(ctx)
	}

	//关闭httpServer
	if c.httpServer != nil {
		c.httpServer.Shutdown(ctx)
	}

	//关闭所有链接
	conn.RemoveAll()

	//关闭信号
	close(finishChan)
}
