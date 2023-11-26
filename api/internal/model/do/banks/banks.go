// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Banks is the golang structure of table app_card_banks for DAO operations like Where/Data.
type Banks struct {
	g.Meta              `orm:"table:app_card_banks, do:true"`
	Id                  interface{} // ID
	Name                interface{} // 银行名称
	Slug                interface{} // slug
	Code                interface{} // code
	FlutterwaveBankCode interface{} // flutterwaveBankCode
	RedbillerBankCode   interface{} // redbillerBankCode
	AnchorBankCode      interface{} // anchorBankCode
	Country             interface{} // 国家
	Currency            interface{} // 货币
	Type                interface{} // 类型
	IsStop              interface{} // 停用：0否 1是
	UpdatedAt           *gtime.Time // 更新时间
	CreatedAt           *gtime.Time // 创建时间
}
