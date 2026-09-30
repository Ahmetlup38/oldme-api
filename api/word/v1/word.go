package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/oldme-git/oldme-api/internal/model"
	"github.com/oldme-git/oldme-api/internal/model/entity"
)

type CreReq struct {
	g.Meta `path:"word/create" method:"post" sm:"新增" tags:"词汇"`
	*model.WordInput
}

type CreRes struct {
}

type UpdReq struct {
	g.Meta `path:"word/update/{id}" method:"post" sm:"修改" tags:"词汇"`
	*model.IdInput
	*model.WordInput
}

type UpdRes struct {
}

type ShowReq struct {
	g.Meta `path:"word/show/{id}" method:"get" sm:"查询详情" tags:"词汇"`
	*model.IdInput
}

type ShowRes struct {
	*entity.Word
}

type DelReq struct {
	g.Meta `path:"word/delete/{id}" method:"post" sm:"删除" tags:"词汇"`
	*model.IdInput
}

type DelRes struct {
}

type ListReq struct {
	g.Meta `path:"word/list" method:"get" sm:"查询列表" tags:"词汇"`
	*model.Paging
	Search string `json:"search" dc:"搜索文本"`
}

type ListRes struct {
	List  []entity.Word `json:"list"`
	Total uint          `json:"total"`
}
