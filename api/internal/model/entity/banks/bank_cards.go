// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BankCards is the golang structure for table bank_cards.
type BankCards struct {
	Id             int         `json:"id"             ` // ID
	UserId         uint        `json:"userId"         ` // 用户id
	BankId         int         `json:"bankId"         ` // 银行id
	CardNumber     string      `json:"cardNumber"     ` // 银行卡号
	CardHolderName string      `json:"cardHolderName" ` // 持卡人姓名
	ExpirationDate *gtime.Time `json:"expirationDate" ` // 银行卡的有效期
	Cvv            string      `json:"cvv"            ` // CVV 安全码
	UpdatedAt      *gtime.Time `json:"updatedAt"      ` // 更新时间
	CreatedAt      *gtime.Time `json:"createdAt"      ` // 创建时间
}
