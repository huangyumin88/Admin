package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

/*--------列表 开始--------*/
type ImageListReq struct {
	g.Meta `path:"/image/list" method:"post" tags:"平台后台/图片" sm:"列表"`
	Filter ImageListFilter `json:"filter" dc:"过滤条件"`
	Field  []string        `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回全部字段。注意：如前端页面所需字段较少，建议传指定字段，可大幅减轻服务器及数据库压力"`
	Sort   string          `json:"sort" default:"id DESC" dc:"排序"`
	Page   int             `json:"page" v:"min:1" default:"1" dc:"页码"`
	Limit  int             `json:"limit" v:"min:0" default:"10" dc:"每页数量。可传0取全部"`
}

type ImageListFilter struct {
	Id             *uint       `json:"id,omitempty" v:"min:1" dc:"ID"`
	IdArr          []uint      `json:"idArr,omitempty" v:"distinct|foreach|min:1" dc:"ID数组"`
	ExcId          *uint       `json:"excId,omitempty" v:"min:1" dc:"排除ID"`
	ExcIdArr       []uint      `json:"excIdArr,omitempty" v:"distinct|foreach|min:1" dc:"排除ID数组"`
	Url            string      `json:"url,omitempty" v:"max-length:255|url" dc:""`
	IsStop         *uint       `json:"isStop,omitempty" v:"in:0,1" dc:"停用：0否 1是"`
	TimeRangeStart *gtime.Time `json:"timeRangeStart,omitempty" v:"date-format:Y-m-d H:i:s" dc:"开始时间：YYYY-mm-dd HH:ii:ss"`
	TimeRangeEnd   *gtime.Time `json:"timeRangeEnd,omitempty" v:"date-format:Y-m-d H:i:s|after-equal:TimeRangeStart" dc:"结束时间：YYYY-mm-dd HH:ii:ss"`
}

type ImageListRes struct {
	Count int             `json:"count" dc:"总数"`
	List  []ImageListItem `json:"list" dc:"列表"`
}

type ImageListItem struct {
	Id        *uint       `json:"id,omitempty" dc:"ID"`
	Url       *string     `json:"url,omitempty" dc:""`
	Avatar    *string     `json:"avatar,omitempty" dc:"头像"`
	IsStop    *uint       `json:"isStop,omitempty" dc:"停用：0否 1是"`
	UpdatedAt *gtime.Time `json:"updatedAt,omitempty" dc:"更新时间"`
	CreatedAt *gtime.Time `json:"createdAt,omitempty" dc:"创建时间"`
}

/*--------列表 结束--------*/

/*--------详情 开始--------*/
type ImageInfoReq struct {
	g.Meta `path:"/image/info" method:"post" tags:"平台后台/图片" sm:"详情"`
	Id     uint     `json:"id" v:"required|min:1" dc:"ID"`
	Field  []string `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回全部字段。注意：如前端页面所需字段较少，建议传指定字段，可大幅减轻服务器及数据库压力"`
}

type ImageInfoRes struct {
	Info ImageInfo `json:"info" dc:"详情"`
}

type ImageInfo struct {
	Id        *uint       `json:"id,omitempty" dc:"ID"`
	Url       *string     `json:"url,omitempty" dc:""`
	Avatar    *string     `json:"avatar,omitempty" dc:"头像"`
	IsStop    *uint       `json:"isStop,omitempty" dc:"停用：0否 1是"`
	UpdatedAt *gtime.Time `json:"updatedAt,omitempty" dc:"更新时间"`
	CreatedAt *gtime.Time `json:"createdAt,omitempty" dc:"创建时间"`
}

/*--------详情 结束--------*/

/*--------新增 开始--------*/
type ImageCreateReq struct {
	g.Meta `path:"/image/create" method:"post" tags:"平台后台/图片" sm:"新增"`
	Url    *string `json:"url,omitempty" v:"max-length:255|url" dc:""`
	Avatar *string `json:"avatar,omitempty" v:"max-length:200|url" dc:"头像"`
	IsStop *uint   `json:"isStop,omitempty" v:"in:0,1" dc:"停用：0否 1是"`
}

/*--------新增 结束--------*/

/*--------修改 开始--------*/
type ImageUpdateReq struct {
	g.Meta `path:"/image/update" method:"post" tags:"平台后台/图片" sm:"修改"`
	IdArr  []uint  `json:"idArr,omitempty" v:"required|distinct|foreach|min:1" dc:"ID数组"`
	Url    *string `json:"url,omitempty" v:"max-length:255|url" dc:""`
	Avatar *string `json:"avatar,omitempty" v:"max-length:200|url" dc:"头像"`
	IsStop *uint   `json:"isStop,omitempty" v:"in:0,1" dc:"停用：0否 1是"`
}

/*--------修改 结束--------*/

/*--------删除 开始--------*/
type ImageDeleteReq struct {
	g.Meta `path:"/image/del" method:"post" tags:"平台后台/图片" sm:"删除"`
	IdArr  []uint `json:"idArr,omitempty" v:"required|distinct|foreach|min:1" dc:"ID数组"`
}

/*--------删除 结束--------*/
