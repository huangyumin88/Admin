// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Account is the golang structure of table apple_account for DAO operations like Where/Data.
type Account struct {
	g.Meta       `orm:"table:apple_account, do:true"`
	Id           interface{} //
	Account      interface{} // 账号
	Pwd          interface{} // 密码
	CountryId    interface{} // 国家id
	CountryCode  interface{} // 国家代码
	Balance      interface{} // 余额
	Dsid         interface{} // Dsid
	Status       interface{} // 禁用：0否 1是
	Info         interface{} // 信息
	Cookies      interface{} // cookies
	LoginStatus  interface{} // 登录：0否 1是
	IsStop       interface{} // 停用：0否 1是
	UpdatedAt    *gtime.Time // 更新时间
	CreatedAt    *gtime.Time // 创建时间
	Stk          interface{} // x-aos-stk
	DeviceId     interface{} // 设备id
	StrTimestamp interface{} // 时间戳
}
