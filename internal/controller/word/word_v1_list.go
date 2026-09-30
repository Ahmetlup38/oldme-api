package word

import (
	"context"

	"github.com/oldme-git/oldme-api/api/word/v1"
	"github.com/oldme-git/oldme-api/internal/logic/word"
	"github.com/oldme-git/oldme-api/internal/model"
)

func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	if req.Paging == nil {
		req.Paging = &model.Paging{
			Page: 1,
			Size: 15,
		}
	}

	query := &model.WordQuery{
		Paging: *req.Paging,
		Search: req.Search,
	}

	list, total, err := word.List(ctx, query)
	if err != nil {
		return nil, err
	}
	return &v1.ListRes{
		List:  list,
		Total: total,
	}, nil
}
