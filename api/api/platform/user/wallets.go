package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

/*--------列表 开始--------*/
type WalletsListReq struct {
	g.Meta `path:"/wallets/list" method:"post" tags:"平台后台/钱包" sm:"列表"`
	Filter WalletsListFilter `json:"filter" dc:"查询条件"`
	Field  []string          `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回全部字段。注意：如前端页面所需字段较少，建议传指定字段，可大幅减轻服务器及数据库压力"`
	Sort   string            `json:"sort" default:"id DESC" dc:"排序"`
	Page   int               `json:"page" v:"min:1" default:"1" dc:"页码"`
	Limit  int               `json:"limit" v:"min:0" default:"10" dc:"每页数量。可传0取全部"`
}

type WalletsListFilter struct {
	Id             *uint       `json:"id,omitempty" v:"min:1" dc:"ID"`
	IdArr          []uint      `json:"idArr,omitempty" v:"distinct|foreach|min:1" dc:"ID数组"`
	ExcId          *uint       `json:"excId,omitempty" v:"min:1" dc:"排除ID"`
	ExcIdArr       []uint      `json:"excIdArr,omitempty" v:"distinct|foreach|min:1" dc:"排除ID数组"`
	WalletId       *uint       `json:"wallet_id,omitempty" v:"min:1" dc:"钱包ID"`
	UserId         *uint       `json:"user_id,omitempty" v:"min:1" dc:"用户ID"`
	Currency       string      `json:"currency,omitempty" v:"max-length:10" dc:"货币 默认尼日利亚 奈拉"`
	IsStop         *uint       `json:"isStop,omitempty" v:"in:0,1" dc:"停用：0否 1是"`
	TimeRangeStart *gtime.Time `json:"timeRangeStart,omitempty" v:"date-format:Y-m-d H:i:s" dc:"开始时间：YYYY-mm-dd HH:ii:ss"`
	TimeRangeEnd   *gtime.Time `json:"timeRangeEnd,omitempty" v:"date-format:Y-m-d H:i:s|after-equal:TimeRangeStart" dc:"结束时间：YYYY-mm-dd HH:ii:ss"`
}

type WalletsListRes struct {
	Count int               `json:"count" dc:"总数"`
	List  []WalletsListItem `json:"list" dc:"列表"`
}

type WalletsListItem struct {
	Id           *uint       `json:"id,omitempty" dc:"ID"`
	WalletId     *uint       `json:"wallet_id,omitempty" dc:"钱包ID"`
	UserId       *uint       `json:"user_id,omitempty" dc:"用户ID"`
	Balance      *float64    `json:"balance,omitempty" dc:"余额"`
	RewardPoints *uint       `json:"reward_points,omitempty" dc:"积分"`
	Currency     *string     `json:"currency,omitempty" dc:"货币 默认尼日利亚 奈拉"`
	IsStop       *uint       `json:"isStop,omitempty" dc:"停用：0否 1是"`
	UpdatedAt    *gtime.Time `json:"updatedAt,omitempty" dc:"更新时间"`
	CreatedAt    *gtime.Time `json:"createdAt,omitempty" dc:"创建时间"`
	UserName     *string     `json:"user_name,omitempty" dc:"用户"`
	Info         *UserInfo   `json:"info" dc:"详情"`
}

/*--------列表 结束--------*/

/*--------详情 开始--------*/
type WalletsInfoReq struct {
	g.Meta `path:"/wallets/info" method:"post" tags:"平台后台/钱包" sm:"详情"`
	Id     uint     `json:"id" v:"required|min:1" dc:"ID"`
	Field  []string `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回全部字段。注意：如前端页面所需字段较少，建议传指定字段，可大幅减轻服务器及数据库压力"`
}

type WalletsInfoRes struct {
	Info WalletsInfo `json:"info" dc:"详情"`
}

type WalletsInfo struct {
	Id           *uint       `json:"id,omitempty" dc:"ID"`
	WalletId     *uint       `json:"wallet_id,omitempty" dc:"钱包ID"`
	UserId       *uint       `json:"user_id,omitempty" dc:"用户ID"`
	Balance      *float64    `json:"balance,omitempty" dc:"余额"`
	RewardPoints *uint       `json:"reward_points,omitempty" dc:"积分"`
	Currency     *string     `json:"currency,omitempty" dc:"货币 默认尼日利亚 奈拉"`
	IsStop       *uint       `json:"isStop,omitempty" dc:"停用：0否 1是"`
	UpdatedAt    *gtime.Time `json:"updatedAt,omitempty" dc:"更新时间"`
	CreatedAt    *gtime.Time `json:"createdAt,omitempty" dc:"创建时间"`
}

/*--------详情 结束--------*/

/*--------修改 开始--------*/
type WalletsUpdateReq struct {
	g.Meta       `path:"/wallets/update" method:"post" tags:"平台后台/钱包" sm:"更新"`
	IdArr        []uint   `json:"idArr,omitempty" v:"required|distinct|foreach|min:1" dc:"ID数组"`
	UserId       *uint    `json:"user_id,omitempty" v:"min:1" dc:"用户ID"`
	Balance      *float64 `json:"balance,omitempty" v:"" dc:"余额"`
	RewardPoints *uint    `json:"reward_points,omitempty" v:"" dc:"积分"`
	Currency     *string  `json:"currency,omitempty" v:"max-length:10" dc:"货币 默认尼日利亚 奈拉"`
	IsStop       *uint    `json:"isStop,omitempty" v:"in:0,1" dc:"停用：0否 1是"`
}

/*--------修改 结束--------*/
