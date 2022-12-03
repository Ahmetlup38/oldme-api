package article_grp

import (
	"context"
	"oldme-api/internal/dao"
	"oldme-api/internal/model"
	"oldme-api/internal/model/do"
	"oldme-api/internal/model/entity"
	"oldme-api/internal/packed"
	"oldme-api/internal/service"
)

type sArticleGrp struct {
}

func init() {
	service.RegisterArticleGrp(&sArticleGrp{})
}

// Cre 创建文章分类
func (s *sArticleGrp) Cre(ctx context.Context, in *model.ArticleGrpInput) (err error) {
	_, err = dao.ArticleGrp.Ctx(ctx).Data(do.ArticleGrp{
		Name:        in.Name,
		Tags:        in.Tags,
		Description: in.Description,
		Onshow:      in.Onshow,
	}).Insert()
	return
}

// Upd 更新文章分类
func (s *sArticleGrp) Upd(ctx context.Context, id uint, in *model.ArticleGrpInput) (err error) {
	_, err = dao.ArticleGrp.Ctx(ctx).Data(do.ArticleGrp{
		Name:        in.Name,
		Tags:        in.Tags,
		Description: in.Description,
		Onshow:      in.Onshow,
	}).Where("id", id).Update()
	return
}

// Del 删除文章分类
func (s *sArticleGrp) Del(ctx context.Context, id uint) (err error) {
	_, err = dao.ArticleGrp.Ctx(ctx).Where("id", id).Delete()
	return
}

// List 读取文章分类列表
func (s *sArticleGrp) List(ctx context.Context) (data *model.ArticleGrpList, err error) {
	data = &model.ArticleGrpList{}
	result, err := dao.ArticleGrp.Ctx(ctx).All()
	_ = result.Structs(data)
	return
}

// Show 读取文章分类详情
func (s *sArticleGrp) Show(ctx context.Context, id uint) (data *entity.ArticleGrp, err error) {
	err = dao.ArticleGrp.Ctx(ctx).Where("id", id).Scan(&data)
	if err != nil {
		err = packed.Code.SetErr(10100)
	}
	return
}

// IsExist 根据id判断一个文章分类是否存在
func (s *sArticleGrp) IsExist(ctx context.Context, id uint) bool {
	num, _ := dao.ArticleGrp.Ctx(ctx).Where("id", id).Count()
	return num == 1
}
