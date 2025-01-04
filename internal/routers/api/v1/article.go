package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/testercc/blog-service/pkg/app"
	"github.com/testercc/blog-service/pkg/errcode"
)

type Article struct {

}

func NewArticle() Article {
	return Article{}
}

func (a Article) Get(c *gin.Context) {
	// todo：正确逻辑不是这样，这里只是测试
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)  // debug: curl -v http://127.0.0.1:8000/api/v1/articles/1  {"code":10000000,"msg":"服务内部错误"}
	return
}
func (a Article) List(c *gin.Context) {}
func (a Article) Create(c *gin.Context) {}
func (a Article) Update(c *gin.Context) {}
func (a Article) Delete(c *gin.Context) {}