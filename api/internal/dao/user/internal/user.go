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

// UserDao is the data access object for table user.
type UserDao struct {
	table      string           // table is the underlying table name of the DAO.
	group      string           // group is the database configuration group name of current DAO.
	columns    UserColumns      // columns contains all the column names of Table for convenient usage.
	primaryKey string           // 主键ID
	columnArr  []string         // 所有字段的数组
	columnArrG *garray.StrArray // 所有字段的数组（该格式更方便使用）
}

// UserColumns defines and stores column names for table user.
type UserColumns struct {
	UserId       string // 用户ID
	Phone        string // 手机
	Email        string // 邮箱
	Account      string // 账号
	Password     string // 密码。md5保存
	Salt         string // 加密盐
	Nickname     string // 昵称
	Avatar       string // 头像
	Gender       string // 性别：0未设置 1男 2女
	Birthday     string // 生日
	Address      string // 详细地址
	IdCardName   string // 身份证姓名
	IdCardNo     string // 身份证号码
	ReferralCode string // 推荐码
	IsStop       string // 停用：0否 1是
	UpdatedAt    string // 更新时间
	CreatedAt    string // 创建时间
	Country      string // 国家
	WalletId	 string // 钱包
	ImUserId     string // ImUserId
	ImUserSig	 string // ImUserSig
}

// userColumns holds the columns for table user.
var userColumns = UserColumns{
	UserId:       "userId",
	Phone:        "phone",
	Email:        "email",
	Account:      "account",
	Password:     "password",
	Salt:         "salt",
	Nickname:     "nickname",
	Avatar:       "avatar",
	Gender:       "gender",
	Birthday:     "birthday",
	Address:      "address",
	IdCardName:   "idCardName",
	IdCardNo:     "idCardNo",
	ReferralCode: "referralCode",
	IsStop:       "isStop",
	UpdatedAt:    "updatedAt",
	CreatedAt:    "createdAt",
	Country:      "country",
	WalletId:     "wallet_id",
	ImUserId:     "imUserId",
	ImUserSig:	  "imUserSig",
}

// NewUserDao creates and returns a new DAO object for table data access.
func NewUserDao() *UserDao {
	return &UserDao{
		group:   `default`,
		table:   `user`,
		columns: userColumns,
		primaryKey: func() string {
			return reflect.ValueOf(userColumns).Field(0).String()
		}(),
		columnArr: func() []string {
			v := reflect.ValueOf(userColumns)
			count := v.NumField()
			column := make([]string, count)
			for i := 0; i < count; i++ {
				column[i] = v.Field(i).String()
			}
			return column
		}(),
		columnArrG: func() *garray.StrArray {
			v := reflect.ValueOf(userColumns)
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
func (dao *UserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
// 使用较为频繁。为优化内存考虑，改成返回指针更为合适，但切忌使用过程中不可修改，否则会污染全局
func (dao *UserDao) Columns() *UserColumns {
	return &dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// 主键ID
func (dao *UserDao) PrimaryKey() string {
	return dao.primaryKey
}

// 所有字段的数组
func (dao *UserDao) ColumnArr() []string {
	return dao.columnArr
}

// 所有字段的数组（该格式更方便使用）
func (dao *UserDao) ColumnArrG() *garray.StrArray {
	return dao.columnArrG
}
