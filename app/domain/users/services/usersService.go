package services

import (
	"github.com/sirupsen/logrus"
	v1 "mangoTV/app/query/api/v1"
	"mangoTV/app/svc"
)

type UsersService struct {
	ctx *svc.ServiceContext
}

func NewUsersService(ctx *svc.ServiceContext) *UsersService {
	return &UsersService{
		ctx: ctx,
	}
}

// GetUserList 获取用户列表
func (u *UsersService) GetUserList(req v1.GetUserListRequest) (res v1.GetUserListResponse, err error) {
	if req.Limit == 0 {
		req.Limit = 10
		res.Limit = 10
	}

	userList, total, err := u.ctx.Provider.UsersDao.GetUserList(req.Limit, req.Page, req.ID, req.UUID, req.IsVip, req.Status, req.STime, req.ETime)

	if err != nil {
		logrus.Errorf("get user list failed, Err:[%+v]", err)
		return res, err
	}

	res.Total = int(total)
	res.Page = req.Page
	res.Items = userList

	return res, nil
}
