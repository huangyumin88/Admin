// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CardCategoriesSub is the golang structure for table card_categories_sub.
type CardCategoriesSub struct {
	Id                  int         `json:"id"                  ` //
	CateId              int         `json:"cateId"              ` // 主分类ID
	Name                string      `json:"name"                ` // card 名字
	Rate                int         `json:"rate"                ` // 费率
	MinAcceptableAmount int         `json:"minAcceptableAmount" ` // 最低可接受金额
	Announcements       string      `json:"announcements"       ` // 注意事项
	Sort                uint        `json:"sort"                ` // 排序值。从小到大排序，默认50，范围0-100
	IsActive            int         `json:"isActive"            ` // 活动：0否 1是
	IsStop              uint        `json:"isStop"              ` // 停用：0否 1是
	SubId               string      `json:"subId"               ` // Card id
	UpdatedAt           *gtime.Time `json:"updatedAt"           ` // 更新时间
	CreatedAt           *gtime.Time `json:"createdAt"           ` // 创建时间
}
