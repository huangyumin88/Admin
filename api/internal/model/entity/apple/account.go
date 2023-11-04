// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Account is the golang structure for table account.
type Account struct {
	Id          int         `json:"id"          ` //
	Account     string      `json:"account"     ` // 账号
	Pwd         string      `json:"pwd"         ` // 密码
	CountryId   int         `json:"countryId"   ` // 国家id
	Balance     string      `json:"balance"     ` // 余额
	Status      int         `json:"status"      ` // 禁用：0否 1是
	Info        string      `json:"info"        ` // 信息
	Cookies     string      `json:"cookies"     ` // cookies
	LoginStatus int         `json:"loginStatus" ` // 登录：0否 1是
	IsStop      uint        `json:"isStop"      ` // 停用：0否 1是
	UpdatedAt   *gtime.Time `json:"updatedAt"   ` // 更新时间
	CreatedAt   *gtime.Time `json:"createdAt"   ` // 创建时间
	Stk         string      `json:"stk"         ` // x-aos-stk
}
