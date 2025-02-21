package models

import (
	"gorm.io/gorm"
	"time"
)

const (
	VipFalse = iota
	VipTrue
)

const (
	AutoUnlockFalse = iota
	AutoUnlockTrue
)

const (
	StatusUnavailable = iota
	StatusAvailable
	StatusBlacklist
)

type Users struct {
	ID                 int64             `gorm:"primary_key;auto_increment;column:id;comment:'ID'" json:"id"`
	UUID               string            `gorm:"size:160;not null;unique;column:uuid;comment:'用户唯一识别'" json:"uuid"`
	Email              string            `gorm:"size:160;not null;column:email;comment:'邮箱'" json:"email"`
	Nickname           string            `gorm:"size:160;not null;column:nickname;comment:'用户昵称'" json:"nickname"`
	Avatar             string            `gorm:"size:160;column:avatar;comment:'管理员头像'" json:"-"`
	FormattedAvatar    string            `gorm:"-" json:"avatar"`
	Coins              int64             `gorm:"type:bigint unsigned;default:0;column:coins;comment:'金币余额'" json:"coins"`
	IsAutoUnlock       int8              `gorm:"not null;default:1;column:is_auto_unlock;comment:'是否自动金币解锁：0否 1是'" json:"isAutoUnlock"`
	IsVip              int8              `gorm:"not null;default:0;column:is_vip;comment:'是否VIP权限：0否 1是'" json:"isVip"`
	VipEx              time.Time         `gorm:"type:datetime;column:vip_ex;default:null" json:"-"`
	FormattedVipEx     string            `gorm:"-" json:"vipEx"`
	Status             int8              `gorm:"not null;default:1;column:status;comment:'状态：0禁用 1启用 2黑名单'" json:"status"`
	Ttclid             string            `gorm:"size:160;column:ttclid;comment:'TikTok广告 ttclid点击编号'" json:"ttclid"`
	Ttp                string            `gorm:"size:160;column:ttp;comment:'TikTok广告 ttp点击编号'" json:"ttp"`
	Fbc                string            `gorm:"size:160;column:fbc;comment:'Facebook广告 fbc点击编号'" json:"fbc"`
	Fbp                string            `gorm:"size:160;column:fbp;comment:'Facebook广告 fbp浏览器编号'" json:"fbp"`
	Fbclid             string            `gorm:"size:160;column:fbclid;comment:'广告fbclid'" json:"fbclid"`
	FbLoginId          string            `gorm:"size:160;default:0;column:fb_login_id;comment:'广告,用户的Facebook 登录编号'" json:"fbLoginId"`
	CampaignName       string            `gorm:"size:160;column:campaign_name;comment:'广告活动名称'" json:"campaignName"`
	CampaignId         string            `gorm:"size:160;column:campaign_id;comment:'广告活动ID'" json:"campaignId"`
	AdsetName          string            `gorm:"size:160;column:adset_name;comment:'广告集名称'" json:"adsetName"`
	AdName             string            `gorm:"size:160;column:ad_name;comment:'广告名称'" json:"adName"`
	AdId               string            `gorm:"size:160;column:ad_id;comment:'广告ID'" json:"adId"`
	AdsetId            string            `gorm:"size:160;column:adset_d;comment:'广告集ID'" json:"adsetId"`
	NovelId            int64             `gorm:"column:novel_id;comment:'广告小说ID'" json:"novelId"`
	ChapterId          int64             `gorm:"not null;column:chapter_id;comment:'书籍章节ID'" json:"chapterId"`
	AffiliateLinkId    int64             `gorm:"column:affiliate_link_id;comment:'广告推广链ID'" json:"affiliateLinkId"`
	NovelArr           map[int64][]int64 `gorm:"-" json:"novelArr"`
	CreatedAt          time.Time         `gorm:"type:datetime;not null;column:created_at;default:CURRENT_TIMESTAMP" json:"-"`
	FormattedCreatedAt string            `gorm:"-" json:"createdAt"`
	UpdatedAt          time.Time         `gorm:"type:datetime;not null;column:updated_at;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP" json:"-"`
	FormattedUpdatedAt string            `gorm:"-" json:"updatedAt"`
}

func (u *Users) Table() map[string]string {
	return map[string]string{
		"ENGINE": "InnoDB", "COMMENT": "用户表",
	}
}

func (u *Users) TableName() string {
	return "c_users"
}

type UsersInfoItem struct {
	ID                 int64             `gorm:"primary_key;auto_increment;column:id;comment:'ID'" json:"id"`
	UUID               string            `gorm:"size:160;not null;unique;column:uuid;comment:'用户唯一识别'" json:"uuid"`
	Email              string            `gorm:"size:160;not null;column:email;comment:'邮箱'" json:"email"`
	Nickname           string            `gorm:"size:160;not null;column:nickname;comment:'用户昵称'" json:"nickname"`
	Avatar             string            `gorm:"size:160;column:avatar;comment:'管理员头像'" json:"-"`
	FormattedAvatar    string            `gorm:"-" json:"avatar"`
	Coins              int64             `gorm:"default:0;column:coins;comment:'金币余额'" json:"coins"`
	IsAutoUnlock       int8              `gorm:"not null;default:1;column:is_auto_unlock;comment:'是否自动金币解锁：0否 1是'" json:"isAutoUnlock"`
	IsVip              int8              `gorm:"not null;default:0;column:is_vip;comment:'是否VIP权限：0否 1是'" json:"isVip"`
	VipEx              time.Time         `gorm:"type:datetime;column:vip_ex;default:null" json:"-"`
	FormattedVipEx     string            `gorm:"-" json:"vipEx"`
	Status             int8              `gorm:"not null;default:1;column:status;comment:'状态：0禁用 1启用 2黑名单'" json:"status"`
	Ttclid             string            `gorm:"size:160;column:ttclid;comment:'TikTok广告 ttclid点击编号'" json:"ttclid"`
	Ttp                string            `gorm:"size:160;column:ttp;comment:'TikTok广告 ttp点击编号'" json:"ttp"`
	Fbc                string            `gorm:"size:160;column:fbc;comment:'广告fbc点击编号'" json:"fbc"`
	Fbp                string            `gorm:"size:160;column:fbp;comment:'广告fbp浏览器编号'" json:"fbp"`
	Fbclid             string            `gorm:"size:160;column:fbclid;comment:'广告fbclid'" json:"fbclid"`
	FbLoginId          string            `gorm:"size:160;default:0;column:fb_login_id;comment:'广告fbLoginId'" json:"fbLoginId"`
	CampaignName       string            `gorm:"size:160;column:campaign_name;comment:'广告活动名称'" json:"campaignName"`
	CampaignId         string            `gorm:"size:160;column:campaign_id;comment:'广告活动ID'" json:"campaignId"`
	AdsetName          string            `gorm:"size:160;column:adset_name;comment:'广告集名称'" json:"adsetName"`
	AdName             string            `gorm:"size:160;column:ad_name;comment:'广告名称'" json:"adName"`
	AdId               string            `gorm:"size:160;column:ad_id;comment:'广告ID'" json:"adId"`
	AdsetId            string            `gorm:"size:160;column:adset_d;comment:'广告集ID'" json:"adsetId"`
	NovelId            int64             `gorm:"column:novel_id;comment:'广告小说ID'" json:"novelId"`
	ChapterId          int64             `gorm:"not null;column:chapter_id;comment:'书籍章节ID'" json:"chapterId"`
	AffiliateLinkId    int64             `gorm:"column:affiliate_link_id;comment:'广告推广链ID'" json:"affiliateLinkId"`
	NovelArr           map[int64][]int64 `gorm:"-" json:"novelArr"`
	CreatedAt          time.Time         `gorm:"type:datetime;not null;column:created_at;default:CURRENT_TIMESTAMP" json:"-"`
	FormattedCreatedAt string            `gorm:"-" json:"createdAt"` // 用于存储格式化后的时间
	UpdatedAt          time.Time         `gorm:"type:datetime;not null;column:updated_at;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP" json:"-"`
	FormattedUpdatedAt string            `gorm:"-" json:"updatedAt"` // 用于存储格式化后的时间
	Device             []UsersDevice     `gorm:"foreignKey:UserID;references:ID" json:"device"`
}

// AfterFind 是一个 GORM 钩子，在查询记录后调用
func (u *Users) AfterFind(tx *gorm.DB) (err error) {
	u.FormattedCreatedAt = u.CreatedAt.Format(time.DateTime)
	u.FormattedUpdatedAt = u.UpdatedAt.Format(time.DateTime)

	// 如果 vipEx 是 "0001-01-01 00:00:00"，将 FormattedVipEx 设置为 "0"
	if u.VipEx.IsZero() {
		u.FormattedVipEx = "0"
	} else {
		u.FormattedVipEx = u.VipEx.Format(time.DateTime)
	}

	u.NovelArr = make(map[int64][]int64)

	var userFlows []*UserFlows
	if err = tx.Where("user_id = ? AND transaction_type = ? AND sub_type = ? AND status = ?", u.ID, TransactionTypeExpense, ExpenseSubTypePurchase, StatusSuccess).Find(&userFlows).Error; err != nil {
		return err
	}

	for _, userFlow := range userFlows {
		if userFlow.NovelsId > 0 && userFlow.ChapterId > 0 {
			if _, exists := u.NovelArr[userFlow.NovelsId]; !exists {
				u.NovelArr[userFlow.NovelsId] = []int64{}
			}
			if !contains(u.NovelArr[userFlow.NovelsId], userFlow.ChapterId) {
				u.NovelArr[userFlow.NovelsId] = append(u.NovelArr[userFlow.NovelsId], userFlow.ChapterId)
			}
		}
	}

	return nil
}

func contains(arr []int64, value int64) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}
