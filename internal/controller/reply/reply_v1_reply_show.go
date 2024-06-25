package reply

import (
	"context"

	"github.com/oldme-git/oldme-api/api/reply/v1"
	"github.com/oldme-git/oldme-api/internal/logic/reply"
)

func (c *ControllerV1) ReplyShow(ctx context.Context, req *v1.ReplyShowReq) (res *v1.ReplyShowRes, err error) {
	info, err := reply.Show(ctx, req.Id)
	if err == nil {
		res = &v1.ReplyShowRes{
			ReplyShow: info,
		}
	}
	return
}
