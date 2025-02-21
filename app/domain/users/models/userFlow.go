package models

import (
	"gorm.io/gorm"
	"time"
)

// 交易状态 (Status)
const (
	StatusPending   = iota + 1 // 待处理
	StatusSuccess              // 成功
	StatusFailed               // 失败
	StatusCancelled            // 已取消
	StatusRefunded             // 已退款
)

// 交易类型 (TransactionType)
const (
	TransactionTypeIncome  = iota + 1 // 收入
	TransactionTypeExpense            // 支出
	TransactionTypeRefund             // 退款
)

// 交易子类型 (SubType)
const (
	// IncomeSubTypeRecharge 收入子类型 (IncomeSubType)
	IncomeSubTypeRecharge       int = 101 // 充值
	IncomeSubTypeActivityReward int = 102 // 活动奖励

	// ExpenseSubTypePurchase 支出子类型 (ExpenseSubType)
	ExpenseSubTypePurchase int = 201 // 消费
	ExpenseSubTypeTransfer int = 202 // 转账

	// RefundSubTypeOrder 退款子类型 (RefundSubType)
	RefundSubTypeOrder int = 501 // 订单退款
	RefundSubTypeOther int = 502 // 其他退款
)

// Source 交易来源 (Source)
const (
	SourceSystem = iota + 1 // 系统发放
	SourceUser              // 用户触发
	SourceAdmin             // 管理员操作
)

type UserFlows struct {
	ID                 int64     `gorm:"primary_key;auto_increment;column:id;comment:'流水ID'" json:"id"`
	UserID             int64     `gorm:"not null;column:user_id;comment:'用户ID';index:idx_user_id" json:"userId"`
	TransactionType    int       `gorm:"not null;column:transaction_type;comment:'交易类型（1: 收入, 2: 支出）';index:idx_transaction_type,priority:1" json:"transactionType"`
	SubType            int       `gorm:"not null;column:sub_type;comment:'交易子类型（如收入中101: 充值, 102: 活动奖励；支出中201: 消费, 202: 转账）';index:idx_transaction_type,priority:2" json:"subType"`
	NovelsId           int64     `gorm:"column:novels_id;comment:'书籍ID'" json:"novelsId,omitempty"`
	ChapterId          int64     `gorm:"column:chapter_id;comment:'书籍章节ID'" json:"chapterId,omitempty"`
	Coins              int64     `gorm:"type:bigint unsigned;not null;column:coins;comment:'本次交易金币'" json:"coins"`
	BeforeCoins        int64     `gorm:"type:bigint unsigned;not null;column:before_coins;comment:'交易前金币余额'" json:"beforeCoins"`
	AfterCoins         int64     `gorm:"type:bigint unsigned;not null;column:after_coins;comment:'交易后金币余额'" json:"afterCoins"`
	ReferenceID        *string   `gorm:"type:varchar(100);default:null;column:reference_id;comment:'关联的业务ID（如订单ID、奖励记录ID）';index:idx_reference_id" json:"referenceId,omitempty"`
	Source             int       `gorm:"not null;column:source;comment:'交易来源（1: 系统发放, 2: 用户触发, 3: 管理员操作）'" json:"source"`
	OperatorID         *int64    `gorm:"type:bigint unsigned;default:null;column:operator_id;comment:'操作人ID（管理员或系统用户ID）'" json:"operatorId,omitempty"`
	Status             int       `gorm:"not null;default:1;type:tinyint;column:status;comment:交易状态（1: 待处理, 2: 成功, 3: 失败, 4: 已取消, 5: 已退款）" json:"status"`
	Description        *string   `gorm:"type:varchar(255);default:null;column:description;comment:'交易描述'" json:"description,omitempty"`
	CreatedAt          time.Time `gorm:"type:datetime;not null;column:created_at;default:CURRENT_TIMESTAMP;comment:'交易时间';index:idx_created_at" json:"-"`
	FormattedCreatedAt string    `gorm:"-" json:"createdAt"`
	UpdatedAt          time.Time `gorm:"type:datetime;not null;column:updated_at;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP;comment:'更新时间'" json:"-"`
	FormattedUpdatedAt string    `gorm:"-" json:"updatedAt"`
}

func (u *UserFlows) TableName() string {
	return "c_user_flows"
}

func (u *UserFlows) Table() map[string]string {
	return map[string]string{
		"ENGINE": "InnoDB", "COMMENT": "用户流水表",
	}
}

func (u *UserFlows) AfterFind(tx *gorm.DB) (err error) {
	u.FormattedCreatedAt = u.CreatedAt.Format(time.DateTime)
	u.FormattedUpdatedAt = u.UpdatedAt.Format(time.DateTime)
	return nil
}

func GetStatusDescription(t int) string {
	switch t {
	case StatusPending:
		return "待处理"
	case StatusSuccess:
		return "成功"
	case StatusFailed:
		return "失败"
	case StatusCancelled:
		return "已取消"
	case StatusRefunded:
		return "已退款"
	default:
		return "未知类型"
	}
}

// GetTransactionTypeDescription 返回交易类型的描述
func GetTransactionTypeDescription(t int) string {
	switch t {
	case TransactionTypeIncome:
		return "收入"
	case TransactionTypeExpense:
		return "支出"
	case TransactionTypeRefund:
		return "退款"
	default:
		return "未知类型"
	}
}

// GetSubTypeDescription 返回交易子类型的描述
func GetSubTypeDescription(t int) string {
	switch t {
	case IncomeSubTypeRecharge:
		return "充值"
	case IncomeSubTypeActivityReward:
		return "活动奖励"
	case ExpenseSubTypePurchase:
		return "消费"
	case ExpenseSubTypeTransfer:
		return "转账"
	case RefundSubTypeOrder:
		return "订单退款"
	case RefundSubTypeOther:
		return "其他退款"
	default:
		return "未知子类型"
	}
}

// GetSourceDescription 返回交易来源描述
func GetSourceDescription(t int) string {
	switch t {
	case SourceSystem:
		return "系统发放"
	case SourceUser:
		return "用户触发"
	case SourceAdmin:
		return "管理员操作"
	default:
		return "未知类型"
	}
}
