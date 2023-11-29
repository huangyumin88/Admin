// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// OrdersActions is the golang structure of table app_card_orders_actions for DAO operations like Where/Data.
type OrdersActions struct {
	g.Meta        `orm:"table:app_card_orders_actions, do:true"`
	Id            interface{} // ID
	ActionsUserId interface{} // 操作人员
	OrderId       interface{} // 订单ID
	BackendStatus interface{} // 操作状态
	Remarks       interface{} // 备注
	UpdatedAt     *gtime.Time // 更新时间
	CreatedAt     *gtime.Time // 创建时间
}
