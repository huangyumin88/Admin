package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

/*--------列表 开始--------*/
type AccountListReq struct {
	g.Meta `path:"/account/list" method:"post" tags:"平台后台/苹果" sm:"列表"`
	Filter AccountListFilter `json:"filter" dc:"查询条件"`
	Field  []string          `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回全部字段。注意：如前端页面所需字段较少，建议传指定字段，可大幅减轻服务器及数据库压力"`
	Sort   string            `json:"sort" default:"id DESC" dc:"排序"`
	Page   int               `json:"page" v:"integer|min:1" default:"1" dc:"页码"`
	Limit  int               `json:"limit" v:"integer|min:0" default:"10" dc:"每页数量。可传0取全部"`
}

type AccountListFilter struct {
	Id             *uint       `json:"id,omitempty" v:"integer|min:1" dc:"ID"`
	IdArr          []uint      `json:"idArr,omitempty" v:"distinct|foreach|integer|foreach|min:1" dc:"ID数组"`
	ExcId          *uint       `json:"excId,omitempty" v:"integer|min:1" dc:"排除ID"`
	ExcIdArr       []uint      `json:"excIdArr,omitempty" v:"distinct|foreach|integer|foreach|min:1" dc:"排除ID数组"`
	Label          string      `json:"label,omitempty" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"标签。常用于前端组件"`
	Account        string      `json:"account,omitempty" v:"length:1,255" dc:"账号"`
	Pwd            string      `json:"pwd,omitempty" v:"length:1,255" dc:"密码"`
	CountryId      *int        `json:"country_id,omitempty" v:"integer|min:1" dc:"国家id"`
	CountryCode    *string     `json:"country_code,omitempty" v:"length:1,255" dc:"国家"`
	Balance        string      `json:"balance,omitempty" v:"length:1,255" dc:"余额"`
	Status         *int        `json:"status,omitempty" v:"integer|in:0,1" dc:"禁用：0否 1是"`
	LoginStatus    *int        `json:"login_status,omitempty" v:"integer|in:0,1" dc:"登录：0否 1是"`
	IsStop         *uint       `json:"isStop,omitempty" v:"integer|in:0,1" dc:"停用：0否 1是"`
	TimeRangeStart *gtime.Time `json:"timeRangeStart,omitempty" v:"date-format:Y-m-d H:i:s" dc:"开始时间：YYYY-mm-dd HH:ii:ss"`
	TimeRangeEnd   *gtime.Time `json:"timeRangeEnd,omitempty" v:"date-format:Y-m-d H:i:s|after-equal:TimeRangeStart" dc:"结束时间：YYYY-mm-dd HH:ii:ss"`
}

type AccountListRes struct {
	Count int               `json:"count" dc:"总数"`
	List  []AccountListItem `json:"list" dc:"列表"`
}

type AccountListItem struct {
	Id            *uint       `json:"id,omitempty" dc:"ID"`
	Label         *string     `json:"label,omitempty" dc:"标签。常用于前端组件"`
	Account       *string     `json:"account,omitempty" dc:"账号"`
	Pwd           *string     `json:"pwd,omitempty" dc:"密码"`
	CountryId     *int        `json:"country_id,omitempty" dc:"国家id"`
	CountryCode   *string     `json:"country_code,omitempty" dc:"国家"`
	Balance       *string     `json:"balance,omitempty" dc:"余额"`
	Status        *int        `json:"status,omitempty" dc:"禁用：0否 1是"`
	Info          *string     `json:"info,omitempty" dc:"信息"`
	Device_id     *string     `json:"device_id,omitempty" dc:"设备"`
	Str_timestamp *string     `json:"str_timestamp,omitempty" dc:"时间"`
	Cookies       *string     `json:"cookies,omitempty" dc:"cookies"`
	Stk           *string     `json:"stk,omitempty" dc:"stk"`
	LoginStatus   *int        `json:"login_status,omitempty" dc:"登录：0否 1是"`
	IsStop        *uint       `json:"isStop,omitempty" dc:"停用：0否 1是"`
	UpdatedAt     *gtime.Time `json:"updatedAt,omitempty" dc:"更新时间"`
	CreatedAt     *gtime.Time `json:"createdAt,omitempty" dc:"创建时间"`
}

/*--------列表 结束--------*/

/*--------详情 开始--------*/
type AccountInfoReq struct {
	g.Meta `path:"/account/info" method:"post" tags:"平台后台/苹果" sm:"详情"`
	Id     uint     `json:"id" v:"required|integer|min:1" dc:"ID"`
	Field  []string `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回全部字段。注意：如前端页面所需字段较少，建议传指定字段，可大幅减轻服务器及数据库压力"`
}

type AccountInfoRes struct {
	Info AccountInfo `json:"info" dc:"详情"`
}

type AccountInfo struct {
	Id            *uint       `json:"id,omitempty" dc:"ID"`
	Label         *string     `json:"label,omitempty" dc:"标签。常用于前端组件"`
	Account       *string     `json:"account,omitempty" dc:"账号"`
	Pwd           *string     `json:"pwd,omitempty" dc:"密码"`
	CountryId     *int        `json:"country_id,omitempty" dc:"国家id"`
	CountryCode   *string     `json:"country_code,omitempty" dc:"国家"`
	Balance       *string     `json:"balance,omitempty" dc:"余额"`
	Status        *int        `json:"status,omitempty" dc:"禁用：0否 1是"`
	Info          *string     `json:"info,omitempty" dc:"信息"`
	Device_id     *string     `json:"device_id,omitempty" dc:"设备"`
	Str_timestamp *string     `json:"str_timestamp,omitempty" dc:"时间"`
	Cookies       *string     `json:"cookies,omitempty" dc:"cookies"`
	Stk           *string     `json:"stk,omitempty" dc:"stk"`
	LoginStatus   *int        `json:"login_status,omitempty" dc:"登录：0否 1是"`
	IsStop        *uint       `json:"isStop,omitempty" dc:"停用：0否 1是"`
	UpdatedAt     *gtime.Time `json:"updatedAt,omitempty" dc:"更新时间"`
	CreatedAt     *gtime.Time `json:"createdAt,omitempty" dc:"创建时间"`
}

/*--------详情 结束--------*/

/*--------新增 开始--------*/
type AccountCreateReq struct {
	g.Meta        `path:"/account/create" method:"post" tags:"平台后台/苹果" sm:"创建"`
	Account       *string `json:"account,omitempty" v:"length:1,255" dc:"账号"`
	Pwd           *string `json:"pwd,omitempty" v:"length:1,255" dc:"密码"`
	CountryId     *int    `json:"country_id,omitempty" v:"integer|min:1" dc:"国家id"`
	CountryCode   *string `json:"country_code,omitempty" v:"length:1,255" dc:"国家"`
	Balance       *string `json:"balance,omitempty" v:"length:1,255" dc:"余额"`
	Status        *int    `json:"status,omitempty" v:"integer|in:0,1" dc:"禁用：0否 1是"`
	Info          *string `json:"info,omitempty" v:"" dc:"信息"`
	Device_id     *string `json:"device_id,omitempty" dc:"设备"`
	Str_timestamp *string `json:"str_timestamp,omitempty" dc:"时间"`
	Cookies       *string `json:"cookies,omitempty" v:"" dc:"cookies"`
	Stk           *string `json:"stk,omitempty" dc:"stk"`
	LoginStatus   *int    `json:"login_status,omitempty" v:"integer|in:0,1" dc:"登录：0否 1是"`
	IsStop        *uint   `json:"isStop,omitempty" v:"integer|in:0,1" dc:"停用：0否 1是"`
}

/*--------新增 结束--------*/

/*--------修改 开始--------*/
type AccountUpdateReq struct {
	g.Meta        `path:"/account/update" method:"post" tags:"平台后台/苹果" sm:"更新"`
	IdArr         []uint  `json:"idArr,omitempty" v:"required|distinct|foreach|integer|foreach|min:1" dc:"ID数组"`
	Account       *string `json:"account,omitempty" v:"length:1,255" dc:"账号"`
	Pwd           *string `json:"pwd,omitempty" v:"length:1,255" dc:"密码"`
	CountryId     *int    `json:"country_id,omitempty" v:"integer|min:1" dc:"国家id"`
	CountryCode   *string `json:"country_code,omitempty" v:"length:1,255" dc:"国家"`
	Balance       *string `json:"balance,omitempty" v:"length:1,255" dc:"余额"`
	Status        *int    `json:"status,omitempty" v:"integer|in:0,1" dc:"禁用：0否 1是"`
	Info          *string `json:"info,omitempty" v:"" dc:"信息"`
	Device_id     *string `json:"device_id,omitempty" v:"" dc:"设备"`
	Str_timestamp *string `json:"str_timestamp,omitempty" v:"length:1,255" dc:"时间"`
	Cookies       *string `json:"cookies,omitempty" v:"" dc:"cookies"`
	Stk           *string `json:"stk,omitempty" dc:"stk"`
	LoginStatus   *int    `json:"login_status,omitempty" v:"integer|in:0,1" dc:"登录：0否 1是"`
	IsStop        *uint   `json:"isStop,omitempty" v:"integer|in:0,1" dc:"停用：0否 1是"`
}

/*--------修改 结束--------*/

/*--------删除 开始--------*/
type AccountDeleteReq struct {
	g.Meta `path:"/account/del" method:"post" tags:"平台后台/苹果" sm:"删除"`
	IdArr  []uint `json:"idArr,omitempty" v:"required|distinct|foreach|integer|foreach|min:1" dc:"ID数组"`
}

/*--------删除 结束--------*/

type AccountLoginReq struct {
	g.Meta  `path:"/account/login" method:"post" tags:"平台后台/苹果" sm:"登录"`
	Account *string `json:"account,omitempty" v:"length:1,255" dc:"账号"`
	Pwd     *string `json:"pwd,omitempty" v:"length:1,255" dc:"密码"`
}

type AccountGiftCardQueryReq struct {
	g.Meta      `path:"/account/giftcard/query" method:"post" tags:"苹果/礼品卡查询" sm:"礼品卡查询"`
	Account     *string `json:"account,omitempty" v:"length:1,255" dc:"账号"`
	Pwd         *string `json:"pwd,omitempty" v:"length:1,255" dc:"密码"`
	GiftCardPin *string `json:"giftCardPin,omitempty" v:"length:1,255" dc:"pin"`
}

type AccountGiftCardInfo struct {
	CountryCode string `json:"country_code,omitempty" v:"length:1,255" dc:"国家"`
	Balance     string `json:"balance,omitempty" v:"length:1,255" dc:"余额"`
}

type AccountGiftCardInfoRes struct {
	Info AccountGiftCardInfo `json:"info" dc:"详情"`
}
