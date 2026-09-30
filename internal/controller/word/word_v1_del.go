package word

import (
	"context"

	"github.com/oldme-git/oldme-api/api/word/v1"
	"github.com/oldme-git/oldme-api/internal/logic/word"
)

func (c *ControllerV1) Del(ctx context.Context, req *v1.DelReq) (res *v1.DelRes, err error) {
	err = word.Del(ctx, req.IdInput.Id)
	return
}
