package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/oldme-git/oldme-api/internal/model"
)

type ArticleListReq struct {
	g.Meta `path:"article/list/*grpId" method:"get" sm:"查询列表" tags:"app"`
	*model.ArticleQuerySafe
}

type ArticleListRes struct {
	List  []model.ArticleListSafe `json:"list"`
	Total uint                    `json:"total"`
}

type ArticleRankReq struct {
	g.Meta `path:"article/rank" method:"get" sm:"查询文章排行" tags:"app"`
	Basis  int `v:"required|in:1,2" dc:"1-热门文章 2-最新文章"`
}

type ArticleRankRes struct {
	List []model.ArticleListSafe `json:"list"`
}

type ArticleShowReq struct {
	g.Meta `path:"article/show/{id}" method:"get" sm:"查询详情" tags:"app"`
	*model.IdInput
}

type ArticleShowRes struct {
	*model.ArticleSafe
}

type AboutShowReq struct {
	g.Meta `path:"about" method:"get" sm:"查询关于我们" tags:"app"`
}

type ArticleHistReq struct {
	g.Meta `path:"article/hist" method:"post" sm:"增加一个点击数" tags:"app"`
	*model.IdInput
}

type ArticleHistRes struct {
}

type ArticleReplyReq struct {
	g.Meta `path:"article/reply/{id}" method:"get" sm:"查看该文章的回复" tags:"app"`
	model.IdInput
}

type ArticleReplyRes struct {
	List  []model.ReplyFloorApp `json:"list"`
	Total uint                  `json:"total"`
}
