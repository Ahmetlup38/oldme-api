package admin

import (
	"context"
	v1 "oldme-api/api/v1"
	"oldme-api/internal/model"
	"oldme-api/internal/model/entity"
	"oldme-api/internal/service"
)

var ArticleGrp = cArticleGrp{}

type cArticleGrp struct {
}

func (c *cArticleGrp) Cre(ctx context.Context, req *v1.ArticleGrpCreReq) (res *v1.ArticleGrpCreRes, err error) {
	err = service.ArticleGrp().Cre(ctx, model.ArticleGrpInput{
		Name:        req.Name,
		Tags:        req.Tags,
		Description: req.Description,
		Onshow:      req.Onshow,
	})
	return
}

func (c *cArticleGrp) Upd(ctx context.Context, req *v1.ArticleGrpUpdReq) (res *v1.ArticleGrpUpdRes, err error) {
	err = service.ArticleGrp().Upd(ctx, req.Id, model.ArticleGrpInput{
		Name:        req.Name,
		Tags:        req.Tags,
		Description: req.Description,
		Onshow:      req.Onshow,
	})
	return
}

func (c *cArticleGrp) Del(ctx context.Context, req *v1.ArticleGrpDelReq) (res *v1.ArticleGrpDelRes, err error) {
	err = service.ArticleGrp().Del(ctx, req.Id)
	return
}

func (c *cArticleGrp) List(ctx context.Context, req *v1.ArticleGrpListReq) (res *v1.ArticleGroListRes, err error) {
	data, err := service.ArticleGrp().List(ctx)
	if err == nil {
		res = &v1.ArticleGroListRes{
			List:  data,
			Total: uint(len(data)),
		}
	}
	return
}

func (c *cArticleGrp) Show(ctx context.Context, req *v1.ArticleGrpShowReq) (res *v1.ArticleGrpShowRes, err error) {
	var data entity.ArticleGrp
	data, err = service.ArticleGrp().Show(ctx, req.Id)
	if err == nil {
		res = &v1.ArticleGrpShowRes{
			ArticleGrp: data,
		}
	}
	return
}
