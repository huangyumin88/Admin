// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Image is the golang structure of table ai_image for DAO operations like Where/Data.
type Image struct {
	g.Meta    `orm:"table:ai_image, do:true"`
	Id        interface{} //
	Url       interface{} //
	Avatar    interface{} // 头像
	IsStop    interface{} // 停用：0否 1是
	UpdatedAt *gtime.Time // 更新时间
	CreatedAt *gtime.Time // 创建时间
}
