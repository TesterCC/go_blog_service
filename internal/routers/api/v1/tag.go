package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/testercc/blog-service/pkg/app"
	"github.com/testercc/blog-service/pkg/errcode"
)

type Tag struct {
}

func NewTag() Tag {
	return Tag{}
}

//func (t Tag) Get(c *gin.Context)    {}

//todo: 标签模块的接口注解编写，接下来应当按照注解的含义和参考上述接口注解，完成文章模块接口注解的编写。

// @Summary 获取多个标签
// @Produce json
// @Param   name      query    string           false "标签名称" maxlength(100)
// @Param   state     query    int              false "状态"   Enums(0, 1) default(1)
// @Param   page      query    int              false "页码"
// @Param   page_size query    int              false "每页数量"
// @Success 200       {object} model.TagSwagger "成功"
// @Failure 400       {object} errcode.Error    "请求错误"
// @Failure 500       {object} errcode.Error    "内部错误"
// @Router  /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {
	param := struct {
		Name  string `form:"name" binding:"max=100"`
		State uint8  `form:"state,default=1" binding:"oneof=0 1"`
	}{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		fmt.Printf("app.BindAndValid errs: %v", errs)       // temp resolve
		//global.Logger.Errorf("app.BindAndValid errs: %v", errs)  // fixme: report error
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	response.ToResponse(gin.H{})  // correct
	//response.ToGeneralResponse(200,"success", nil)   // debug ok
	return
}

// @Summary 新增标签
// @Produce json
// @Param   name       body     string           true  "标签名称" minlength(3) maxlength(100)
// @Param   state      body     int              false "状态"   Enums(0, 1)  default(1)
// @Param   created_by body     string           true  "创建者"  minlength(3) maxlength(100)
// @Success 200        {object} model.TagSwagger "成功"
// @Failure 400        {object} errcode.Error    "请求错误"
// @Failure 500        {object} errcode.Error    "内部错误"
// @Router  /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {}

// @Summary 更新标签
// @Produce json
// @Param   id          path     int              true  "标签 ID"
// @Param   name        body     string           false "标签名称" minlength(3) maxlength(100)
// @Param   state       body     int              false "状态"   Enums(0, 1)  default(1)
// @Param   modified_by body     string           true  "修改者"  minlength(3) maxlength(100)
// @Success 200         {array}  model.TagSwagger "成功"
// @Failure 400         {object} errcode.Error    "请求错误"
// @Failure 500         {object} errcode.Error    "内部错误"
// @Router  /api/v1/tags/{id} [put]
func (t Tag) Update(c *gin.Context) {}

// @Summary 删除标签
// @Produce json
// @Param   id  path     int           true "标签 ID"
// @Success 200 {string} string        "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router  /api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {}
