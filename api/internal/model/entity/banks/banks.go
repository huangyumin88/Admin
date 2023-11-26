// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Banks is the golang structure for table banks.
type Banks struct {
	Id                  int         `json:"id"                  ` // ID
	Name                string      `json:"name"                ` // 银行名称
	Slug                string      `json:"slug"                ` // slug
	Code                string      `json:"code"                ` // code
	FlutterwaveBankCode string      `json:"flutterwaveBankCode" ` // flutterwaveBankCode
	RedbillerBankCode   string      `json:"redbillerBankCode"   ` // redbillerBankCode
	AnchorBankCode      string      `json:"anchorBankCode"      ` // anchorBankCode
	Country             string      `json:"country"             ` // 国家
	Currency            string      `json:"currency"            ` // 货币
	Type                string      `json:"type"                ` // 类型
	IsStop              uint        `json:"isStop"              ` // 停用：0否 1是
	UpdatedAt           *gtime.Time `json:"updatedAt"           ` // 更新时间
	CreatedAt           *gtime.Time `json:"createdAt"           ` // 创建时间
}
