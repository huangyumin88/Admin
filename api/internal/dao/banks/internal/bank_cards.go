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

// BankCardsDao is the data access object for table app_card_bank_cards.
type BankCardsDao struct {
	table      string           // table is the underlying table name of the DAO.
	group      string           // group is the database configuration group name of current DAO.
	columns    BankCardsColumns // columns contains all the column names of Table for convenient usage.
	primaryKey string           // 主键ID
	columnArr  []string         // 所有字段的数组
	columnArrG *garray.StrArray // 所有字段的数组（该格式更方便使用）
}

// BankCardsColumns defines and stores column names for table app_card_bank_cards.
type BankCardsColumns struct {
	Id             string // ID
	UserId         string // 用户id
	BankId         string // 银行id
	CardNumber     string // 银行卡号
	CardHolderName string // 持卡人姓名
	ExpirationDate string // 银行卡的有效期
	Cvv            string // CVV 安全码
	UpdatedAt      string // 更新时间
	CreatedAt      string // 创建时间
}

// bankCardsColumns holds the columns for table app_card_bank_cards.
var bankCardsColumns = BankCardsColumns{
	Id:             "id",
	UserId:         "user_id",
	BankId:         "bank_id",
	CardNumber:     "card_number",
	CardHolderName: "card_holder_name",
	ExpirationDate: "expiration_date",
	Cvv:            "cvv",
	UpdatedAt:      "updatedAt",
	CreatedAt:      "createdAt",
}

// NewBankCardsDao creates and returns a new DAO object for table data access.
func NewBankCardsDao() *BankCardsDao {
	return &BankCardsDao{
		group:   `default`,
		table:   `app_card_bank_cards`,
		columns: bankCardsColumns,
		primaryKey: func() string {
			return reflect.ValueOf(bankCardsColumns).Field(0).String()
		}(),
		columnArr: func() []string {
			v := reflect.ValueOf(bankCardsColumns)
			count := v.NumField()
			column := make([]string, count)
			for i := 0; i < count; i++ {
				column[i] = v.Field(i).String()
			}
			return column
		}(),
		columnArrG: func() *garray.StrArray {
			v := reflect.ValueOf(bankCardsColumns)
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
func (dao *BankCardsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *BankCardsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
// 使用较为频繁。为优化内存考虑，改成返回指针更为合适，但切忌使用过程中不可修改，否则会污染全局
func (dao *BankCardsDao) Columns() *BankCardsColumns {
	return &dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *BankCardsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *BankCardsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *BankCardsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// 主键ID
func (dao *BankCardsDao) PrimaryKey() string {
	return dao.primaryKey
}

// 所有字段的数组
func (dao *BankCardsDao) ColumnArr() []string {
	return dao.columnArr
}

// 所有字段的数组（该格式更方便使用）
func (dao *BankCardsDao) ColumnArrG() *garray.StrArray {
	return dao.columnArrG
}
