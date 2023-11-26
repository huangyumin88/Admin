package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

/*--------列表 开始--------*/
type BanksUserListReq struct {
	g.Meta `path:"/banksUser/list" method:"post" tags:"平台后台/银行用户中间表" sm:"列表"`
	Filter BanksUserListFilter `json:"filter" dc:"查询条件"`
	Field  []string            `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回全部字段。注意：如前端页面所需字段较少，建议传指定字段，可大幅减轻服务器及数据库压力"`
	Sort   string              `json:"sort" default:"id DESC" dc:"排序"`
	Page   int                 `json:"page" v:"min:1" default:"1" dc:"页码"`
	Limit  int                 `json:"limit" v:"min:0" default:"10" dc:"每页数量。可传0取全部"`
}

type BanksUserListFilter struct {
	Id       *uint  `json:"id,omitempty" v:"min:1" dc:"ID"`
	IdArr    []uint `json:"idArr,omitempty" v:"distinct|foreach|min:1" dc:"ID数组"`
	ExcId    *uint  `json:"excId,omitempty" v:"min:1" dc:"排除ID"`
	ExcIdArr []uint `json:"excIdArr,omitempty" v:"distinct|foreach|min:1" dc:"排除ID数组"`
	//Id             *int        `json:"id,omitempty" v:"min:1" dc:"ID"`
	UserId         *uint       `json:"user_id,omitempty" v:"min:1" dc:"用户id"`
	BankId         *int        `json:"bank_id,omitempty" v:"min:1" dc:"银行id"`
	TimeRangeStart *gtime.Time `json:"timeRangeStart,omitempty" v:"date-format:Y-m-d H:i:s" dc:"开始时间：YYYY-mm-dd HH:ii:ss"`
	TimeRangeEnd   *gtime.Time `json:"timeRangeEnd,omitempty" v:"date-format:Y-m-d H:i:s|after-equal:TimeRangeStart" dc:"结束时间：YYYY-mm-dd HH:ii:ss"`
}

type BanksUserListRes struct {
	Count int                 `json:"count" dc:"总数"`
	List  []BanksUserListItem `json:"list" dc:"列表"`
}

type BanksUserListItem struct {
	Id *uint `json:"id,omitempty" dc:"ID"`
	//Id        *int        `json:"id,omitempty" dc:"ID"`
	UserId    *uint       `json:"user_id,omitempty" dc:"用户id"`
	BankId    *int        `json:"bank_id,omitempty" dc:"银行id"`
	UpdatedAt *gtime.Time `json:"updatedAt,omitempty" dc:"更新时间"`
	CreatedAt *gtime.Time `json:"createdAt,omitempty" dc:"创建时间"`
	UserName  *string     `json:"user_name,omitempty" dc:"用户"`
}

/*--------列表 结束--------*/

/*--------删除 开始--------*/
type BanksUserDeleteReq struct {
	g.Meta `path:"/banksUser/del" method:"post" tags:"平台后台/银行用户中间表" sm:"删除"`
	IdArr  []uint `json:"idArr,omitempty" v:"required|distinct|foreach|min:1" dc:"ID数组"`
}

/*--------删除 结束--------*/
