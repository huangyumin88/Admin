package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

/*--------列表 开始--------*/
type OrdersListReq struct {
	g.Meta `path:"/orders/list" method:"post" tags:"平台后台/订单" sm:"列表"`
	Filter OrdersListFilter `json:"filter" dc:"查询条件"`
	Field  []string         `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回全部字段。注意：如前端页面所需字段较少，建议传指定字段，可大幅减轻服务器及数据库压力"`
	Sort   string           `json:"sort" default:"id DESC" dc:"排序"`
	Page   int              `json:"page" v:"min:1" default:"1" dc:"页码"`
	Limit  int              `json:"limit" v:"min:0" default:"10" dc:"每页数量。可传0取全部"`
}

type OrdersListFilter struct {
	Id             *uint       `json:"id,omitempty" v:"min:1" dc:"ID"`
	IdArr          []uint      `json:"idArr,omitempty" v:"distinct|foreach|min:1" dc:"ID数组"`
	ExcId          *uint       `json:"excId,omitempty" v:"min:1" dc:"排除ID"`
	ExcIdArr       []uint      `json:"excIdArr,omitempty" v:"distinct|foreach|min:1" dc:"排除ID数组"`
	OrderId        *uint       `json:"order_id,omitempty" v:"min:1" dc:"订单ID"`
	OrderNo        string      `json:"order_no,omitempty" v:"max-length:255" dc:"订单单号"`
	UserId         *uint       `json:"user_id,omitempty" v:"min:1" dc:"用户ID"`
	SalespersonId  *uint       `json:"salesperson_id,omitempty" v:"min:1" dc:"业务员ID"`
	ClientStatus   string      `json:"client_status,omitempty" v:"" dc:"用户订单状态：Pending - 交易中; Failed - 交易失败; Completed - 交易完成; Closed - 关闭;"`
	BackendStatus  string      `json:"backend_status,omitempty" v:"" dc:"后台订单状态：Pending - 等待审核; Loading - 加载中;  Failed - 加载失败; Pledging - 质押中; Completed - 交易完成;"`
	CardCateSubId  *uint       `json:"card_cate_sub_id,omitempty" v:"min:1" dc:"子分类ID"`
	Device         string      `json:"device,omitempty" v:"max-length:30" dc:"使用设备"`
	Wallet         string      `json:"wallet,omitempty" v:"max-length:10" dc:"结算货币"`
	TimeRangeStart *gtime.Time `json:"timeRangeStart,omitempty" v:"date-format:Y-m-d H:i:s" dc:"开始时间：YYYY-mm-dd HH:ii:ss"`
	TimeRangeEnd   *gtime.Time `json:"timeRangeEnd,omitempty" v:"date-format:Y-m-d H:i:s|after-equal:TimeRangeStart" dc:"结束时间：YYYY-mm-dd HH:ii:ss"`
}

type OrdersListRes struct {
	Count int              `json:"count" dc:"总数"`
	List  []OrdersListItem `json:"list" dc:"列表"`
}

type OrdersListItem struct {
	Id            *uint       `json:"id,omitempty" dc:"ID"`
	OrderId       *uint       `json:"order_id,omitempty" dc:"订单ID"`
	OrderNo       *string     `json:"order_no,omitempty" dc:"订单单号"`
	UserId        *uint       `json:"user_id,omitempty" dc:"用户ID"`
	SalespersonId *uint       `json:"salesperson_id,omitempty" dc:"业务员ID"`
	ClientStatus  *string     `json:"client_status,omitempty" dc:"用户订单状态：Pending - 交易中; Failed - 交易失败; Completed - 交易完成; Closed - 关闭;"`
	BackendStatus *string     `json:"backend_status,omitempty" dc:"后台订单状态：Pending - 等待审核; Loading - 加载中;  Failed - 加载失败; Pledging - 质押中; Completed - 交易完成;"`
	FailedReason  *string     `json:"failed_reason,omitempty" dc:"拒绝原因"`
	FailedFiles   *string     `json:"failed_files,omitempty" dc:"拒绝图片"`
	TradeAmount   *float64    `json:"trade_amount,omitempty" dc:"交易金额（AUD）"`
	PayableAmount *uint       `json:"payable_amount,omitempty" dc:"需要支付金额"`
	CardCateSubId *uint       `json:"card_cate_sub_id,omitempty" dc:"子分类ID"`
	Device        *string     `json:"device,omitempty" dc:"使用设备"`
	Wallet        *string     `json:"wallet,omitempty" dc:"结算货币"`
	UpdatedAt     *gtime.Time `json:"updatedAt,omitempty" dc:"更新时间"`
	CreatedAt     *gtime.Time `json:"createdAt,omitempty" dc:"创建时间"`
	UserName      *string     `json:"user_name,omitempty" dc:"用户"`
}

/*--------列表 结束--------*/

/*--------详情 开始--------*/
type OrdersInfoReq struct {
	g.Meta `path:"/orders/info" method:"post" tags:"平台后台/订单" sm:"详情"`
	Id     uint     `json:"id" v:"required|min:1" dc:"ID"`
	Field  []string `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回全部字段。注意：如前端页面所需字段较少，建议传指定字段，可大幅减轻服务器及数据库压力"`
}

type OrdersInfoRes struct {
	Info OrdersInfo `json:"info" dc:"详情"`
}

type OrdersInfo struct {
	Id            *uint       `json:"id,omitempty" dc:"ID"`
	OrderId       *uint       `json:"order_id,omitempty" dc:"订单ID"`
	OrderNo       *string     `json:"order_no,omitempty" dc:"订单单号"`
	UserId        *uint       `json:"user_id,omitempty" dc:"用户ID"`
	SalespersonId *uint       `json:"salesperson_id,omitempty" dc:"业务员ID"`
	ClientStatus  *string     `json:"client_status,omitempty" dc:"用户订单状态：Pending - 交易中; Failed - 交易失败; Completed - 交易完成; Closed - 关闭;"`
	BackendStatus *string     `json:"backend_status,omitempty" dc:"后台订单状态：Pending - 等待审核; Loading - 加载中;  Failed - 加载失败; Pledging - 质押中; Completed - 交易完成;"`
	FailedReason  *string     `json:"failed_reason,omitempty" dc:"拒绝原因"`
	FailedFiles   *string     `json:"failed_files,omitempty" dc:"拒绝图片"`
	TradeAmount   *float64    `json:"trade_amount,omitempty" dc:"交易金额（AUD）"`
	PayableAmount *uint       `json:"payable_amount,omitempty" dc:"需要支付金额"`
	CardCateSubId *uint       `json:"card_cate_sub_id,omitempty" dc:"子分类ID"`
	Device        *string     `json:"device,omitempty" dc:"使用设备"`
	Wallet        *string     `json:"wallet,omitempty" dc:"结算货币"`
	UpdatedAt     *gtime.Time `json:"updatedAt,omitempty" dc:"更新时间"`
	CreatedAt     *gtime.Time `json:"createdAt,omitempty" dc:"创建时间"`
}

/*--------详情 结束--------*/

/*--------新增 开始--------*/
type OrdersCreateReq struct {
	g.Meta        `path:"/orders/create" method:"post" tags:"平台后台/订单" sm:"创建"`
	OrderNo       *string  `json:"order_no,omitempty" v:"max-length:255" dc:"订单单号"`
	UserId        *uint    `json:"user_id,omitempty" v:"min:1" dc:"用户ID"`
	SalespersonId *uint    `json:"salesperson_id,omitempty" v:"min:1" dc:"业务员ID"`
	ClientStatus  *string  `json:"client_status,omitempty" v:"" dc:"用户订单状态：Pending - 交易中; Failed - 交易失败; Completed - 交易完成; Closed - 关闭;"`
	BackendStatus *string  `json:"backend_status,omitempty" v:"" dc:"后台订单状态：Pending - 等待审核; Loading - 加载中;  Failed - 加载失败; Pledging - 质押中; Completed - 交易完成;"`
	FailedReason  *string  `json:"failed_reason,omitempty" v:"" dc:"拒绝原因"`
	FailedFiles   *string  `json:"failed_files,omitempty" v:"" dc:"拒绝图片"`
	TradeAmount   *float64 `json:"trade_amount,omitempty" v:"min:0" dc:"交易金额（AUD）"`
	PayableAmount *uint    `json:"payable_amount,omitempty" v:"" dc:"需要支付金额"`
	CardCateSubId *uint    `json:"card_cate_sub_id,omitempty" v:"min:1" dc:"子分类ID"`
	Device        *string  `json:"device,omitempty" v:"max-length:30" dc:"使用设备"`
	Wallet        *string  `json:"wallet,omitempty" v:"max-length:10" dc:"结算货币"`
}

/*--------新增 结束--------*/

/*--------修改 开始--------*/
type OrdersUpdateReq struct {
	g.Meta        `path:"/orders/update" method:"post" tags:"平台后台/订单" sm:"更新"`
	IdArr         []uint   `json:"idArr,omitempty" v:"required|distinct|foreach|min:1" dc:"ID数组"`
	OrderNo       *string  `json:"order_no,omitempty" v:"max-length:255" dc:"订单单号"`
	UserId        *uint    `json:"user_id,omitempty" v:"min:1" dc:"用户ID"`
	SalespersonId *uint    `json:"salesperson_id,omitempty" v:"min:1" dc:"业务员ID"`
	ClientStatus  *string  `json:"client_status,omitempty" v:"" dc:"用户订单状态：Pending - 交易中; Failed - 交易失败; Completed - 交易完成; Closed - 关闭;"`
	BackendStatus *string  `json:"backend_status,omitempty" v:"" dc:"后台订单状态：Pending - 等待审核; Loading - 加载中;  Failed - 加载失败; Pledging - 质押中; Completed - 交易完成;"`
	FailedReason  *string  `json:"failed_reason,omitempty" v:"" dc:"拒绝原因"`
	FailedFiles   *string  `json:"failed_files,omitempty" v:"" dc:"拒绝图片"`
	TradeAmount   *float64 `json:"trade_amount,omitempty" v:"min:0" dc:"交易金额（AUD）"`
	PayableAmount *uint    `json:"payable_amount,omitempty" v:"" dc:"需要支付金额"`
	CardCateSubId *uint    `json:"card_cate_sub_id,omitempty" v:"min:1" dc:"子分类ID"`
	Device        *string  `json:"device,omitempty" v:"max-length:30" dc:"使用设备"`
	Wallet        *string  `json:"wallet,omitempty" v:"max-length:10" dc:"结算货币"`
}

/*--------修改 结束--------*/

/*--------删除 开始--------*/
type OrdersDeleteReq struct {
	g.Meta `path:"/orders/del" method:"post" tags:"平台后台/订单" sm:"删除"`
	IdArr  []uint `json:"idArr,omitempty" v:"required|distinct|foreach|min:1" dc:"ID数组"`
}

/*--------删除 结束--------*/

/*--------后台订单状态查询 开始--------*/

//type OrdersQueryOrderStatusReq struct {
//	g.Meta    `path:"/orders/checkorderstatus" method:"post" tags:"平台后台/订单" sm:"查询订单状态"`
//	OrderType int  `json:"order_type" v:"min:1" default:"0" dc:"查看订单类型的状态， 0 - 操作员订单状态; 1 - 用户订单状态"`
//	OrderId   uint `json:"order_id" v:"min:1" default:"0" dc:"订单ID  可以不传"`
//}
//
//type OrdersQueryOrderStatusRes struct {
//	List []OrdersQueryOrderStatusItem `json:"list" dc:"列表"`
//}
//
//type OrdersQueryOrderStatusItem struct {
//	status *string `json:"status,omitempty" dc:"用户订单状态：Pending - 交易中; Failed - 交易失败; Completed - 交易完成; Closed - 关闭;"`
//	Label  *string `json:"label,omitempty" dc:"标签。常用于前端组件"`
//	//BackendStatus *string     `json:"backend_status,omitempty" dc:"后台订单状态：Pending - 等待审核; Loading - 加载中;  Failed - 加载失败; Pledging - 质押中; Completed - 交易完成;"`
//}

/*--------后台订单状态查询 结束--------*/
