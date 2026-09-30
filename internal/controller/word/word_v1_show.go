package word

import (
	"context"

	"github.com/oldme-git/oldme-api/api/word/v1"
	"github.com/oldme-git/oldme-api/internal/logic/word"
)

func (c *ControllerV1) Show(ctx context.Context, req *v1.ShowReq) (res *v1.ShowRes, err error) {
	info, err := word.Show(ctx, req.IdInput.Id)
	if err != nil {
		return nil, err
	}
	return &v1.ShowRes{
		Word: info,
	}, nil
}
