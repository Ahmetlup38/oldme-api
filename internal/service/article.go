// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"oldme-api/internal/model"
	"oldme-api/internal/model/entity"
)

type (
	IArticle interface {
		Cre(ctx context.Context, in *model.ArticleInput) (lastId uint, err error)
		Upd(ctx context.Context, id uint, in *model.ArticleInput) (err error)
		List(ctx context.Context, query *model.ArticleQuery) (list *[]model.ArticleList, total uint, err error)
		Show(ctx context.Context, id uint) (info *entity.Article, err error)
		Del(ctx context.Context, id uint, isReal bool) (err error)
		ReCre(ctx context.Context, id uint) (err error)
		Hist(ctx context.Context, id uint) (err error)
		UptLastedAt(ctx context.Context, id uint) (err error)
	}
)

var (
	localArticle IArticle
)

func Article() IArticle {
	if localArticle == nil {
		panic("implement not found for interface IArticle, forgot register?")
	}
	return localArticle
}

func RegisterArticle(i IArticle) {
	localArticle = i
}
