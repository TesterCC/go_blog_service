package main

import (
	"github.com/testercc/blog-service/internal/routers"
	"net/http"
	"time"
)

/*
// curl http://127.0.0.1:8080/ping

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	r.Run(":7777")  // default port 8080
}
*/

// 改造为本项目的启动文件

func main() {
	router := routers.NewRouter()
	// 通过自定义 http.Server，设置了监听的 TCP Endpoint、处理的程序、允许读取/写入的最大时间、请求头的最大字节数等基础参数
	s := &http.Server {
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	// 开始监听
	s.ListenAndServe()
}

// test: curl http://127.0.0.1:8080/api/v1/tags