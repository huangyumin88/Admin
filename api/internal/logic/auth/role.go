package logic

import (
	daoAuth "api/internal/dao/auth"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sAuthRole struct{}

func NewAuthRole() *sAuthRole {
	return &sAuthRole{}
}

func init() {
	service.RegisterAuthRole(NewAuthRole())
}

// 新增
func (logicThis *sAuthRole) Create(ctx context.Context, data map[string]interface{}) (id int64, err error) {
	daoThis := daoAuth.Role
	_, okMenuIdArr := data[`menuIdArr`]
	if okMenuIdArr {
		menuIdArr := gconv.SliceUint(data[`menuIdArr`])
		filterTmp := g.Map{daoAuth.Menu.PrimaryKey(): menuIdArr, daoAuth.Menu.Columns().SceneId: data[`sceneId`]}
		count, _ := daoAuth.Menu.ParseDbCtx(ctx).Where(filterTmp).Count()
		if len(menuIdArr) != count {
			err = utils.NewErrorCode(ctx, 89999998, ``)
			return
		}
	}
	_, okActionIdArr := data[`actionIdArr`]
	if okActionIdArr {
		actionIdArr := gconv.SliceUint(data[`actionIdArr`])
		filterTmp := g.Map{daoAuth.ActionRelToScene.Columns().ActionId: actionIdArr, daoAuth.ActionRelToScene.Columns().SceneId: data[`sceneId`]}
		count, _ := daoAuth.ActionRelToScene.ParseDbCtx(ctx).Where(filterTmp).Count()
		if len(actionIdArr) != count {
			err = utils.NewErrorCode(ctx, 89999998, ``)
			return
		}
	}

	id, err = daoThis.HandlerCtx(ctx).Insert(data).GetModel().InsertAndGetId()
	return
}

// 修改
func (logicThis *sAuthRole) Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (row int64, err error) {
	daoThis := daoAuth.Role
	daoHandlerThis := daoThis.HandlerCtx(ctx).Filter(filter, true)
	if len(daoHandlerThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	_, okMenuIdArr := data[`menuIdArr`]
	if okMenuIdArr {
		menuIdArr := gconv.SliceUint(data[`menuIdArr`])
		for _, id := range daoHandlerThis.IdArr {
			oldInfo, _ := daoThis.ParseDbCtx(ctx).Where(daoThis.PrimaryKey(), id).One()
			filterTmp := g.Map{daoAuth.Menu.PrimaryKey(): menuIdArr, daoAuth.Menu.Columns().SceneId: oldInfo[`sceneId`]}
			_, okSceneId := data[`sceneId`]
			if okSceneId {
				filterTmp[daoAuth.Menu.Columns().SceneId] = data[`sceneId`]
			}
			count, _ := daoAuth.Menu.ParseDbCtx(ctx).Where(filterTmp).Count()
			if len(menuIdArr) != count {
				err = utils.NewErrorCode(ctx, 89999998, ``)
				return
			}
		}
	}
	_, okActionIdArr := data[`actionIdArr`]
	if okActionIdArr {
		actionIdArr := gconv.SliceUint(data[`actionIdArr`])
		for _, id := range daoHandlerThis.IdArr {
			oldInfo, _ := daoThis.ParseDbCtx(ctx).Where(daoThis.PrimaryKey(), id).One()
			filterTmp := g.Map{daoAuth.ActionRelToScene.Columns().ActionId: actionIdArr, daoAuth.ActionRelToScene.Columns().SceneId: oldInfo[`sceneId`]}
			_, okSceneId := data[`sceneId`]
			if okSceneId {
				filterTmp[daoAuth.ActionRelToScene.Columns().SceneId] = data[`sceneId`]
			}
			count, _ := daoAuth.ActionRelToScene.ParseDbCtx(ctx).Where(filterTmp).Count()
			if len(actionIdArr) != count {
				err = utils.NewErrorCode(ctx, 89999998, ``)
				return
			}
		}
	}

	row, err = daoHandlerThis.Update(data).GetModel().UpdateAndGetAffected()
	return
}

// 删除
func (logicThis *sAuthRole) Delete(ctx context.Context, filter map[string]interface{}) (row int64, err error) {
	daoThis := daoAuth.Role
	daoHandlerThis := daoThis.HandlerCtx(ctx).Filter(filter, true)
	if len(daoHandlerThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	result, err := daoHandlerThis.Delete().GetModel().Delete()
	row, _ = result.RowsAffected()
	return
}
