// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Wallets is the golang structure for table wallets.
type Wallets struct {
	WalletId     int         `json:"walletId"     ` // 钱包ID
	UserId       uint        `json:"userId"       ` // 用户ID
	Balance      float64     `json:"balance"      ` // 余额
	RewardPoints uint        `json:"rewardPoints" ` // 积分
	Currency     string      `json:"currency"     ` // 货币 默认尼日利亚 奈拉
	IsStop       uint        `json:"isStop"       ` // 停用：0否 1是
	UpdatedAt    *gtime.Time `json:"updatedAt"    ` // 更新时间
	CreatedAt    *gtime.Time `json:"createdAt"    ` // 创建时间
}
