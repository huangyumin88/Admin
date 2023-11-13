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

// CardUrlDao is the data access object for table apple_card_url.
type CardUrlDao struct {
	table      string           // table is the underlying table name of the DAO.
	group      string           // group is the database configuration group name of current DAO.
	columns    CardUrlColumns   // columns contains all the column names of Table for convenient usage.
	primaryKey string           // 主键ID
	columnArr  []string         // 所有字段的数组
	columnArrG *garray.StrArray // 所有字段的数组（该格式更方便使用）
}

// CardUrlColumns defines and stores column names for table apple_card_url.
type CardUrlColumns struct {
	Id          string //
	Url         string // 请求链接
	AccountId   string // 苹果账号ID
	CountryCode string // 国家代码
	IsStop      string // 停用：0否 1是
	IsAutoLogin string // 自动登录：0否 1是
	UpdatedAt   string // 更新时间
	CreatedAt   string // 创建时间
}

// cardUrlColumns holds the columns for table apple_card_url.
var cardUrlColumns = CardUrlColumns{
	Id:          "id",
	Url:         "url",
	AccountId:   "account_id",
	CountryCode: "country_code",
	IsStop:      "isStop",
	IsAutoLogin: "isAutoLogin",
	UpdatedAt:   "updatedAt",
	CreatedAt:   "createdAt",
}

// NewCardUrlDao creates and returns a new DAO object for table data access.
func NewCardUrlDao() *CardUrlDao {
	return &CardUrlDao{
		group:   `default`,
		table:   `apple_card_url`,
		columns: cardUrlColumns,
		primaryKey: func() string {
			return reflect.ValueOf(cardUrlColumns).Field(0).String()
		}(),
		columnArr: func() []string {
			v := reflect.ValueOf(cardUrlColumns)
			count := v.NumField()
			column := make([]string, count)
			for i := 0; i < count; i++ {
				column[i] = v.Field(i).String()
			}
			return column
		}(),
		columnArrG: func() *garray.StrArray {
			v := reflect.ValueOf(cardUrlColumns)
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
func (dao *CardUrlDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CardUrlDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
// 使用较为频繁。为优化内存考虑，改成返回指针更为合适，但切忌使用过程中不可修改，否则会污染全局
func (dao *CardUrlDao) Columns() *CardUrlColumns {
	return &dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CardUrlDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CardUrlDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CardUrlDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// 主键ID
func (dao *CardUrlDao) PrimaryKey() string {
	return dao.primaryKey
}

// 所有字段的数组
func (dao *CardUrlDao) ColumnArr() []string {
	return dao.columnArr
}

// 所有字段的数组（该格式更方便使用）
func (dao *CardUrlDao) ColumnArrG() *garray.StrArray {
	return dao.columnArrG
}
