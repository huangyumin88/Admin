// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CardCategories is the golang structure of table app_card_categories for DAO operations like Where/Data.
type CardCategories struct {
	g.Meta    `orm:"table:app_card_categories, do:true"`
	Id        interface{} // ID
	SubId     interface{} // Card id
	Avatar    interface{} // 图片
	AvatarUrl interface{} // 转发图片地址
	Name      interface{} // card 名字
	Sort      interface{} // 排序值。从小到大排序，默认50，范围0-100
	IsActive  interface{} // 活动：0否 1是
	IsStop    interface{} // 停用：0否 1是
	UpdatedAt *gtime.Time // 更新时间
	CreatedAt *gtime.Time // 创建时间
}
