// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CardUrl is the golang structure of table apple_card_url for DAO operations like Where/Data.
type CardUrl struct {
	g.Meta      `orm:"table:apple_card_url, do:true"`
	Id          interface{} //
	Url         interface{} // 请求链接
	AccountId   interface{} // 苹果账号ID
	CountryCode interface{} // 国家代码
	IsStop      interface{} // 停用：0否 1是
	IsAutoLogin interface{} // 自动登录：0否 1是
	UpdatedAt   *gtime.Time // 更新时间
	CreatedAt   *gtime.Time // 创建时间
}
