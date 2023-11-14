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

// CookiesDao is the data access object for table apple_cookies.
type CookiesDao struct {
	table      string           // table is the underlying table name of the DAO.
	group      string           // group is the database configuration group name of current DAO.
	columns    CookiesColumns   // columns contains all the column names of Table for convenient usage.
	primaryKey string           // 主键ID
	columnArr  []string         // 所有字段的数组
	columnArrG *garray.StrArray // 所有字段的数组（该格式更方便使用）
}

// CookiesColumns defines and stores column names for table apple_cookies.
type CookiesColumns struct {
	Id           string // Id
	Cookies      string // cookies
	CountryId    string // 国家id
	CountryCode  string // 国家代码
	UpdatedAt    string // 更新时间
	CreatedAt    string // 创建时间
	Headers      string // 信息头
	StrTimestamp string // 时间戳
	Account      string // 账号
	DeviceId     string // 设备id
}

// cookiesColumns holds the columns for table apple_cookies.
var cookiesColumns = CookiesColumns{
	Id:           "id",
	Cookies:      "cookies",
	CountryId:    "country_id",
	CountryCode:  "country_code",
	UpdatedAt:    "updatedAt",
	CreatedAt:    "createdAt",
	Headers:      "headers",
	StrTimestamp: "str_timestamp",
	Account:      "account",
	DeviceId:     "device_id",
}

// NewCookiesDao creates and returns a new DAO object for table data access.
func NewCookiesDao() *CookiesDao {
	return &CookiesDao{
		group:   `default`,
		table:   `apple_cookies`,
		columns: cookiesColumns,
		primaryKey: func() string {
			return reflect.ValueOf(cookiesColumns).Field(0).String()
		}(),
		columnArr: func() []string {
			v := reflect.ValueOf(cookiesColumns)
			count := v.NumField()
			column := make([]string, count)
			for i := 0; i < count; i++ {
				column[i] = v.Field(i).String()
			}
			return column
		}(),
		columnArrG: func() *garray.StrArray {
			v := reflect.ValueOf(cookiesColumns)
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
func (dao *CookiesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CookiesDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
// 使用较为频繁。为优化内存考虑，改成返回指针更为合适，但切忌使用过程中不可修改，否则会污染全局
func (dao *CookiesDao) Columns() *CookiesColumns {
	return &dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CookiesDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CookiesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CookiesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// 主键ID
func (dao *CookiesDao) PrimaryKey() string {
	return dao.primaryKey
}

// 所有字段的数组
func (dao *CookiesDao) ColumnArr() []string {
	return dao.columnArr
}

// 所有字段的数组（该格式更方便使用）
func (dao *CookiesDao) ColumnArrG() *garray.StrArray {
	return dao.columnArrG
}
