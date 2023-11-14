// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Cookies is the golang structure for table cookies.
type Cookies struct {
	Id           int         `json:"id"           ` // Id
	Cookies      string      `json:"cookies"      ` // cookies
	CountryId    int         `json:"countryId"    ` // 国家id
	CountryCode  string      `json:"countryCode"  ` // 国家代码
	UpdatedAt    *gtime.Time `json:"updatedAt"    ` // 更新时间
	CreatedAt    *gtime.Time `json:"createdAt"    ` // 创建时间
	Headers      string      `json:"headers"      ` // 信息头
	StrTimestamp string      `json:"strTimestamp" ` // 时间戳
	Account      string      `json:"account"      ` // 账号
	DeviceId     string      `json:"deviceId"     ` // 设备id
}
