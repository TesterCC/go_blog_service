package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/testercc/blog-service/global"
	"github.com/testercc/blog-service/internal/service"
	"github.com/testercc/blog-service/pkg/app"
	"github.com/testercc/blog-service/pkg/convert"
	"github.com/testercc/blog-service/pkg/errcode"
)

type Tag struct {
}

func NewTag() Tag {
	return Tag{}
}

//func (t Tag) Get(c *gin.Context)    {}

//todo: 标签模块的接口注解编写，接下来应当按照注解的含义和参考上述接口注解，完成文章模块接口注解的编写。
//  curl -X GET 'http://127.0.0.1:8888/api/v1/tags?page=1&page_size=10

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
	param := service.TagListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	totalRows, err := svc.CountTag(&service.CountTagRequest{Name: param.Name, State: param.State})
	if err != nil {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.ErrorCountTagFail)
		return
	}

	tags, err := svc.GetTagList(&param, &pager)
	if err != nil {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}

	response.ToResponseList(tags, totalRows)
	return
}

/*
在方法中完成了入参校验和绑定、获取标签总数、获取标签列表、 序列化结果集等四大功能板块的逻辑串联和日志、错误处理。
*/

// testcase:  curl -X POST http://127.0.0.1:8888/api/v1/tags -F 'name=Go' -F created_by=testercc

// @Summary 新增标签
// @Produce json
// @Param   name       body     string           true  "标签名称" minlength(3) maxlength(100)
// @Param   state      body     int              false "状态"   Enums(0, 1)  default(1)
// @Param   created_by body     string           true  "创建者"  minlength(3) maxlength(100)
// @Success 200        {object} model.TagSwagger "成功"
// @Failure 400        {object} errcode.Error    "请求错误"
// @Failure 500        {object} errcode.Error    "内部错误"
// @Router  /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {
	param := service.CreateTagRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CreateTag(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.CreateTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

// testcase: curl -X PUT http://127.0.0.1:8888/api/v1/tags/4 -F state=0 -F modified_by=hacker

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
func (t Tag) Update(c *gin.Context) {
	param := service.UpdateTagRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.UpdateTag(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.UpdateTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateTagFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}


// testcase: curl -X DELETE http://127.0.0.1:8888/api/v1/tags/5

// @Summary 删除标签
// @Produce json
// @Param   id  path     int           true "标签 ID"
// @Success 200 {string} string        "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router  /api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {
	param := service.DeleteTagRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.DeleteTag(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.DeleteTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteTagFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}
