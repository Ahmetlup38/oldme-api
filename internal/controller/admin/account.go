package admin

import (
	"context"
	v1 "oldme-api/api/v1"
	"oldme-api/internal/service"
)

var Account = &cAccount{}

type cAccount struct {
}

// Info 获取账户信息
func (c cAccount) Info(ctx context.Context, req *v1.AccountReq) (res *v1.AccountRes, err error) {
	admin := service.Account().Info(ctx)
	res = &v1.AccountRes{
		Username:  admin.Username,
		Nickname:  admin.Nickname,
		Avatar:    admin.Avatar,
		Register:  admin.Register,
		LastLogin: admin.LastLogin,
	}
	return
}

// Logout 退出登录
func (c cAccount) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	err = service.Account().Logout(ctx)
	return
}
