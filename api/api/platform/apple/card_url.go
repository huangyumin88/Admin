package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

/*--------列表 开始--------*/
type CardUrlListReq struct {
	g.Meta `path:"/cardUrl/list" method:"post" tags:"平台后台/苹果礼品卡查询地址" sm:"列表"`
	Filter CardUrlListFilter `json:"filter" dc:"查询条件"`
	Field  []string          `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回全部字段。注意：如前端页面所需字段较少，建议传指定字段，可大幅减轻服务器及数据库压力"`
	Sort   string            `json:"sort" default:"id DESC" dc:"排序"`
	Page   int               `json:"page" v:"integer|min:1" default:"1" dc:"页码"`
	Limit  int               `json:"limit" v:"integer|min:0" default:"10" dc:"每页数量。可传0取全部"`
}

type CardUrlListFilter struct {
	Id             *uint       `json:"id,omitempty" v:"integer|min:1" dc:"ID"`
	IdArr          []uint      `json:"idArr,omitempty" v:"distinct|foreach|integer|foreach|min:1" dc:"ID数组"`
	ExcId          *uint       `json:"excId,omitempty" v:"integer|min:1" dc:"排除ID"`
	ExcIdArr       []uint      `json:"excIdArr,omitempty" v:"distinct|foreach|integer|foreach|min:1" dc:"排除ID数组"`
	Url            string      `json:"url,omitempty" v:"length:1,255|url" dc:"请求链接"`
	AccountId      *int        `json:"account_id,omitempty" v:"integer|min:1" dc:"苹果账号ID"`
	CountryCode    string      `json:"country_code,omitempty" v:"length:1,255|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"国家代码"`
	IsStop         *uint       `json:"isStop,omitempty" v:"integer|in:0,1" dc:"停用：0否 1是"`
	IsAutoLogin    *int        `json:"isAutoLogin,omitempty" v:"integer|in:0,1" dc:"自动登录：0否 1是"`
	TimeRangeStart *gtime.Time `json:"timeRangeStart,omitempty" v:"date-format:Y-m-d H:i:s" dc:"开始时间：YYYY-mm-dd HH:ii:ss"`
	TimeRangeEnd   *gtime.Time `json:"timeRangeEnd,omitempty" v:"date-format:Y-m-d H:i:s|after-equal:TimeRangeStart" dc:"结束时间：YYYY-mm-dd HH:ii:ss"`
}

type CardUrlListRes struct {
	Count int               `json:"count" dc:"总数"`
	List  []CardUrlListItem `json:"list" dc:"列表"`
}

type CardUrlListItem struct {
	Id          *uint        `json:"id,omitempty" dc:"ID"`
	Label       *string      `json:"label,omitempty" dc:"标签。常用于前端组件"`
	Url         *string      `json:"url,omitempty" dc:"请求链接"`
	AccountId   *int         `json:"account_id,omitempty" dc:"苹果账号ID"`
	CountryCode *string      `json:"country_code,omitempty" dc:"国家代码"`
	IsStop      *uint        `json:"isStop,omitempty" dc:"停用：0否 1是"`
	IsAutoLogin *int         `json:"isAutoLogin,omitempty" dc:"自动登录：0否 1是"`
	UpdatedAt   *gtime.Time  `json:"updatedAt,omitempty" dc:"更新时间"`
	CreatedAt   *gtime.Time  `json:"createdAt,omitempty" dc:"创建时间"`
	AccountInfo *AccountInfo `json:"account_info,omitempty" dc:"账号详情"`
}

/*--------列表 结束--------*/

/*--------详情 开始--------*/
type CardUrlInfoReq struct {
	g.Meta `path:"/cardUrl/info" method:"post" tags:"平台后台/苹果礼品卡查询地址" sm:"详情"`
	Id     uint     `json:"id" v:"required|integer|min:1" dc:"ID"`
	Field  []string `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回全部字段。注意：如前端页面所需字段较少，建议传指定字段，可大幅减轻服务器及数据库压力"`
}

type CardUrlAutomaticLoginReq struct {
	g.Meta `path:"/cardUrl/automaticLogin" method:"post" tags:"平台后台/苹果礼品卡查询地址" sm:"自动登录"`
}

type CardUrlInfoRes struct {
	Info CardUrlInfo `json:"info" dc:"详情"`
}

type CardUrlInfo struct {
	Id          *uint       `json:"id,omitempty" dc:"ID"`
	Url         *string     `json:"url,omitempty" dc:"请求链接"`
	AccountId   *int        `json:"account_id,omitempty" dc:"苹果账号ID"`
	CountryCode *string     `json:"country_code,omitempty" dc:"国家代码"`
	IsStop      *uint       `json:"isStop,omitempty" dc:"停用：0否 1是"`
	IsAutoLogin *int        `json:"isAutoLogin,omitempty" dc:"自动登录：0否 1是"`
	UpdatedAt   *gtime.Time `json:"updatedAt,omitempty" dc:"更新时间"`
	CreatedAt   *gtime.Time `json:"createdAt,omitempty" dc:"创建时间"`
}

/*--------详情 结束--------*/

/*--------新增 开始--------*/
type CardUrlCreateReq struct {
	g.Meta      `path:"/cardUrl/create" method:"post" tags:"平台后台/苹果礼品卡查询地址" sm:"创建"`
	Url         *string `json:"url,omitempty" v:"length:1,255|url" dc:"请求链接"`
	AccountId   *int    `json:"account_id,omitempty" v:"integer|min:1" dc:"苹果账号ID"`
	CountryCode *string `json:"country_code,omitempty" v:"length:1,255|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"国家代码"`
	IsStop      *uint   `json:"isStop,omitempty" v:"integer|in:0,1" dc:"停用：0否 1是"`
	IsAutoLogin *int    `json:"isAutoLogin,omitempty" v:"integer|in:0,1" dc:"自动登录：0否 1是"`
}

/*--------新增 结束--------*/

/*--------修改 开始--------*/
type CardUrlUpdateReq struct {
	g.Meta      `path:"/cardUrl/update" method:"post" tags:"平台后台/苹果礼品卡查询地址" sm:"更新"`
	IdArr       []uint  `json:"idArr,omitempty" v:"required|distinct|foreach|integer|foreach|min:1" dc:"ID数组"`
	Url         *string `json:"url,omitempty" v:"length:1,255" dc:"请求链接"`
	AccountId   *int    `json:"account_id,omitempty" v:"integer|min:1" dc:"苹果账号ID"`
	CountryCode *string `json:"country_code,omitempty" v:"length:1,255|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"国家代码"`
	IsStop      *uint   `json:"isStop,omitempty" v:"integer|in:0,1" dc:"停用：0否 1是"`
	IsAutoLogin *int    `json:"isAutoLogin,omitempty" v:"integer|in:0,1" dc:"自动登录：0否 1是"`
}

/*--------修改 结束--------*/

/*--------删除 开始--------*/
type CardUrlDeleteReq struct {
	g.Meta `path:"/cardUrl/del" method:"post" tags:"平台后台/苹果礼品卡查询地址" sm:"删除"`
	IdArr  []uint `json:"idArr,omitempty" v:"required|distinct|foreach|integer|foreach|min:1" dc:"ID数组"`
}

type CardUrlQueryReq struct {
	g.Meta      `path:"/cardUrl/giftcard/query" method:"post" tags:"平台后台/苹果礼品卡查询地址" sm:"自动查询"`
	GiftCardPin *string `json:"giftCardPin,omitempty" v:"length:1,255" dc:"pin"`
}

/*--------删除 结束--------*/
