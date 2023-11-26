package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

/*--------列表 开始--------*/
type BanksListReq struct {
	g.Meta `path:"/banks/list" method:"post" tags:"平台后台/银行" sm:"列表"`
	Filter BanksListFilter `json:"filter" dc:"查询条件"`
	Field  []string        `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回全部字段。注意：如前端页面所需字段较少，建议传指定字段，可大幅减轻服务器及数据库压力"`
	Sort   string          `json:"sort" default:"id DESC" dc:"排序"`
	Page   int             `json:"page" v:"min:1" default:"1" dc:"页码"`
	Limit  int             `json:"limit" v:"min:0" default:"10" dc:"每页数量。可传0取全部"`
}

type BanksListFilter struct {
	Id                  *uint       `json:"id,omitempty" v:"min:1" dc:"ID"`
	IdArr               []uint      `json:"idArr,omitempty" v:"distinct|foreach|min:1" dc:"ID数组"`
	ExcId               *uint       `json:"excId,omitempty" v:"min:1" dc:"排除ID"`
	ExcIdArr            []uint      `json:"excIdArr,omitempty" v:"distinct|foreach|min:1" dc:"排除ID数组"`
	Label               string      `json:"label,omitempty" v:"max-length:30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"标签。常用于前端组件"`
	Name                string      `json:"name,omitempty" v:"max-length:50|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"银行名称"`
	Slug                string      `json:"slug,omitempty" v:"max-length:50" dc:"slug"`
	Code                string      `json:"code,omitempty" v:"max-length:50|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"code"`
	FlutterwaveBankCode string      `json:"flutterwaveBankCode,omitempty" v:"max-length:50|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"flutterwaveBankCode"`
	RedbillerBankCode   string      `json:"redbillerBankCode,omitempty" v:"max-length:50|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"redbillerBankCode"`
	AnchorBankCode      string      `json:"anchorBankCode,omitempty" v:"max-length:50|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"anchorBankCode"`
	Country             string      `json:"country,omitempty" v:"max-length:50" dc:"国家"`
	Currency            string      `json:"currency,omitempty" v:"max-length:50" dc:"货币"`
	Type                string      `json:"type,omitempty" v:"max-length:50" dc:"类型"`
	IsStop              *uint       `json:"isStop,omitempty" v:"in:0,1" dc:"停用：0否 1是"`
	TimeRangeStart      *gtime.Time `json:"timeRangeStart,omitempty" v:"date-format:Y-m-d H:i:s" dc:"开始时间：YYYY-mm-dd HH:ii:ss"`
	TimeRangeEnd        *gtime.Time `json:"timeRangeEnd,omitempty" v:"date-format:Y-m-d H:i:s|after-equal:TimeRangeStart" dc:"结束时间：YYYY-mm-dd HH:ii:ss"`
}

type BanksListRes struct {
	Count int             `json:"count" dc:"总数"`
	List  []BanksListItem `json:"list" dc:"列表"`
}

type BanksListItem struct {
	Id                  *uint       `json:"id,omitempty" dc:"ID"`
	Label               *string     `json:"label,omitempty" dc:"标签。常用于前端组件"`
	Name                *string     `json:"name,omitempty" dc:"银行名称"`
	Slug                *string     `json:"slug,omitempty" dc:"slug"`
	Code                *string     `json:"code,omitempty" dc:"code"`
	FlutterwaveBankCode *string     `json:"flutterwaveBankCode,omitempty" dc:"flutterwaveBankCode"`
	RedbillerBankCode   *string     `json:"redbillerBankCode,omitempty" dc:"redbillerBankCode"`
	AnchorBankCode      *string     `json:"anchorBankCode,omitempty" dc:"anchorBankCode"`
	Country             *string     `json:"country,omitempty" dc:"国家"`
	Currency            *string     `json:"currency,omitempty" dc:"货币"`
	Type                *string     `json:"type,omitempty" dc:"类型"`
	IsStop              *uint       `json:"isStop,omitempty" dc:"停用：0否 1是"`
	UpdatedAt           *gtime.Time `json:"updatedAt,omitempty" dc:"更新时间"`
	CreatedAt           *gtime.Time `json:"createdAt,omitempty" dc:"创建时间"`
}

/*--------列表 结束--------*/

/*--------详情 开始--------*/
type BanksInfoReq struct {
	g.Meta `path:"/banks/info" method:"post" tags:"平台后台/银行" sm:"详情"`
	Id     uint     `json:"id" v:"required|min:1" dc:"ID"`
	Field  []string `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回全部字段。注意：如前端页面所需字段较少，建议传指定字段，可大幅减轻服务器及数据库压力"`
}

type BanksInfoRes struct {
	Info BanksInfo `json:"info" dc:"详情"`
}

type BanksInfo struct {
	Id                  *uint       `json:"id,omitempty" dc:"ID"`
	Label               *string     `json:"label,omitempty" dc:"标签。常用于前端组件"`
	Name                *string     `json:"name,omitempty" dc:"银行名称"`
	Slug                *string     `json:"slug,omitempty" dc:"slug"`
	Code                *string     `json:"code,omitempty" dc:"code"`
	FlutterwaveBankCode *string     `json:"flutterwaveBankCode,omitempty" dc:"flutterwaveBankCode"`
	RedbillerBankCode   *string     `json:"redbillerBankCode,omitempty" dc:"redbillerBankCode"`
	AnchorBankCode      *string     `json:"anchorBankCode,omitempty" dc:"anchorBankCode"`
	Country             *string     `json:"country,omitempty" dc:"国家"`
	Currency            *string     `json:"currency,omitempty" dc:"货币"`
	Type                *string     `json:"type,omitempty" dc:"类型"`
	IsStop              *uint       `json:"isStop,omitempty" dc:"停用：0否 1是"`
	UpdatedAt           *gtime.Time `json:"updatedAt,omitempty" dc:"更新时间"`
	CreatedAt           *gtime.Time `json:"createdAt,omitempty" dc:"创建时间"`
}

/*--------详情 结束--------*/

/*--------新增 开始--------*/
type BanksCreateReq struct {
	g.Meta              `path:"/banks/create" method:"post" tags:"平台后台/银行" sm:"创建"`
	Name                *string `json:"name,omitempty" v:"required|max-length:50|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"银行名称"`
	Slug                *string `json:"slug,omitempty" v:"max-length:50" dc:"slug"`
	Code                *string `json:"code,omitempty" v:"max-length:50|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"code"`
	FlutterwaveBankCode *string `json:"flutterwaveBankCode,omitempty" v:"max-length:50|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"flutterwaveBankCode"`
	RedbillerBankCode   *string `json:"redbillerBankCode,omitempty" v:"max-length:50|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"redbillerBankCode"`
	AnchorBankCode      *string `json:"anchorBankCode,omitempty" v:"max-length:50|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"anchorBankCode"`
	Country             *string `json:"country,omitempty" v:"max-length:50" dc:"国家"`
	Currency            *string `json:"currency,omitempty" v:"max-length:50" dc:"货币"`
	Type                *string `json:"type,omitempty" v:"max-length:50" dc:"类型"`
	IsStop              *uint   `json:"isStop,omitempty" v:"in:0,1" dc:"停用：0否 1是"`
}

/*--------新增 结束--------*/

/*--------修改 开始--------*/
type BanksUpdateReq struct {
	g.Meta              `path:"/banks/update" method:"post" tags:"平台后台/银行" sm:"更新"`
	IdArr               []uint  `json:"idArr,omitempty" v:"required|distinct|foreach|min:1" dc:"ID数组"`
	Name                *string `json:"name,omitempty" v:"max-length:50|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"银行名称"`
	Slug                *string `json:"slug,omitempty" v:"max-length:50" dc:"slug"`
	Code                *string `json:"code,omitempty" v:"max-length:50|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"code"`
	FlutterwaveBankCode *string `json:"flutterwaveBankCode,omitempty" v:"max-length:50|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"flutterwaveBankCode"`
	RedbillerBankCode   *string `json:"redbillerBankCode,omitempty" v:"max-length:50|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"redbillerBankCode"`
	AnchorBankCode      *string `json:"anchorBankCode,omitempty" v:"max-length:50|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"anchorBankCode"`
	Country             *string `json:"country,omitempty" v:"max-length:50" dc:"国家"`
	Currency            *string `json:"currency,omitempty" v:"max-length:50" dc:"货币"`
	Type                *string `json:"type,omitempty" v:"max-length:50" dc:"类型"`
	IsStop              *uint   `json:"isStop,omitempty" v:"in:0,1" dc:"停用：0否 1是"`
}

/*--------修改 结束--------*/

/*--------删除 开始--------*/
type BanksDeleteReq struct {
	g.Meta `path:"/banks/del" method:"post" tags:"平台后台/银行" sm:"删除"`
	IdArr  []uint `json:"idArr,omitempty" v:"required|distinct|foreach|min:1" dc:"ID数组"`
}

/*--------删除 结束--------*/
