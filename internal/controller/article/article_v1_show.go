package article

import (
	"context"
	"github.com/oldme-git/oldme-api/internal/service"

	"github.com/oldme-git/oldme-api/api/article/v1"
)

func (c *ControllerV1) Show(ctx context.Context, req *v1.ShowReq) (res *v1.ShowRes, err error) {
	info, err := service.Article().Show(ctx, req.Id)
	if err == nil {
		res = &v1.ShowRes{
			Article: info,
		}
	}
	return
}
