package word

import (
	"context"

	"github.com/oldme-git/oldme-api/api/word/v1"
	"github.com/oldme-git/oldme-api/internal/logic/word"
	"github.com/oldme-git/oldme-api/internal/model"
)

func (c *ControllerV1) Upd(ctx context.Context, req *v1.UpdReq) (res *v1.UpdRes, err error) {
	in := req.WordInput
	in.Id = model.Id(req.IdInput.Id)
	err = word.Upd(ctx, in)
	return
}
