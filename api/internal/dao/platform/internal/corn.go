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

// CornDao is the data access object for table platform_corn.
type CornDao struct {
	table      string           // table is the underlying table name of the DAO.
	group      string           // group is the database configuration group name of current DAO.
	columns    CornColumns      // columns contains all the column names of Table for convenient usage.
	primaryKey string           // 主键ID
	columnArr  []string         // 所有字段的数组
	columnArrG *garray.StrArray // 所有字段的数组（该格式更方便使用）
}

// CornColumns defines and stores column names for table platform_corn.
type CornColumns struct {
	CornId      string // 定时器ID
	CornCode    string // 标识
	CornName    string // 名称
	CornPattern string // 表达式
	Remark      string // 备注
	IsStop      string // 是否停用：0否 1是
	UpdatedAt   string // 更新时间
	CreatedAt   string // 创建时间
}

// cornColumns holds the columns for table platform_corn.
var cornColumns = CornColumns{
	CornId:      "cornId",
	CornCode:    "cornCode",
	CornName:    "cornName",
	CornPattern: "cornPattern",
	Remark:      "remark",
	IsStop:      "isStop",
	UpdatedAt:   "updatedAt",
	CreatedAt:   "createdAt",
}

// NewCornDao creates and returns a new DAO object for table data access.
func NewCornDao() *CornDao {
	return &CornDao{
		group:   `default`,
		table:   `platform_corn`,
		columns: cornColumns,
		primaryKey: func() string {
			return reflect.ValueOf(cornColumns).Field(0).String()
		}(),
		columnArr: func() []string {
			v := reflect.ValueOf(cornColumns)
			count := v.NumField()
			column := make([]string, count)
			for i := 0; i < count; i++ {
				column[i] = v.Field(i).String()
			}
			return column
		}(),
		columnArrG: func() *garray.StrArray {
			v := reflect.ValueOf(cornColumns)
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
func (dao *CornDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CornDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CornDao) Columns() CornColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CornDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CornDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CornDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// 主键ID
func (dao *CornDao) PrimaryKey() string {
	return dao.primaryKey
}

// 所有字段的数组
func (dao *CornDao) ColumnArr() []string {
	return dao.columnArr
}

// 所有字段的数组（该格式更方便使用）
func (dao *CornDao) ColumnArrG() *garray.StrArray {
	return dao.columnArrG
}
