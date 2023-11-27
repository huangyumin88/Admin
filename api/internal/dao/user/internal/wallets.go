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

// WalletsDao is the data access object for table app_card_wallets.
type WalletsDao struct {
	table      string           // table is the underlying table name of the DAO.
	group      string           // group is the database configuration group name of current DAO.
	columns    WalletsColumns   // columns contains all the column names of Table for convenient usage.
	primaryKey string           // 主键ID
	columnArr  []string         // 所有字段的数组
	columnArrG *garray.StrArray // 所有字段的数组（该格式更方便使用）
}

// WalletsColumns defines and stores column names for table app_card_wallets.
type WalletsColumns struct {
	WalletId     string // 钱包ID
	UserId       string // 用户ID
	Balance      string // 余额
	RewardPoints string // 积分
	Currency     string // 货币 默认尼日利亚 奈拉
	IsStop       string // 停用：0否 1是
	UpdatedAt    string // 更新时间
	CreatedAt    string // 创建时间
}

// walletsColumns holds the columns for table app_card_wallets.
var walletsColumns = WalletsColumns{
	WalletId:     "wallet_id",
	UserId:       "user_id",
	Balance:      "balance",
	RewardPoints: "reward_points",
	Currency:     "currency",
	IsStop:       "isStop",
	UpdatedAt:    "updatedAt",
	CreatedAt:    "createdAt",
}

// NewWalletsDao creates and returns a new DAO object for table data access.
func NewWalletsDao() *WalletsDao {
	return &WalletsDao{
		group:   `default`,
		table:   `app_card_wallets`,
		columns: walletsColumns,
		primaryKey: func() string {
			return reflect.ValueOf(walletsColumns).Field(0).String()
		}(),
		columnArr: func() []string {
			v := reflect.ValueOf(walletsColumns)
			count := v.NumField()
			column := make([]string, count)
			for i := 0; i < count; i++ {
				column[i] = v.Field(i).String()
			}
			return column
		}(),
		columnArrG: func() *garray.StrArray {
			v := reflect.ValueOf(walletsColumns)
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
func (dao *WalletsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *WalletsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
// 使用较为频繁。为优化内存考虑，改成返回指针更为合适，但切忌使用过程中不可修改，否则会污染全局
func (dao *WalletsDao) Columns() *WalletsColumns {
	return &dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *WalletsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *WalletsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *WalletsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// 主键ID
func (dao *WalletsDao) PrimaryKey() string {
	return dao.primaryKey
}

// 所有字段的数组
func (dao *WalletsDao) ColumnArr() []string {
	return dao.columnArr
}

// 所有字段的数组（该格式更方便使用）
func (dao *WalletsDao) ColumnArrG() *garray.StrArray {
	return dao.columnArrG
}
