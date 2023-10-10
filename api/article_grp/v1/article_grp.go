package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/oldme-git/oldme-api/internal/model"
	"github.com/oldme-git/oldme-api/internal/model/entity"
)

type CreReq struct {
	g.Meta `path:"article/group/create" method:"post" sm:"新增" tags:"文章分类"`
	*model.ArticleGrpInput
}

type CreRes struct {
}

type UpdReq struct {
	g.Meta `path:"article/group/update/{id}" method:"post" sm:"修改" tags:"文章分类"`
	*model.IdInput
	*model.ArticleGrpInput
}

type UpdRes struct {
}

type DelReq struct {
	g.Meta `path:"article/group/delete/{id}" method:"post" sm:"删除" tags:"文章分类"`
	*model.IdInput
}

type DelRes struct {
}

type ListReq struct {
	g.Meta `path:"article/group/list" method:"get" sm:"查询列表" tags:"文章分类"`
}

type ListRes struct {
	List  []entity.ArticleGrp `json:"list"`
	Total uint                `json:"total"`
}

type ShowReq struct {
	g.Meta `path:"article/group/show/{id}" method:"get" sm:"查询详情" tags:"文章分类"`
	*model.IdInput
}

type ShowRes struct {
	*entity.ArticleGrp
}
