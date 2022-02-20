package routers

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())   // 输出请求日志，并标准化日志的格式
	r.Use(gin.Recovery()) // 异常捕获，针对每次请求处理进行Recovery处理，防止因为出现panic导致服务崩溃，同时将异常日志的格式标准化

	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/tags")
		apiv1.DELETE("/tags/:id")
		apiv1.PUT("/tags/:id")
		apiv1.PATCH("/tags/:id/state")
		apiv1.GET("/tags")

		apiv1.POST("/articles")
		apiv1.DELETE("/articles/:id")
		apiv1.PUT("/articles/:id")
		apiv1.PATCH("/articles/:id/state")
		apiv1.GET("/articles/:id")
		apiv1.GET("/articles")
	}

	return r
}
