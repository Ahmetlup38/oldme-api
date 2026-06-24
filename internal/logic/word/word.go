package word

import (
	"context"

	"github.com/gogf/gf/v2/util/grand"

	"github.com/oldme-git/oldme-api/internal/dao"
	"github.com/oldme-git/oldme-api/internal/model/entity"
	"github.com/oldme-git/oldme-api/internal/utility"
)

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
