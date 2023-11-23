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

// CardCountriesDao is the data access object for table app_card_countries.
type CardCountriesDao struct {
	table      string               // table is the underlying table name of the DAO.
	group      string               // group is the database configuration group name of current DAO.
	columns    CardCountriesColumns // columns contains all the column names of Table for convenient usage.
	primaryKey string               // 主键ID
	columnArr  []string             // 所有字段的数组
	columnArrG *garray.StrArray     // 所有字段的数组（该格式更方便使用）
}

// CardCountriesColumns defines and stores column names for table app_card_countries.
type CardCountriesColumns struct {
	Id           string //
	Name         string //
	IsoName      string //
	CurrencyCode string // 国家代码
	CurrencyName string // 国家名称
	FlagUrl      string //
	FlagAvatar   string // 国家图片
	FlagAvatarID string // 图片
	IsStop       string // 停用：0否 1是
	UpdatedAt    string // 更新时间
	CreatedAt    string // 创建时间
}

// cardCountriesColumns holds the columns for table app_card_countries.
var cardCountriesColumns = CardCountriesColumns{
	Id:           "id",
	Name:         "name",
	IsoName:      "isoName",
	CurrencyCode: "currencyCode",
	CurrencyName: "currencyName",
	FlagUrl:      "flagUrl",
	FlagAvatar:   "flagAvatar",
	FlagAvatarID: "flagAvatarID",
	IsStop:       "isStop",
	UpdatedAt:    "updatedAt",
	CreatedAt:    "createdAt",
}

// NewCardCountriesDao creates and returns a new DAO object for table data access.
func NewCardCountriesDao() *CardCountriesDao {
	return &CardCountriesDao{
		group:   `default`,
		table:   `app_card_countries`,
		columns: cardCountriesColumns,
		primaryKey: func() string {
			return reflect.ValueOf(cardCountriesColumns).Field(0).String()
		}(),
		columnArr: func() []string {
			v := reflect.ValueOf(cardCountriesColumns)
			count := v.NumField()
			column := make([]string, count)
			for i := 0; i < count; i++ {
				column[i] = v.Field(i).String()
			}
			return column
		}(),
		columnArrG: func() *garray.StrArray {
			v := reflect.ValueOf(cardCountriesColumns)
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
func (dao *CardCountriesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CardCountriesDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
// 使用较为频繁。为优化内存考虑，改成返回指针更为合适，但切忌使用过程中不可修改，否则会污染全局
func (dao *CardCountriesDao) Columns() *CardCountriesColumns {
	return &dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CardCountriesDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CardCountriesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CardCountriesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// 主键ID
func (dao *CardCountriesDao) PrimaryKey() string {
	return dao.primaryKey
}

// 所有字段的数组
func (dao *CardCountriesDao) ColumnArr() []string {
	return dao.columnArr
}

// 所有字段的数组（该格式更方便使用）
func (dao *CardCountriesDao) ColumnArrG() *garray.StrArray {
	return dao.columnArrG
}
