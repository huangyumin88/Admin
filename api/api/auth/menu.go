package api

import (
	apiCommon "api/api"
)

type MenuListReq struct {
	Filter MenuListFilterReq `p:"filter"`
	apiCommon.CommonListReq
}

type MenuListFilterReq struct {
	apiCommon.CommonListFilterReq `c:",omitempty"`
	// 下面根据自己需求修改
	IsStop   *uint  `c:"isStop,omitempty" p:"isStop" v:"in:0,1"`
	MenuId   *uint  `c:"menuId,omitempty" p:"menuId" v:"min:1"`
	MenuName string `c:"menuName,omitempty" p:"menuName" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	/* type Menu struct {
	    MenuId     uint        `json:"menuId"     `// 权限菜单ID
	    SceneId    uint        `json:"sceneId"    `// 权限场景ID（只能是auth_scene表中sceneType为0的菜单类型场景）
	    Pid        uint        `json:"pid"        `// 父ID
	    MenuName   string      `json:"menuName"   `// 名称
	    MenuIcon   string      `json:"menuIcon"   `// 图标
	    MenuUrl    string      `json:"menuUrl"    `// 链接
	    Level      uint        `json:"level"      `// 层级
	    PidPath    string      `json:"pidPath"    `// 层级路径
	    ExtraData  string      `json:"extraData"  `// 额外数据。（json格式：{"i18n（国际化设置）": {"title": {"语言标识":"标题",...}}）
	    Sort       uint        `json:"sort"       `// 排序值（从小到大排序，默认50，范围0-100）
	    IsStop     uint        `json:"isStop"     `// 是否停用：0否 1是
	    UpdateTime *gtime.Time `json:"updateTime" `// 更新时间
	    CreateTime *gtime.Time `json:"createTime" `// 创建时间
	} */
}
