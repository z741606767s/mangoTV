package iface

import v1 "mangoTV/app/query/api/v1"

type IUsersService interface {
	GetUserList(req v1.GetUserListRequest) (res v1.GetUserListResponse, err error)
}
