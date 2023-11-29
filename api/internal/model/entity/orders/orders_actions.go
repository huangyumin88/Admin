// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// OrdersActions is the golang structure for table orders_actions.
type OrdersActions struct {
	Id            int         `json:"id"            ` // ID
	ActionsUserId uint        `json:"actionsUserId" ` // 操作人员
	OrderId       uint        `json:"orderId"       ` // 订单ID
	BackendStatus string      `json:"backendStatus" ` // 操作状态
	Remarks       string      `json:"remarks"       ` // 备注
	UpdatedAt     *gtime.Time `json:"updatedAt"     ` // 更新时间
	CreatedAt     *gtime.Time `json:"createdAt"     ` // 创建时间
}
