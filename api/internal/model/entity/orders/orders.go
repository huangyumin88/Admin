// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Orders is the golang structure for table orders.
type Orders struct {
	OrderId       uint        `json:"orderId"       ` // 订单ID
	OrderNo       string      `json:"orderNo"       ` // 订单单号
	UserId        uint        `json:"userId"        ` // 用户ID
	SalespersonId uint        `json:"salespersonId" ` // 业务员ID
	ClientStatus  string      `json:"clientStatus"  ` // 用户订单状态：Pending - 交易中; Failed - 交易失败; Completed - 交易完成; Closed - 关闭;
	BackendStatus string      `json:"backendStatus" ` // 后台订单状态：Pending - 等待审核; Loading - 加载中;  Failed - 加载失败; Pledging - 质押中; Completed - 交易完成;
	TradeFiles    string      `json:"tradeFiles"    ` // 交易图片
	FailedReason  string      `json:"failedReason"  ` // 拒绝原因
	FailedFiles   string      `json:"failedFiles"   ` // 拒绝图片
	TradeAmount   float64     `json:"tradeAmount"   ` // 交易金额（AUD）
	PayableAmount uint        `json:"payableAmount" ` // 需要支付金额
	CardCateId    uint        `json:"cardCateId"    ` // 主分类ID
	CardCateSubId uint        `json:"cardCateSubId" ` // 子分类ID
	Device        string      `json:"device"        ` // 使用设备
	Wallet        string      `json:"wallet"        ` // 结算货币
	UpdatedAt     *gtime.Time `json:"updatedAt"     ` // 更新时间
	CreatedAt     *gtime.Time `json:"createdAt"     ` // 创建时间
}
