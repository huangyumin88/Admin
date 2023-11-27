// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Wallets is the golang structure of table app_card_wallets for DAO operations like Where/Data.
type Wallets struct {
	g.Meta       `orm:"table:app_card_wallets, do:true"`
	WalletId     interface{} // 钱包ID
	UserId       interface{} // 用户ID
	Balance      interface{} // 余额
	RewardPoints interface{} // 积分
	Currency     interface{} // 货币 默认尼日利亚 奈拉
	IsStop       interface{} // 停用：0否 1是
	UpdatedAt    *gtime.Time // 更新时间
	CreatedAt    *gtime.Time // 创建时间
}
