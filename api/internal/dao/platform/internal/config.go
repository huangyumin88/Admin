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

// ConfigDao is the data access object for table platform_config.
type ConfigDao struct {
	table      string           // table is the underlying table name of the DAO.
	group      string           // group is the database configuration group name of current DAO.
	columns    ConfigColumns    // columns contains all the column names of Table for convenient usage.
	primaryKey string           // 主键ID
	columnArr  []string         // 所有字段的数组
	columnArrG *garray.StrArray // 所有字段的数组（该格式更方便使用）
}

// ConfigColumns defines and stores column names for table platform_config.
type ConfigColumns struct {
	ConfigId    string // 配置ID
	ConfigKey   string // 配置项Key
	ConfigValue string // 配置项值（设置大点。以后可能需要保存富文本内容，如公司简介或协议等等）
	UpdatedAt   string // 更新时间
	CreatedAt   string // 创建时间
}

// configColumns holds the columns for table platform_config.
var configColumns = ConfigColumns{
	ConfigId:    "configId",
	ConfigKey:   "configKey",
	ConfigValue: "configValue",
	UpdatedAt:   "updatedAt",
	CreatedAt:   "createdAt",
}

// NewConfigDao creates and returns a new DAO object for table data access.
func NewConfigDao() *ConfigDao {
	return &ConfigDao{
		group:   `default`,
		table:   `platform_config`,
		columns: configColumns,
		primaryKey: func() string {
			return reflect.ValueOf(configColumns).Field(0).String()
		}(),
		columnArr: func() []string {
			v := reflect.ValueOf(configColumns)
			count := v.NumField()
			column := make([]string, count)
			for i := 0; i < count; i++ {
				column[i] = v.Field(i).String()
			}
			return column
		}(),
		columnArrG: func() *garray.StrArray {
			v := reflect.ValueOf(configColumns)
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
func (dao *ConfigDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ConfigDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ConfigDao) Columns() ConfigColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ConfigDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ConfigDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ConfigDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// 主键ID
func (dao *ConfigDao) PrimaryKey() string {
	return dao.primaryKey
}

// 所有字段的数组
func (dao *ConfigDao) ColumnArr() []string {
	return dao.columnArr
}

// 所有字段的数组（该格式更方便使用）
func (dao *ConfigDao) ColumnArrG() *garray.StrArray {
	return dao.columnArrG
}
