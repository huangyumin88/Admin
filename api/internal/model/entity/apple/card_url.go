// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CardUrl is the golang structure for table card_url.
type CardUrl struct {
	Id          int         `json:"id"          ` //
	Url         string      `json:"url"         ` // 请求链接
	AccountId   int         `json:"accountId"   ` // 苹果账号ID
	CountryCode string      `json:"countryCode" ` // 国家代码
	IsStop      uint        `json:"isStop"      ` // 停用：0否 1是
	IsAutoLogin int         `json:"isAutoLogin" ` // 自动登录：0否 1是
	UpdatedAt   *gtime.Time `json:"updatedAt"   ` // 更新时间
	CreatedAt   *gtime.Time `json:"createdAt"   ` // 创建时间
}
