// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CardCountries is the golang structure of table app_card_countries for DAO operations like Where/Data.
type CardCountries struct {
	g.Meta       `orm:"table:app_card_countries, do:true"`
	Id           interface{} //
	Name         interface{} //
	IsoName      interface{} //
	CurrencyCode interface{} // 国家代码
	CurrencyName interface{} // 国家名称
	FlagUrl      interface{} //
	FlagAvatar   interface{} // 国家图片
	FlagAvatarID interface{} // 图片
	IsStop       interface{} // 停用：0否 1是
	UpdatedAt    *gtime.Time // 更新时间
	CreatedAt    *gtime.Time // 创建时间
}
