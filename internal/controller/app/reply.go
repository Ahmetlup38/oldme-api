package app

import (
	"context"
	v1 "github.com/oldme-git/oldme-api/api/app/v1"
	"github.com/oldme-git/oldme-api/internal/model"
	"github.com/oldme-git/oldme-api/internal/model/entity"
	"github.com/oldme-git/oldme-api/internal/service"
)

var Reply = &cReply{}

type cReply struct {
}

func (c *cReply) Reply(ctx context.Context, req *v1.ReplyReq) (res *v1.ReplyRes, err error) {
	_, err = service.Reply().Cre(ctx, &entity.Reply{
		Aid:     int(req.Aid),
		Pid:     int(req.Pid),
		Email:   req.Email,
		Site:    req.Site,
		Name:    req.Name,
		Content: req.Content,
		Status:  model.CheckStatus,
	})
	return
}
