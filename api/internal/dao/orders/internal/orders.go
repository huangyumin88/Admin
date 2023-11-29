// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"reflect"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OrdersDao is the data access object for table app_card_orders.
type OrdersDao struct {
	table      string           // table is the underlying table name of the DAO.
	group      string           // group is the database configuration group name of current DAO.
	columns    OrdersColumns    // columns contains all the column names of Table for convenient usage.
	primaryKey string           // 主键ID
	columnArr  []string         // 所有字段的数组
	columnArrG *garray.StrArray // 所有字段的数组（该格式更方便使用）
}

// OrdersColumns defines and stores column names for table app_card_orders.
type OrdersColumns struct {
	OrderId       string // 订单ID
	OrderNo       string // 订单单号
	UserId        string // 用户ID
	SalespersonId string // 业务员ID
	ClientStatus  string // 用户订单状态：Pending - 交易中; Failed - 交易失败; Completed - 交易完成; Closed - 关闭;
	BackendStatus string // 后台订单状态：Pending - 等待审核; Loading - 加载中;  Failed - 加载失败; Pledging - 质押中; Completed - 交易完成;
	TradeFiles	  string // 交易图片
	FailedReason  string // 拒绝原因
	FailedFiles   string // 拒绝图片
	TradeAmount   string // 交易金额（AUD）
	PayableAmount string // 需要支付金额
	CardCateSubId string // 子分类ID
	Device        string // 使用设备
	Wallet        string // 结算货币
	UpdatedAt     string // 更新时间
	CreatedAt     string // 创建时间
}

// ordersColumns holds the columns for table app_card_orders.
var ordersColumns = OrdersColumns{
	OrderId:       "order_id",
	OrderNo:       "order_no",
	UserId:        "user_id",
	SalespersonId: "salesperson_id",
	ClientStatus:  "client_status",
	BackendStatus: "backend_status",
	TradeFiles:	   "trade_files",
	FailedReason:  "failed_reason",
	FailedFiles:   "failed_files",
	TradeAmount:   "trade_amount",
	PayableAmount: "payable_amount",
	CardCateSubId: "card_cate_sub_id",
	Device:        "device",
	Wallet:        "wallet",
	UpdatedAt:     "updatedAt",
	CreatedAt:     "createdAt",
}

// NewOrdersDao creates and returns a new DAO object for table data access.
func NewOrdersDao() *OrdersDao {
	return &OrdersDao{
		group:   `default`,
		table:   `app_card_orders`,
		columns: ordersColumns,
		primaryKey: func() string {
			return reflect.ValueOf(ordersColumns).Field(0).String()
		}(),
		columnArr: func() []string {
			v := reflect.ValueOf(ordersColumns)
			count := v.NumField()
			column := make([]string, count)
			for i := 0; i < count; i++ {
				column[i] = v.Field(i).String()
			}
			return column
		}(),
		columnArrG: func() *garray.StrArray {
			v := reflect.ValueOf(ordersColumns)
			count := v.NumField()
			column := make([]string, count)
			for i := 0; i < count; i++ {
				column[i] = v.Field(i).String()
			}
			return garray.NewStrArrayFrom(column)
		}(),
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *OrdersDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *OrdersDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
// 使用较为频繁。为优化内存考虑，改成返回指针更为合适，但切忌使用过程中不可修改，否则会污染全局
func (dao *OrdersDao) Columns() *OrdersColumns {
	return &dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *OrdersDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *OrdersDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *OrdersDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// 主键ID
func (dao *OrdersDao) PrimaryKey() string {
	return dao.primaryKey
}

// 所有字段的数组
func (dao *OrdersDao) ColumnArr() []string {
	return dao.columnArr
}

// 所有字段的数组（该格式更方便使用）
func (dao *OrdersDao) ColumnArrG() *garray.StrArray {
	return dao.columnArrG
}
