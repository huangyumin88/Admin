package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

/*--------列表 开始--------*/
type SceneListReq struct {
	g.Meta `path:"/list" method:"post" tags:"平台-场景" sm:"列表"`
	Filter SceneListFilter `json:"filter" dc:"查询条件"`
	// apiCommon.CommonListReq
	Field []string `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段。默认会返回全部查询字段。如果需要的字段较少，建议指定字段，传值参考默认返回的字段"`
	Sort  string   `json:"sort" default:"id DESC" dc:"排序"`
	Page  int      `json:"page" v:"integer|min:1" default:"1" dc:"页码"`
	Limit int      `json:"limit" v:"integer|min:0" default:"10" dc:"每页数量。可传0取全部"`
}

type SceneListFilter struct {
	/*--------公共参数 开始--------*/
	// apiCommon.CommonListFilterReq `c:",omitempty"`	// 代码中用到转换成map，且必须用omitempty过滤空参数。而规范路由自动生成swagger会因omitempty导致这些字段不生成。故直接写这里
	Id        *uint       `c:"id,omitempty" json:"id" v:"integer|min:1" dc:"ID"`
	IdArr     []uint      `c:"idArr,omitempty" json:"idArr" v:"distinct|foreach|integer|foreach|min:1" dc:"ID数组"`
	ExcId     *uint       `c:"excId,omitempty" json:"excId" v:"integer|min:1" dc:"排除ID"`
	ExcIdArr  []uint      `c:"excIdArr,omitempty" json:"excIdArr" v:"distinct|foreach|integer|foreach|min:1" dc:"排除ID数组"`
	StartTime *gtime.Time `c:"startTime,omitempty" json:"startTime" v:"date-format:Y-m-d H:i:s" dc:"开始时间。示例：2000-01-01 00:00:00"`
	EndTime   *gtime.Time `c:"endTime,omitempty" json:"endTime" v:"date-format:Y-m-d H:i:s|after-equal:StartTime" dc:"结束时间。示例：2000-01-01 00:00:00"`
	Name      string      `c:"name,omitempty" json:"name" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"名称。后台公共列表常用"`
	/*--------公共参数 结束--------*/
	SceneId   *uint  `c:"sceneId,omitempty" json:"sceneId" v:"integer|min:1" dc:"场景ID"`
	SceneCode string `c:"sceneCode,omitempty" json:"sceneCode" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"场景Code"`
	SceneName string `c:"sceneName,omitempty" json:"sceneName" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"场景名称"`
	IsStop    *uint  `c:"isStop,omitempty" json:"isStop" v:"integer|in:0,1" dc:"是否停用：0否 1是"`
}

type SceneListRes struct {
	// apiCommon.CommonListRes
	Count int         `json:"count" dc:"总数"`
	List  []SceneList `json:"list" dc:"列表"`
}

type SceneList struct {
	Id          uint   `json:"id" dc:"ID"`
	Name        string `json:"name" dc:"名称"`
	SceneId     uint   `json:"sceneId" dc:"场景ID"`
	SceneCode   string `json:"sceneCode" dc:"场景标识"`
	SceneName   string `json:"sceneName" dc:"场景名称"`
	SceneConfig string `json:"sceneConfig" dc:"场景配置"`
	IsStop      uint   `json:"isStop" dc:"是否停用：0否 1是"`
	UpdatedAt   string `json:"updatedAt" dc:"更新时间"`
	CreatedAt   string `json:"createdAt" dc:"创建时间"`
}

/*--------列表 结束--------*/

/*--------详情 开始--------*/
type SceneInfoReq struct {
	g.Meta `path:"/info" method:"post" tags:"平台-场景" sm:"详情"`
	// apiCommon.CommonInfoReq
	Id    uint     `json:"id" v:"required|integer|min:1" dc:"ID"`
	Field []string `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段。默认会返回全部查询字段。如果需要的字段较少，建议指定字段，传值参考默认返回的字段"`
}

type SceneInfoRes struct {
	Info SceneInfo `json:"info" dc:"详情"`
}

type SceneInfo struct {
	Id          uint   `json:"id" dc:"ID"`
	Name        string `json:"name" dc:"名称"`
	SceneId     uint   `json:"sceneId" dc:"场景ID"`
	SceneCode   string `json:"sceneCode" dc:"场景标识"`
	SceneName   string `json:"sceneName" dc:"场景名称"`
	SceneConfig string `json:"sceneConfig" dc:"场景配置"`
	IsStop      uint   `json:"isStop" dc:"是否停用：0否 1是"`
	UpdatedAt   string `json:"updatedAt" dc:"更新时间"`
	CreatedAt   string `json:"createdAt" dc:"创建时间"`
}

/*--------详情 结束--------*/

/*--------新增 开始--------*/
type SceneCreateReq struct {
	g.Meta      `path:"/create" method:"post" tags:"平台-场景" sm:"创建"`
	SceneCode   *string `c:"sceneCode,omitempty" json:"sceneCode" v:"required|length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"场景标识"`
	SceneName   *string `c:"sceneName,omitempty" json:"sceneName" v:"required|length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"场景名称"`
	SceneConfig *string `c:"sceneConfig,omitempty" json:"sceneConfig" v:"json" dc:"场景配置"`
	IsStop      *uint   `c:"isStop,omitempty" json:"isStop" v:"integer|in:0,1" dc:"是否停用：0否 1是"`
}

type SceneCreateRes struct {
	Id int64 `json:"id" dc:"ID"`
}

/*--------新增 结束--------*/

/*--------修改 开始--------*/
type SceneUpdateReq struct {
	g.Meta `path:"/update" method:"post" tags:"平台-场景" sm:"更新"`
	// apiCommon.CommonUpdateDeleteIdArrReq `c:",omitempty"`
	IdArr       []uint  `c:"idArr,omitempty" json:"idArr" v:"required|distinct|foreach|integer|foreach|min:1" dc:"ID数组"`
	SceneCode   *string `c:"sceneCode,omitempty" json:"sceneCode" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"场景标识"`
	SceneName   *string `c:"sceneName,omitempty" json:"sceneName" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"场景名称"`
	SceneConfig *string `c:"sceneConfig,omitempty" json:"sceneConfig" v:"json" dc:"场景配置"`
	IsStop      *uint   `c:"isStop,omitempty" json:"isStop" v:"integer|in:0,1" dc:"是否停用：0否 1是"`
}

type SceneUpdateRes struct {
}

/*--------修改 结束--------*/

/*--------删除 开始--------*/
type SceneDeleteReq struct {
	g.Meta `path:"/del" method:"post" tags:"平台-场景" sm:"删除"`
	// apiCommon.CommonUpdateDeleteIdArrReq
	IdArr []uint `c:"idArr,omitempty" json:"idArr" v:"required|distinct|foreach|integer|foreach|min:1" dc:"ID数组"`
}

type SceneDeleteRes struct {
}

/*--------删除 结束--------*/