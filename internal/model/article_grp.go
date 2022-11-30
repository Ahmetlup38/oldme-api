package model

import "oldme-api/internal/model/entity"

type ArticleGrpList []entity.ArticleGrp

type ArticleGrpInput struct {
	Name        string `v:"required|length:2, 30"`
	Tags        string `v:"length:1, 200"`
	Description string `v:"length:2, 200"`
	Onshow      bool   `v:"required"`
}
