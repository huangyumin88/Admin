// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CardCategories is the golang structure for table card_categories.
type CardCategories struct {
	Id        int         `json:"id"        ` // ID
	SubId     string      `json:"subId"     ` // Card id
	Avatar    string      `json:"avatar"    ` // 图片
	AvatarUrl string      `json:"avatarUrl" ` // 转发图片地址
	Name      string      `json:"name"      ` // card 名字
	Sort      uint        `json:"sort"      ` // 排序值。从小到大排序，默认50，范围0-100
	IsActive  int         `json:"isActive"  ` // 活动：0否 1是
	IsStop    uint        `json:"isStop"    ` // 停用：0否 1是
	UpdatedAt *gtime.Time `json:"updatedAt" ` // 更新时间
	CreatedAt *gtime.Time `json:"createdAt" ` // 创建时间
}
