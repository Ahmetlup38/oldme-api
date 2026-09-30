package word

import (
	"context"

	"github.com/gogf/gf/v2/util/grand"

	"github.com/oldme-git/oldme-api/internal/dao"
	"github.com/oldme-git/oldme-api/internal/model"
	"github.com/oldme-git/oldme-api/internal/model/do"
	"github.com/oldme-git/oldme-api/internal/model/entity"
	"github.com/oldme-git/oldme-api/internal/utility"
)

// Cre 创建词汇
func Cre(ctx context.Context, in *model.WordInput) (err error) {
	_, err = dao.Word.Ctx(ctx).Data(do.Word{
		Word:    in.Word,
		Explain: in.Explain,
	}).Insert()
	if err != nil {
		err = utility.Err.Sys(err)
	}
	return
}

// Upd 更新词汇
func Upd(ctx context.Context, in *model.WordInput) (err error) {
	_, err = dao.Word.Ctx(ctx).Data(do.Word{
		Word:    in.Word,
		Explain: in.Explain,
	}).Where("id", in.Id).Update()
	if err != nil {
		err = utility.Err.Sys(err)
	}
	return
}

// Del 删除词汇
func Del(ctx context.Context, id model.Id) (err error) {
	_, err = dao.Word.Ctx(ctx).Where("id", id).Delete()
	if err != nil {
		err = utility.Err.Sys(err)
	}
	return
}

// Show 读取词汇详情
func Show(ctx context.Context, id model.Id) (info *entity.Word, err error) {
	info = &entity.Word{}
	err = dao.Word.Ctx(ctx).Where("id", id).Scan(&info)
	if err != nil {
		err = utility.Err.TableNotData(err)
	}
	return
}

// List 读取词汇列表
func List(ctx context.Context, query *model.WordQuery) (list []entity.Word, total uint, err error) {
	if query == nil {
		query = &model.WordQuery{}
	}
	// 对于查询初始值的处理
	if query.Page == 0 {
		query.Page = 1
	}
	if query.Size == 0 {
		query.Size = 15
	}

	db := dao.Word.Ctx(ctx)
	if query.Search != "" {
		db = db.Where("word like ?", "%"+query.Search+"%")
	}
	data, totalInt, err := db.Page(query.Page, query.Size).OrderDesc("id").AllAndCount(true)
	if err != nil {
		err = utility.Err.Sys(err)
		return
	}
	list = []entity.Word{}
	_ = data.Structs(&list)
	total = uint(totalInt)

	return
}

// ReadDay 随机读取一些词
func ReadDay(ctx context.Context, readNum int) (list []entity.Word, err error) {
	// 读取总数
	count, err := dao.Word.Ctx(ctx).
		Fields("id").
		Count()
	if err != nil {
		err = utility.Err.Sys(err)
		return
	}
	// 随机偏移量
	offset := grand.Intn(int(count)) - readNum
	if offset < 0 {
		offset = 0
	}

	// 读取随机句子id
	data, err := dao.Word.Ctx(ctx).
		Fields("word", "explain").
		Limit(readNum).
		Offset(offset).
		All()

	if err != nil {
		err = utility.Err.Sys(err)
	}
	list = []entity.Word{}
	_ = data.Structs(&list)
	return
}
