// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BankCards is the golang structure of table app_card_bank_cards for DAO operations like Where/Data.
type BankCards struct {
	g.Meta         `orm:"table:app_card_bank_cards, do:true"`
	Id             interface{} // ID
	UserId         interface{} // 用户id
	BankId         interface{} // 银行id
	CardNumber     interface{} // 银行卡号
	CardHolderName interface{} // 持卡人姓名
	ExpirationDate *gtime.Time // 银行卡的有效期
	Cvv            interface{} // CVV 安全码
	UpdatedAt      *gtime.Time // 更新时间
	CreatedAt      *gtime.Time // 创建时间
}
