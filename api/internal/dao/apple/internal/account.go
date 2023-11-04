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

// AccountDao is the data access object for table apple_account.
type AccountDao struct {
	table      string           // table is the underlying table name of the DAO.
	group      string           // group is the database configuration group name of current DAO.
	columns    AccountColumns   // columns contains all the column names of Table for convenient usage.
	primaryKey string           // 主键ID
	columnArr  []string         // 所有字段的数组
	columnArrG *garray.StrArray // 所有字段的数组（该格式更方便使用）
}

// AccountColumns defines and stores column names for table apple_account.
type AccountColumns struct {
	Id           string //
	Account      string // 账号
	Pwd          string // 密码
	CountryId    string // 国家id
	CountryCode  string // 国家代码
	Balance      string // 余额
	Status       string // 禁用：0否 1是
	Info         string // 信息
	Cookies      string // cookies
	LoginStatus  string // 登录：0否 1是
	IsStop       string // 停用：0否 1是
	UpdatedAt    string // 更新时间
	CreatedAt    string // 创建时间
	Stk          string // x-aos-stk
	DeviceId     string // 设备id
	StrTimestamp string // 时间戳
}

// accountColumns holds the columns for table apple_account.
var accountColumns = AccountColumns{
	Id:           "id",
	Account:      "account",
	Pwd:          "pwd",
	CountryId:    "country_id",
	CountryCode:  "country_code",
	Balance:      "balance",
	Status:       "status",
	Info:         "info",
	Cookies:      "cookies",
	LoginStatus:  "login_status",
	IsStop:       "isStop",
	UpdatedAt:    "updatedAt",
	CreatedAt:    "createdAt",
	Stk:          "stk",
	DeviceId:     "device_id",
	StrTimestamp: "str_timestamp",
}

// NewAccountDao creates and returns a new DAO object for table data access.
func NewAccountDao() *AccountDao {
	return &AccountDao{
		group:   `default`,
		table:   `apple_account`,
		columns: accountColumns,
		primaryKey: func() string {
			return reflect.ValueOf(accountColumns).Field(0).String()
		}(),
		columnArr: func() []string {
			v := reflect.ValueOf(accountColumns)
			count := v.NumField()
			column := make([]string, count)
			for i := 0; i < count; i++ {
				column[i] = v.Field(i).String()
			}
			return column
		}(),
		columnArrG: func() *garray.StrArray {
			v := reflect.ValueOf(accountColumns)
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
func (dao *AccountDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AccountDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
// 使用较为频繁。为优化内存考虑，改成返回指针更为合适，但切忌使用过程中不可修改，否则会污染全局
func (dao *AccountDao) Columns() *AccountColumns {
	return &dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AccountDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AccountDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AccountDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// 主键ID
func (dao *AccountDao) PrimaryKey() string {
	return dao.primaryKey
}

// 所有字段的数组
func (dao *AccountDao) ColumnArr() []string {
	return dao.columnArr
}

// 所有字段的数组（该格式更方便使用）
func (dao *AccountDao) ColumnArrG() *garray.StrArray {
	return dao.columnArrG
}
