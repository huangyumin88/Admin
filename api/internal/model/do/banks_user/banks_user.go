// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BanksUser is the golang structure of table app_card_banks_user for DAO operations like Where/Data.
type BanksUser struct {
	g.Meta    `orm:"table:app_card_banks_user, do:true"`
	Id        interface{} // ID
	UserId    interface{} // 用户id
	BankId    interface{} // 银行id
	UpdatedAt *gtime.Time // 更新时间
	CreatedAt *gtime.Time // 创建时间
}
