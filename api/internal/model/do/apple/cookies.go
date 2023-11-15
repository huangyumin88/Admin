// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Cookies is the golang structure of table apple_cookies for DAO operations like Where/Data.
type Cookies struct {
	g.Meta       `orm:"table:apple_cookies, do:true"`
	Id           interface{} // Id
	Cookies      interface{} // cookies
	CountryId    interface{} // 国家id
	CountryCode  interface{} // 国家代码
	Url          interface{} // 请求链接
	UpdatedAt    *gtime.Time // 更新时间
	CreatedAt    *gtime.Time // 创建时间
	Headers      interface{} // 信息头
	StrTimestamp interface{} // 时间戳
	Account      interface{} // 账号
	DeviceId     interface{} // 设备id
}
