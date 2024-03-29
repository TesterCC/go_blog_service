package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/testercc/blog-service/docs"   // 注意import，否则初始化报错
	"github.com/testercc/blog-service/global"
	"github.com/testercc/blog-service/internal/model"
	"github.com/testercc/blog-service/internal/routers"
	"github.com/testercc/blog-service/pkg/logger"
	"github.com/testercc/blog-service/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
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

/*
在 Go 语言中，init 方法常用于应用程序内的一些初始化操作，
它在 main 方法之前自动执行，它的执行顺序是：全局变量初始化 =》init 方法 =》main 方法，
但并不是建议滥用，因为如果 init 过多，你可能会迷失在各个库的 init 方法中，会非常麻烦。

该 init 方法主要作用是进行应用程序的初始化流程控制，整个应用代码里也只会有一个 init 方法，
因此我们在这里调用了初始化配置的方法，达到配置文件内容映射到应用配置结构体的作用。
*/

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
}

// 在 init 方法中新增了日志组件的流程，并在 setupLogger 方法内部对 global 的包全局变量 Logger 进行了初始化，
// 需要注意的是我们使用了 lumberjack 作为日志库的 io.Writer，并且设置日志文件所允许的最大占用空间为 600MB、日志文件最大生存周期为 10 天，并且设置日志文件名的时间格式为本地时间。
func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   500,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}

func setupDBEngine() error {
	var err error
	//注意不要用 := , 因为 := 会重新声明并创建了左侧的新局部变量，因此在其它包中调用 global.DBEngine 变量时，它仍然是 nil，
	//仍然是达不到可用标准，因为根本就没有赋值到真正需要赋值的包全局变量 global.DBEngine 上。
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}

	return nil
}

func setupSetting() error {
	setting, err := setting.NewSetting()

	if err != nil {
		return err
	}

	err = setting.ReadSection("Server", &global.ServerSetting)

	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	return nil
}

// v2: main()
//func main() {
//
//	router := routers.NewRouter()
//	// 通过自定义 http.Server，设置了监听的 TCP Endpoint、处理的程序、允许读取/写入的最大时间、请求头的最大字节数等基础参数
//	s := &http.Server{
//		Addr:           ":8080",
//		Handler:        router,
//		ReadTimeout:    10 * time.Second,
//		WriteTimeout:   10 * time.Second,
//		MaxHeaderBytes: 1 << 20,
//	}
//	// 开始监听
//	s.ListenAndServe()
//}

// v3 启动配置通过配置获取
// 针对项目写注解

// @title          Security Development Blog System
// @version        1.0
// @description    TesterCC - Security Development Blog API Docs.
// @termsOfService https://github.com/TesterCC/

// @contact.name  API Support
// @contact.url   https://github.com/swaggo/swag/blob/master/README_zh-CN.md
// @contact.email testerlyx@foxmail.com

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html

// @host     localhost:8888
// @BasePath /api/v1
func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	// 通过自定义 http.Server，设置了监听的 TCP Endpoint、处理的程序、允许读取/写入的最大时间、请求头的最大字节数等基础参数
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	// for debug
	fmt.Println("[D] Launch server at 127.0.0.1:" + global.ServerSetting.HttpPort)
	// for debug
	//fmt.Println(global.ServerSetting)
	//fmt.Println(global.AppSetting)
	//fmt.Println(global.DatabaseSetting)
	// global.Logger.Infof(nil,"%s: go-programming-tour-book/%s", "TesterCC","blog-service")  // debug test logger, log in /storage/file
	//log.Fatal("Server Test Log")  // debug
	// 开始监听
	s.ListenAndServe()

}

// test: curl http://127.0.0.1:8080/api/v1/tags
// Swagger 的地址 http://127.0.0.1:8000/swagger/index.html