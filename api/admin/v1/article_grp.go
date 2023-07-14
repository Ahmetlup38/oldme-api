package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"oldme-api/internal/model"
	"oldme-api/internal/model/entity"
)

type ArticleGrpCreReq struct {
	g.Meta `path:"article/group/create" method:"post" sm:"新增" tags:"文章分类"`
	*model.ArticleGrpInput
}

type ArticleGrpCreRes struct {
}

type ArticleGrpUpdReq struct {
	g.Meta `path:"article/group/update/{id}" method:"post" sm:"修改" tags:"文章分类"`
	*model.IdInput
	*model.ArticleGrpInput
}

type ArticleGrpUpdRes struct {
}

type ArticleGrpDelReq struct {
	g.Meta `path:"article/group/delete/{id}" method:"post" sm:"删除" tags:"文章分类"`
	*model.IdInput
}

type ArticleGrpDelRes struct {
}

type ArticleGrpListReq struct {
	g.Meta `path:"article/group/list" method:"get" sm:"查询列表" tags:"文章分类"`
}

type ArticleGrpListRes struct {
	List  *[]entity.ArticleGrp `json:"list"`
	Total uint                 `json:"total"`
}

type ArticleGrpShowReq struct {
	g.Meta `path:"article/group/show/{id}" method:"get" sm:"查询详情" tags:"文章分类"`
	*model.IdInput
}

type ArticleGrpShowRes struct {
	*entity.ArticleGrp
}
