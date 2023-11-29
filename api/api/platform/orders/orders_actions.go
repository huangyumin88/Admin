package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

/*--------列表 开始--------*/
type OrdersActionsListReq struct {
	g.Meta `path:"/ordersActions/list" method:"post" tags:"平台后台/订单操作记录" sm:"列表"`
	Filter OrdersActionsListFilter `json:"filter" dc:"查询条件"`
	Field  []string                `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回全部字段。注意：如前端页面所需字段较少，建议传指定字段，可大幅减轻服务器及数据库压力"`
	Sort   string                  `json:"sort" default:"id DESC" dc:"排序"`
	Page   int                     `json:"page" v:"min:1" default:"1" dc:"页码"`
	Limit  int                     `json:"limit" v:"min:0" default:"10" dc:"每页数量。可传0取全部"`
}

type OrdersActionsListFilter struct {
	Id             *uint       `json:"id,omitempty" v:"min:1" dc:"ID"`
	IdArr          []uint      `json:"idArr,omitempty" v:"distinct|foreach|min:1" dc:"ID数组"`
	ExcId          *uint       `json:"excId,omitempty" v:"min:1" dc:"排除ID"`
	ExcIdArr       []uint      `json:"excIdArr,omitempty" v:"distinct|foreach|min:1" dc:"排除ID数组"`
	ActionsUserId  *uint       `json:"actions_user_id,omitempty" v:"min:1" dc:"操作人员"`
	OrderId        *uint       `json:"order_id,omitempty" v:"min:1" dc:"订单ID"`
	BackendStatus  string      `json:"backend_status,omitempty" v:"max-length:10" dc:"操作状态"`
	Remarks        string      `json:"remarks,omitempty" v:"max-length:255" dc:"备注"`
	TimeRangeStart *gtime.Time `json:"timeRangeStart,omitempty" v:"date-format:Y-m-d H:i:s" dc:"开始时间：YYYY-mm-dd HH:ii:ss"`
	TimeRangeEnd   *gtime.Time `json:"timeRangeEnd,omitempty" v:"date-format:Y-m-d H:i:s|after-equal:TimeRangeStart" dc:"结束时间：YYYY-mm-dd HH:ii:ss"`
}

type OrdersActionsListRes struct {
	Count int                     `json:"count" dc:"总数"`
	List  []OrdersActionsListItem `json:"list" dc:"列表"`
}

type OrdersActionsListItem struct {
	Id            *uint       `json:"id,omitempty" dc:"ID"`
	ActionsUserId *uint       `json:"actions_user_id,omitempty" dc:"操作人员"`
	OrderId       *uint       `json:"order_id,omitempty" dc:"订单ID"`
	BackendStatus *string     `json:"backend_status,omitempty" dc:"操作状态"`
	Remarks       *string     `json:"remarks,omitempty" dc:"备注"`
	UpdatedAt     *gtime.Time `json:"updatedAt,omitempty" dc:"更新时间"`
	CreatedAt     *gtime.Time `json:"createdAt,omitempty" dc:"创建时间"`
}

/*--------列表 结束--------*/

/*--------详情 开始--------*/
type OrdersActionsInfoReq struct {
	g.Meta `path:"/ordersActions/info" method:"post" tags:"平台后台/订单操作记录" sm:"详情"`
	Id     uint     `json:"id" v:"required|min:1" dc:"ID"`
	Field  []string `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回全部字段。注意：如前端页面所需字段较少，建议传指定字段，可大幅减轻服务器及数据库压力"`
}

type OrdersActionsInfoRes struct {
	Info OrdersActionsInfo `json:"info" dc:"详情"`
}

type OrdersActionsInfo struct {
	Id            *uint       `json:"id,omitempty" dc:"ID"`
	ActionsUserId *uint       `json:"actions_user_id,omitempty" dc:"操作人员"`
	OrderId       *uint       `json:"order_id,omitempty" dc:"订单ID"`
	BackendStatus *string     `json:"backend_status,omitempty" dc:"操作状态"`
	Remarks       *string     `json:"remarks,omitempty" dc:"备注"`
	UpdatedAt     *gtime.Time `json:"updatedAt,omitempty" dc:"更新时间"`
	CreatedAt     *gtime.Time `json:"createdAt,omitempty" dc:"创建时间"`
}

/*--------详情 结束--------*/

/*--------新增 开始--------*/
type OrdersActionsCreateReq struct {
	g.Meta        `path:"/ordersActions/create" method:"post" tags:"平台后台/订单操作记录" sm:"创建"`
	Id            *int    `json:"id,omitempty" v:"min:1" dc:"ID"`
	ActionsUserId *uint   `json:"actions_user_id,omitempty" v:"min:1" dc:"操作人员"`
	OrderId       *uint   `json:"order_id,omitempty" v:"min:1" dc:"订单ID"`
	BackendStatus *string `json:"backend_status,omitempty" v:"max-length:10" dc:"操作状态"`
	Remarks       *string `json:"remarks,omitempty" v:"max-length:255" dc:"备注"`
}

/*--------新增 结束--------*/

/*--------修改 开始--------*/
type OrdersActionsUpdateReq struct {
	g.Meta        `path:"/ordersActions/update" method:"post" tags:"平台后台/订单操作记录" sm:"更新"`
	IdArr         []uint  `json:"idArr,omitempty" v:"required|distinct|foreach|min:1" dc:"ID数组"`
	Id            *int    `json:"id,omitempty" v:"min:1" dc:"ID"`
	ActionsUserId *uint   `json:"actions_user_id,omitempty" v:"min:1" dc:"操作人员"`
	OrderId       *uint   `json:"order_id,omitempty" v:"min:1" dc:"订单ID"`
	BackendStatus *string `json:"backend_status,omitempty" v:"max-length:10" dc:"操作状态"`
	Remarks       *string `json:"remarks,omitempty" v:"max-length:255" dc:"备注"`
}

/*--------修改 结束--------*/

/*--------删除 开始--------*/
type OrdersActionsDeleteReq struct {
	g.Meta `path:"/ordersActions/del" method:"post" tags:"平台后台/订单操作记录" sm:"删除"`
	IdArr  []uint `json:"idArr,omitempty" v:"required|distinct|foreach|min:1" dc:"ID数组"`
}

/*--------删除 结束--------*/
