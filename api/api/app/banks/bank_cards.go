package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

/*--------列表 开始--------*/
type BankCardsListReq struct {
	g.Meta `path:"/bankCards/list" method:"post" tags:"APP/用户银行" sm:"列表"`
	Filter BankCardsListFilter `json:"filter" dc:"查询条件"`
	Field  []string            `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回全部字段。注意：如前端页面所需字段较少，建议传指定字段，可大幅减轻服务器及数据库压力"`
	Sort   string              `json:"sort" default:"id DESC" dc:"排序"`
	Page   int                 `json:"page" v:"min:1" default:"1" dc:"页码"`
	Limit  int                 `json:"limit" v:"min:0" default:"10" dc:"每页数量。可传0取全部"`
}

type BankCardsListFilter struct {
	Id             *uint       `json:"id,omitempty" v:"min:1" dc:"ID"`
	IdArr          []uint      `json:"idArr,omitempty" v:"distinct|foreach|min:1" dc:"ID数组"`
	ExcId          *uint       `json:"excId,omitempty" v:"min:1" dc:"排除ID"`
	ExcIdArr       []uint      `json:"excIdArr,omitempty" v:"distinct|foreach|min:1" dc:"排除ID数组"`
	UserId         *uint       `json:"user_id,omitempty" v:"min:1" dc:"用户id"`
	BankId         *int        `json:"bank_id,omitempty" v:"min:1" dc:"银行id"`
	CardNumber     string      `json:"card_number,omitempty" v:"max-length:255" dc:"银行卡号"`
	CardHolderName string      `json:"card_holder_name,omitempty" v:"max-length:255|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"持卡人姓名"`
	ExpirationDate *gtime.Time `json:"expiration_date,omitempty" v:"date-format:Y-m-d" dc:"银行卡的有效期"`
	Cvv            string      `json:"cvv,omitempty" v:"max-length:255" dc:"CVV 安全码"`
	TimeRangeStart *gtime.Time `json:"timeRangeStart,omitempty" v:"date-format:Y-m-d H:i:s" dc:"开始时间：YYYY-mm-dd HH:ii:ss"`
	TimeRangeEnd   *gtime.Time `json:"timeRangeEnd,omitempty" v:"date-format:Y-m-d H:i:s|after-equal:TimeRangeStart" dc:"结束时间：YYYY-mm-dd HH:ii:ss"`
}

type BankCardsListRes struct {
	Count int                 `json:"count" dc:"总数"`
	List  []BankCardsListItem `json:"list" dc:"列表"`
}

type BankCardsListItem struct {
	Id             *uint       `json:"id,omitempty" dc:"ID"`
	UserId         *uint       `json:"user_id,omitempty" dc:"用户id"`
	BankId         *int        `json:"bank_id,omitempty" dc:"银行id"`
	CardNumber     *string     `json:"card_number,omitempty" dc:"银行卡号"`
	CardHolderName *string     `json:"card_holder_name,omitempty" dc:"持卡人姓名"`
	ExpirationDate *string     `json:"expiration_date,omitempty" dc:"银行卡的有效期"`
	Cvv            *string     `json:"cvv,omitempty" dc:"CVV 安全码"`
	UpdatedAt      *gtime.Time `json:"updatedAt,omitempty" dc:"更新时间"`
	CreatedAt      *gtime.Time `json:"createdAt,omitempty" dc:"创建时间"`
	UserName       *string     `json:"user_name,omitempty" dc:"用户"`
}

/*--------列表 结束--------*/

/*--------详情 开始--------*/
type BankCardsInfoReq struct {
	g.Meta `path:"/bankCards/info" method:"post" tags:"APP/用户银行" sm:"详情"`
	Id     uint     `json:"id" v:"required|min:1" dc:"ID"`
	Field  []string `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回全部字段。注意：如前端页面所需字段较少，建议传指定字段，可大幅减轻服务器及数据库压力"`
}

type BankCardsInfoRes struct {
	Info BankCardsInfo `json:"info" dc:"详情"`
}

type BankCardsInfo struct {
	Id             *uint       `json:"id,omitempty" dc:"ID"`
	UserId         *uint       `json:"user_id,omitempty" dc:"用户id"`
	BankId         *int        `json:"bank_id,omitempty" dc:"银行id"`
	CardNumber     *string     `json:"card_number,omitempty" dc:"银行卡号"`
	CardHolderName *string     `json:"card_holder_name,omitempty" dc:"持卡人姓名"`
	ExpirationDate *string     `json:"expiration_date,omitempty" dc:"银行卡的有效期"`
	Cvv            *string     `json:"cvv,omitempty" dc:"CVV 安全码"`
	UpdatedAt      *gtime.Time `json:"updatedAt,omitempty" dc:"更新时间"`
	CreatedAt      *gtime.Time `json:"createdAt,omitempty" dc:"创建时间"`
}

/*--------详情 结束--------*/

/*--------新增 开始--------*/
type BankCardsCreateReq struct {
	g.Meta         `path:"/bankCards/create" method:"post" tags:"APP/用户银行" sm:"创建"`
	UserId         *uint   `json:"user_id,omitempty" v:"min:1" dc:"用户id"`
	BankId         *int    `json:"bank_id,omitempty" v:"min:1" dc:"银行id"`
	CardNumber     *string `json:"card_number,omitempty" v:"max-length:50" dc:"银行卡号"`
	CardHolderName *string `json:"card_holder_name,omitempty" v:"max-length:50" dc:"持卡人姓名"`
	ExpirationDate *string `json:"expiration_date,omitempty" v:"" dc:"银行卡的有效期"`
	Cvv            *string `json:"cvv,omitempty" v:"" dc:"CVV 安全码"`
}

/*--------新增 结束--------*/

/*--------修改 开始--------*/
type BankCardsUpdateReq struct {
	g.Meta `path:"/bankCards/update" method:"post" tags:"APP/用户银行" sm:"更新"`
	IdArr  []uint `json:"idArr,omitempty" v:"required|distinct|foreach|min:1" dc:"ID数组"`
	//UserId         *uint   `json:"user_id,omitempty" v:"min:1" dc:"用户id"`
	//BankId         *int    `json:"bank_id,omitempty" v:"min:1" dc:"银行id"`
	CardNumber     *string `json:"card_number,omitempty" v:"max-length:50" dc:"银行卡号"`
	CardHolderName *string `json:"card_holder_name,omitempty" v:"max-length:50" dc:"持卡人姓名"`
	ExpirationDate *string `json:"expiration_date,omitempty" v:"" dc:"银行卡的有效期"`
	Cvv            *string `json:"cvv,omitempty" v:"max-length:255" dc:"CVV 安全码"`
}

/*--------修改 结束--------*/

/*--------删除 开始--------*/
type BankCardsDeleteReq struct {
	g.Meta `path:"/bankCards/del" method:"post" tags:"APP/用户银行" sm:"删除"`
	IdArr  []uint `json:"idArr,omitempty" v:"required|distinct|foreach|min:1" dc:"ID数组"`
}

/*--------删除 结束--------*/
