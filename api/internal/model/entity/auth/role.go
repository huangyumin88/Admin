// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Role is the golang structure for table role.
type Role struct {
	RoleId    uint        `json:"roleId"    ` // 权限角色ID
	SceneId   uint        `json:"sceneId"   ` // 权限场景ID
	TableId   uint        `json:"tableId"   ` // 关联表ID（0表示平台创建，其他值根据authSceneId对应不同表，表示是哪个表内哪个机构或个人创建）
	RoleName  string      `json:"roleName"  ` // 名称
	IsStop    uint        `json:"isStop"    ` // 停用：0否 1是
	UpdatedAt *gtime.Time `json:"updatedAt" ` // 更新时间
	CreatedAt *gtime.Time `json:"createdAt" ` // 创建时间
}
