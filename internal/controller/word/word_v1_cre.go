package word

import (
	"context"

	"github.com/oldme-git/oldme-api/api/word/v1"
	"github.com/oldme-git/oldme-api/internal/logic/word"
)

func (c *ControllerV1) Cre(ctx context.Context, req *v1.CreReq) (res *v1.CreRes, err error) {
	err = word.Cre(ctx, req.WordInput)
	return
}
