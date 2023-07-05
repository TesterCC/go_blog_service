package service

import (
	"github.com/testercc/blog-service/internal/model"
	"github.com/testercc/blog-service/pkg/app"
)

// 针对入参校验增加绑定/验证结构体
// form 代表着表单的映射字段名
// binding 入参校验的规则内容

/*
主要是定义了 Request 结构体作为接口入参的基准，而本项目由于并不会太复杂，所以直接放在了 service 层中便于使用，
若后续业务不断增长，程序越来越复杂，service 也冗杂了，可以考虑将抽离一层接口校验层，便于解耦逻辑。
*/

type CountTagRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type TagListRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateTagRequest struct {
	Name      string `form:"name" binding:"required,min=1,max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=2,max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateTagRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	Name       string `form:"name" binding:"max=100"`
	State      uint8  `form:"state" binding:"oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,min=2,max=100"`
}

type DeleteTagRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

// 用于处理标签模块的业务逻辑

func (svc *Service) CountTag(param *CountTagRequest) (int, error) {
	return svc.dao.CountTag(param.Name, param.State)
}

func (svc *Service) GetTagList(param *TagListRequest, pager *app.Pager) ([]*model.Tag, error) {
	return svc.dao.GetTagList(param.Name, param.State, pager.Page, pager.PageSize)
}

func (svc *Service) CreateTag(param *CreateTagRequest) error {
	return svc.dao.CreateTag(param.Name, param.State, param.CreatedBy)
}

func (svc *Service) UpdateTag(param *UpdateTagRequest) error {
	return svc.dao.UpdateTag(param.ID, param.Name, param.State, param.ModifiedBy)
}

func (svc *Service) DeleteTag(param *DeleteTagRequest) error {
	return svc.dao.DeleteTag(param.ID)
}

/*
主要针对业务接口中定义的的增删改查和统计行为进行了 Request 结构体编写，而在结构体中，应用到了两个 tag 标签，分别是 form 和 binding，它们分别代表着表单的映射字段名和入参校验的规则内容，其主要功能是实现参数绑定和参数检验。
*/
