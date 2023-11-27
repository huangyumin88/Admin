// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CardCategoriesSub is the golang structure of table app_card_categories_sub for DAO operations like Where/Data.
type CardCategoriesSub struct {
	g.Meta              `orm:"table:app_card_categories_sub, do:true"`
	Id                  interface{} //
	CateId              interface{} // 主分类ID
	Name                interface{} // card 名字
	Rate                interface{} // 费率
	MinAcceptableAmount interface{} // 最低可接受金额
	Announcements       interface{} // 注意事项
	Sort                interface{} // 排序值。从小到大排序，默认50，范围0-100
	IsActive            interface{} // 活动：0否 1是
	IsStop              interface{} // 停用：0否 1是
	SubId               interface{} // Card id
	UpdatedAt           *gtime.Time // 更新时间
	CreatedAt           *gtime.Time // 创建时间
}
