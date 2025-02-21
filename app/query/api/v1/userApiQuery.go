package v1

import "mangoTV/app/domain/users/models"

type GetUserListRequest struct {
	Limit  int    `form:"limit" json:"limit"`   // 每页数量
	Page   int    `form:"page" json:"page"`     // 页码
	ID     int64  `form:"id" json:"id"`         // 用户ID
	STime  int64  `form:"sTime" json:"sTime"`   // 开始时间
	ETime  int64  `form:"eTime" json:"eTime"`   // 结束时间
	UUID   string `form:"uuid" json:"uuid"`     // UUID唯一标识
	IsVip  int8   `form:"isVip" json:"isVip"`   // 1:是会员 2:会员或过期 3:非会员
	Status int8   `form:"status" json:"status"` // 1:正常 2:禁用
}

type GetUserListResponse struct {
	Total int             `json:"total"` // 总数
	Limit int             `json:"limit"` // 每页数量
	Page  int             `json:"page"`  // 页码
	Items []*models.Users `json:"items"` // 列表
}
