package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

/*--------列表 开始--------*/
type CardCategoriesSubListReq struct {
	g.Meta `path:"/cardCategoriesSub/list" method:"post" tags:"APP/子分类" sm:"列表"`
	Filter CardCategoriesSubListFilter `json:"filter" dc:"查询条件"`
	Field  []string                    `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回全部字段。注意：如前端页面所需字段较少，建议传指定字段，可大幅减轻服务器及数据库压力"`
	Sort   string                      `json:"sort" default:"id DESC" dc:"排序"`
	Page   int                         `json:"page" v:"min:1" default:"1" dc:"页码"`
	Limit  int                         `json:"limit" v:"min:0" default:"10" dc:"每页数量。可传0取全部"`
}

type CardCategoriesSubListFilter struct {
	Id             *uint       `json:"id,omitempty" v:"min:1" dc:"ID"`
	IdArr          []uint      `json:"idArr,omitempty" v:"distinct|foreach|min:1" dc:"ID数组"`
	ExcId          *uint       `json:"excId,omitempty" v:"min:1" dc:"排除ID"`
	ExcIdArr       []uint      `json:"excIdArr,omitempty" v:"distinct|foreach|min:1" dc:"排除ID数组"`
	Label          string      `json:"label,omitempty" v:"max-length:30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"标签。常用于前端组件"`
	CateId         *int        `json:"cate_id,omitempty" v:"min:1" dc:"主分类ID"`
	Name           string      `json:"name,omitempty" v:"max-length:255|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"card 名字"`
	IsActive       *int        `json:"isActive,omitempty" v:"in:0,1" dc:"活动：0否 1是"`
	IsStop         *uint       `json:"isStop,omitempty" v:"in:0,1" dc:"停用：0否 1是"`
	SubId          string      `json:"sub_id,omitempty" v:"max-length:30" dc:"Card id"`
	TimeRangeStart *gtime.Time `json:"timeRangeStart,omitempty" v:"date-format:Y-m-d H:i:s" dc:"开始时间：YYYY-mm-dd HH:ii:ss"`
	TimeRangeEnd   *gtime.Time `json:"timeRangeEnd,omitempty" v:"date-format:Y-m-d H:i:s|after-equal:TimeRangeStart" dc:"结束时间：YYYY-mm-dd HH:ii:ss"`
}

type CardCategoriesSubListRes struct {
	Count int                         `json:"count" dc:"总数"`
	List  []CardCategoriesSubListItem `json:"list" dc:"列表"`
}

type CardCategoriesSubListItem struct {
	Id                  *uint       `json:"id,omitempty" dc:"ID"`
	Label               *string     `json:"label,omitempty" dc:"标签。常用于前端组件"`
	CateId              *int        `json:"cate_id,omitempty" dc:"主分类ID"`
	Name                *string     `json:"name,omitempty" dc:"card 名字"`
	Rate                *int        `json:"rate,omitempty" dc:"费率"`
	MinAcceptableAmount *int        `json:"minAcceptableAmount,omitempty" dc:"最低可接受金额"`
	Announcements       *string     `json:"announcements,omitempty" dc:"注意事项"`
	Sort                *uint       `json:"sort,omitempty" dc:"排序值。从小到大排序，默认50，范围0-100"`
	IsActive            *int        `json:"isActive,omitempty" dc:"活动：0否 1是"`
	IsStop              *uint       `json:"isStop,omitempty" dc:"停用：0否 1是"`
	SubId               *string     `json:"sub_id,omitempty" dc:"Card id"`
	UpdatedAt           *gtime.Time `json:"updatedAt,omitempty" dc:"更新时间"`
	CreatedAt           *gtime.Time `json:"createdAt,omitempty" dc:"创建时间"`
}

/*--------列表 结束--------*/
