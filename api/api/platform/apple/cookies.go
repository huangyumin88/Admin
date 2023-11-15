package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

/*--------列表 开始--------*/
type CookiesListReq struct {
	g.Meta `path:"/cookies/list" method:"post" tags:"平台后台/苹果官网登录保存信息" sm:"列表"`
	Filter CookiesListFilter `json:"filter" dc:"查询条件"`
	Field  []string          `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回全部字段。注意：如前端页面所需字段较少，建议传指定字段，可大幅减轻服务器及数据库压力"`
	Sort   string            `json:"sort" default:"id DESC" dc:"排序"`
	Page   int               `json:"page" v:"integer|min:1" default:"1" dc:"页码"`
	Limit  int               `json:"limit" v:"integer|min:0" default:"10" dc:"每页数量。可传0取全部"`
}

type CookiesListFilter struct {
	Id             *uint       `json:"id,omitempty" v:"integer|min:1" dc:"ID"`
	IdArr          []uint      `json:"idArr,omitempty" v:"distinct|foreach|integer|foreach|min:1" dc:"ID数组"`
	ExcId          *uint       `json:"excId,omitempty" v:"integer|min:1" dc:"排除ID"`
	ExcIdArr       []uint      `json:"excIdArr,omitempty" v:"distinct|foreach|integer|foreach|min:1" dc:"排除ID数组"`
	Label          string      `json:"label,omitempty" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"标签。常用于前端组件"`
	CountryId      *int        `json:"country_id,omitempty" v:"integer|min:1" dc:"国家id"`
	CountryCode    string      `json:"country_code,omitempty" v:"length:1,255|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"国家代码"`
	Url            *string     `json:"url,omitempty" v:"length:1,255" dc:"请求链接"`
	TimeRangeStart *gtime.Time `json:"timeRangeStart,omitempty" v:"date-format:Y-m-d H:i:s" dc:"开始时间：YYYY-mm-dd HH:ii:ss"`
	TimeRangeEnd   *gtime.Time `json:"timeRangeEnd,omitempty" v:"date-format:Y-m-d H:i:s|after-equal:TimeRangeStart" dc:"结束时间：YYYY-mm-dd HH:ii:ss"`
	StrTimestamp   string      `json:"str_timestamp,omitempty" v:"length:1,30" dc:"时间戳"`
	Account        string      `json:"account,omitempty" v:"length:1,255" dc:"账号"`
}

type CookiesListRes struct {
	Count int               `json:"count" dc:"总数"`
	List  []CookiesListItem `json:"list" dc:"列表"`
}

type CookiesListItem struct {
	Id           *uint       `json:"id,omitempty" dc:"ID"`
	Label        *string     `json:"label,omitempty" dc:"标签。常用于前端组件"`
	Cookies      *string     `json:"cookies,omitempty" dc:"cookies"`
	CountryId    *int        `json:"country_id,omitempty" dc:"国家id"`
	CountryCode  *string     `json:"country_code,omitempty" dc:"国家代码"`
	Url          *string     `json:"url,omitempty"  dc:"请求链接"`
	UpdatedAt    *gtime.Time `json:"updatedAt,omitempty" dc:"更新时间"`
	CreatedAt    *gtime.Time `json:"createdAt,omitempty" dc:"创建时间"`
	Headers      *string     `json:"headers,omitempty" dc:"信息头"`
	StrTimestamp *string     `json:"str_timestamp,omitempty" dc:"时间戳"`
	Account      *string     `json:"account,omitempty" dc:"账号"`
	DeviceId     *string     `json:"device_id,omitempty" dc:"设备id"`
}

/*--------列表 结束--------*/

/*--------详情 开始--------*/
type CookiesInfoReq struct {
	g.Meta `path:"/cookies/info" method:"post" tags:"平台后台/苹果官网登录保存信息" sm:"详情"`
	Id     uint     `json:"id" v:"required|integer|min:1" dc:"ID"`
	Field  []string `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回全部字段。注意：如前端页面所需字段较少，建议传指定字段，可大幅减轻服务器及数据库压力"`
}

type CookiesInfoRes struct {
	Info CookiesInfo `json:"info" dc:"详情"`
}

type CookiesInfo struct {
	Id           *uint       `json:"id,omitempty" dc:"ID"`
	Label        *string     `json:"label,omitempty" dc:"标签。常用于前端组件"`
	Cookies      *string     `json:"cookies,omitempty" dc:"cookies"`
	CountryId    *int        `json:"country_id,omitempty" dc:"国家id"`
	CountryCode  *string     `json:"country_code,omitempty" dc:"国家代码"`
	Url          *string     `json:"url,omitempty" dc:"请求链接"`
	UpdatedAt    *gtime.Time `json:"updatedAt,omitempty" dc:"更新时间"`
	CreatedAt    *gtime.Time `json:"createdAt,omitempty" dc:"创建时间"`
	Headers      *string     `json:"headers,omitempty" dc:"信息头"`
	StrTimestamp *string     `json:"str_timestamp,omitempty" dc:"时间戳"`
	Account      *string     `json:"account,omitempty" dc:"账号"`
	DeviceId     *string     `json:"device_id,omitempty" dc:"设备id"`
}

/*--------详情 结束--------*/

/*--------新增 开始--------*/
type CookiesCreateReq struct {
	g.Meta       `path:"/cookies/create" method:"post" tags:"平台后台/苹果官网登录保存信息" sm:"创建"`
	Cookies      *string `json:"cookies,omitempty" v:"" dc:"cookies"`
	CountryId    *int    `json:"country_id,omitempty" v:"integer|min:1" dc:"国家id"`
	CountryCode  *string `json:"country_code,omitempty" v:"length:1,255|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"国家代码"`
	Url          *string `json:"url,omitempty" v:"length:1,255" dc:"请求链接"`
	Headers      *string `json:"headers,omitempty" v:"" dc:"信息头"`
	StrTimestamp *string `json:"str_timestamp,omitempty" v:"length:1,30" dc:"时间戳"`
	Account      *string `json:"account,omitempty" v:"length:1,255" dc:"账号"`
	DeviceId     *string `json:"device_id,omitempty" v:"" dc:"设备id"`
}

/*--------新增 结束--------*/

/*--------修改 开始--------*/
type CookiesUpdateReq struct {
	g.Meta       `path:"/cookies/update" method:"post" tags:"平台后台/苹果官网登录保存信息" sm:"更新"`
	IdArr        []uint  `json:"idArr,omitempty" v:"required|distinct|foreach|integer|foreach|min:1" dc:"ID数组"`
	Cookies      *string `json:"cookies,omitempty" v:"" dc:"cookies"`
	CountryId    *int    `json:"country_id,omitempty" v:"integer|min:1" dc:"国家id"`
	CountryCode  *string `json:"country_code,omitempty" v:"length:1,255|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"国家代码"`
	Url          *string `json:"url,omitempty" v:"length:1,255" dc:"请求链接"`
	Headers      *string `json:"headers,omitempty" v:"" dc:"信息头"`
	StrTimestamp *string `json:"str_timestamp,omitempty" v:"length:1,30" dc:"时间戳"`
	Account      *string `json:"account,omitempty" v:"length:1,255" dc:"账号"`
	DeviceId     *string `json:"device_id,omitempty" v:"" dc:"设备id"`
}

/*--------修改 结束--------*/

/*--------删除 开始--------*/
type CookiesDeleteReq struct {
	g.Meta `path:"/cookies/del" method:"post" tags:"平台后台/苹果官网登录保存信息" sm:"删除"`
	IdArr  []uint `json:"idArr,omitempty" v:"required|distinct|foreach|integer|foreach|min:1" dc:"ID数组"`
}

/*--------删除 结束--------*/
