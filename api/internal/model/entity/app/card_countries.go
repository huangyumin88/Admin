// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CardCountries is the golang structure for table card_countries.
type CardCountries struct {
	Id           int         `json:"id"           ` //
	Name         string      `json:"name"         ` //
	IsoName      string      `json:"isoName"      ` //
	CurrencyCode string      `json:"currencyCode" ` // 国家代码
	CurrencyName string      `json:"currencyName" ` // 国家名称
	FlagUrl      string      `json:"flagUrl"      ` //
	FlagAvatar   string      `json:"flagAvatar"   ` // 国家图片
	FlagAvatarID string      `json:"flagAvatarID" ` // 图片
	IsStop       uint        `json:"isStop"       ` // 停用：0否 1是
	UpdatedAt    *gtime.Time `json:"updatedAt"    ` // 更新时间
	CreatedAt    *gtime.Time `json:"createdAt"    ` // 创建时间
}
