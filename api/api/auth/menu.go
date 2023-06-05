package api

import (
	apiCommon "api/api"
)

type MenuListReq struct {
	apiCommon.CommonListReq
	Filter MenuListFilterReq `p:"filter"`
}

type MenuListFilterReq struct {
	apiCommon.CommonListFilterReq `c:",omitempty"`
	MenuId                        *uint  `c:"menuId,omitempty" p:"menuId" v:"min:1"`
	SceneId                       *uint  `c:"sceneId,omitempty" p:"sceneId" v:"min:1"`
	Pid                           *uint  `c:"pid,omitempty" p:"pid" v:"min:0"`
	MenuName                      string `c:"menuName,omitempty" p:"menuName" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	IsStop                        *uint  `c:"isStop,omitempty" p:"isStop" v:"in:0,1"`
}

type MenuInfoReq struct {
	apiCommon.CommonInfoReq
}

type MenuCreateReq struct {
	SceneId   *uint   `c:"sceneId,omitempty" p:"sceneId" v:"required|min:1"`
	Pid       *uint   `c:"sceneId,omitempty" p:"sceneId" v:"min:0"`
	MenuName  *string `c:"menuName,omitempty" p:"menuName" v:"required|length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	MenuIcon  *string `c:"menuIcon,omitempty" p:"menuIcon" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	MenuUrl   *string `c:"menuUrl,omitempty" p:"menuUrl" v:"length:1,120"`
	ExtraData *string `c:"extraData,omitempty" p:"extraData" v:"json"`
	Sort      *uint   `c:"sort,omitempty" p:"sort" v:"between:0,100"`
	IsStop    *uint   `c:"isStop,omitempty" p:"isStop" v:"in:0,1"`
}

type MenuUpdateReq struct {
	apiCommon.CommonUpdateDeleteIdArrReq `c:",omitempty"`
	SceneId                              *uint   `c:"sceneId,omitempty" p:"sceneId" v:"min:1"`
	Pid                                  *uint   `c:"sceneId,omitempty" p:"sceneId" v:"min:0"`
	MenuName                             *string `c:"menuName,omitempty" p:"menuName" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	MenuIcon                             *string `c:"menuIcon,omitempty" p:"menuIcon" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	MenuUrl                              *string `c:"menuUrl,omitempty" p:"menuUrl" v:"length:1,120"`
	ExtraData                            *string `c:"extraData,omitempty" p:"extraData" v:"json"`
	Sort                                 *uint   `c:"sort,omitempty" p:"sort" v:"between:0,100"`
	IsStop                               *uint   `c:"isStop,omitempty" p:"isStop" v:"in:0,1"`
}

type MenuDeleteReq struct {
	apiCommon.CommonUpdateDeleteIdArrReq
}