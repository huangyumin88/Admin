// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Image is the golang structure for table image.
type Image struct {
	Id        uint        `json:"id"        ` //
	Url       string      `json:"url"       ` //
	Avatar    string      `json:"avatar"    ` // 头像
	IsStop    uint        `json:"isStop"    ` // 停用：0否 1是
	UpdatedAt *gtime.Time `json:"updatedAt" ` // 更新时间
	CreatedAt *gtime.Time `json:"createdAt" ` // 创建时间
}
