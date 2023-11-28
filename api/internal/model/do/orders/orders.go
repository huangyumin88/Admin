// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Orders is the golang structure of table app_card_orders for DAO operations like Where/Data.
type Orders struct {
	g.Meta        `orm:"table:app_card_orders, do:true"`
	OrderId       interface{} // 订单ID
	OrderNo       interface{} // 订单单号
	UserId        interface{} // 用户ID
	SalespersonId interface{} // 业务员ID
	ClientStatus  interface{} // 用户订单状态：Pending - 交易中; Failed - 交易失败; Completed - 交易完成; Closed - 关闭;
	BackendStatus interface{} // 后台订单状态：Pending - 等待审核; Loading - 加载中;  Failed - 加载失败; Pledging - 质押中; Completed - 交易完成;
	FailedReason  interface{} // 拒绝原因
	FailedFiles   interface{} // 拒绝图片
	TradeAmount   interface{} // 交易金额（AUD）
	PayableAmount interface{} // 需要支付金额
	CardCateSubId interface{} // 子分类ID
	Device        interface{} // 使用设备
	Wallet        interface{} // 结算货币
	UpdatedAt     *gtime.Time // 更新时间
	CreatedAt     *gtime.Time // 创建时间
}
