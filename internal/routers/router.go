package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	v1 "github.com/testercc/blog-service/internal/routers/api/v1"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())   // 输出请求日志，并标准化日志的格式
	r.Use(gin.Recovery()) // 异常捕获，针对每次请求处理进行Recovery处理，防止因为出现panic导致服务崩溃，同时将异常日志的格式标准化
	// 先同通过 swag init 把 Swagger API 所需要的文件都生成，再在 routers 中进行swagger默认初始化和注册对应的路由，在浏览器中访问  http://127.0.0.1:8000/swagger/index.html#/  swagger可根据需求修改
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 编写好路由的 Handler 方法后，只需要将其注册到对应的路由规则上
	article := v1.NewArticle()
	tag := v1.NewTag()


	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/tags", tag.Create)
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.GET("/tags", tag.List)

		apiv1.POST("/articles",article.Create)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.GET("/articles", article.List)
	}

	return r
}
