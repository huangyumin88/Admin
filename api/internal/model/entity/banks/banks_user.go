// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BanksUser is the golang structure for table banks_user.
type BanksUser struct {
	Id        int         `json:"id"        ` // ID
	UserId    uint        `json:"userId"    ` // 用户id
	BankId    int         `json:"bankId"    ` // 银行id
	UpdatedAt *gtime.Time `json:"updatedAt" ` // 更新时间
	CreatedAt *gtime.Time `json:"createdAt" ` // 创建时间
}
