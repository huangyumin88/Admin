package controller

import (
	apiAuth "api/api/auth"
	daoAuth "api/internal/model/dao/auth"
	"api/internal/service"
	"api/internal/utils"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

type Menu struct{}

func NewMenu() *Menu {
	return &Menu{}
}

// 列表
func (controllerThis *Menu) List(r *ghttp.Request) {
	/**--------参数处理 开始--------**/
	var param *apiAuth.MenuListReq
	err := r.Parse(&param)
	if err != nil {
		utils.HttpFailJson(r, err)
		return
	}
	filter := gconv.Map(param.Filter)
	order := [][2]string{{"id", "DESC"}}
	if param.Sort.Key != "" {
		order[0][0] = param.Sort.Key
	}
	if param.Sort.Order != "" {
		order[0][1] = param.Sort.Order
	}
	if param.Page <= 0 {
		param.Page = 1
	}
	if param.Limit <= 0 {
		param.Limit = 10
	}
	/**--------参数处理 结束--------**/

	sceneCode := utils.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	case "platformAdmin":
		/**--------权限验证 开始--------**/
		isAuth, _ := service.Action().CheckAuth(r.Context(), "authMenuLook")
		allowField := []string{"menuId", "menuName", "id"}
		if isAuth {
			allowField = daoAuth.Menu.ColumnArr()
			allowField = append(allowField, "id", "sceneName", "pMenuName")
			//allowField = gset.NewStrSetFrom(allowField).Diff(gset.NewStrSetFrom([]string{"password"})).Slice() //移除敏感字段
		}
		field := allowField
		if len(param.Field) > 0 {
			field = gset.NewStrSetFrom(param.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
			if len(field) == 0 {
				field = allowField
			}
		}
		/**--------权限验证 结束--------**/

		count, err := service.Menu().Count(r.Context(), filter)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		list, err := service.Menu().List(r.Context(), filter, field, order, param.Page, param.Limit)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		utils.HttpSuccessJson(r, map[string]interface{}{"count": count, "list": list}, 0)
	}
}

// 详情
func (controllerThis *Menu) Info(r *ghttp.Request) {
	sceneCode := utils.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	case "platformAdmin":
		/**--------参数处理 开始--------**/
		var param *apiAuth.MenuInfoReq
		err := r.Parse(&param)
		if err != nil {
			utils.HttpFailJson(r, utils.NewErrorCode(r.GetCtx(), 89999999, err.Error()))
			return
		}

		allowField := daoAuth.Menu.ColumnArr()
		allowField = append(allowField, "id")
		//allowField = gset.NewStrSetFrom(allowField).Diff(gset.NewStrSetFrom([]string{"password"})).Slice() //移除敏感字段
		field := allowField
		if len(param.Field) > 0 {
			field = gset.NewStrSetFrom(param.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
			if len(field) == 0 {
				field = allowField
			}
		}
		filter := map[string]interface{}{"id": param.Id}
		/**--------参数处理 结束--------**/

		/**--------权限验证 开始--------**/
		_, err = service.Action().CheckAuth(r.Context(), "authMenuLook")
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		/**--------权限验证 结束--------**/

		info, err := service.Menu().Info(r.Context(), filter, field)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		utils.HttpSuccessJson(r, map[string]interface{}{"info": info}, 0)
	}
}

// 创建
func (controllerThis *Menu) Create(r *ghttp.Request) {
	sceneCode := utils.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	case "platformAdmin":
		/**--------参数处理 开始--------**/
		var param *apiAuth.MenuCreateReq
		err := r.Parse(&param)
		if err != nil {
			utils.HttpFailJson(r, utils.NewErrorCode(r.GetCtx(), 89999999, err.Error()))
			return
		}
		data := gconv.Map(param)
		/**--------参数处理 结束--------**/

		/**--------权限验证 开始--------**/
		_, err = service.Action().CheckAuth(r.Context(), "authMenuCreate")
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		/**--------权限验证 结束--------**/

		_, err = service.Menu().Create(r.Context(), data)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		utils.HttpSuccessJson(r, map[string]interface{}{}, 0)
	}
}

// 更新
func (controllerThis *Menu) Update(r *ghttp.Request) {
	sceneCode := utils.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	case "platformAdmin":
		/**--------参数处理 开始--------**/
		var param *apiAuth.MenuUpdateReq
		err := r.Parse(&param)
		if err != nil {
			utils.HttpFailJson(r, utils.NewErrorCode(r.GetCtx(), 89999999, err.Error()))
			return
		}
		data := gconv.Map(param)
		delete(data, "idArr")
		if len(data) == 0 {
			utils.HttpFailJson(r, utils.NewErrorCode(r.GetCtx(), 89999999, ""))
			return
		}
		filter := map[string]interface{}{"id": param.IdArr}
		/**--------参数处理 结束--------**/

		/**--------权限验证 开始--------**/
		_, err = service.Action().CheckAuth(r.Context(), "authMenuUpdate")
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		/**--------权限验证 结束--------**/

		_, err = service.Menu().Update(r.Context(), data, filter)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		utils.HttpSuccessJson(r, map[string]interface{}{}, 0)
	}
}

// 删除
func (controllerThis *Menu) Delete(r *ghttp.Request) {
	sceneCode := utils.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	case "platformAdmin":
		/**--------参数处理 开始--------**/
		var param *apiAuth.MenuDeleteReq
		err := r.Parse(&param)
		if err != nil {
			utils.HttpFailJson(r, utils.NewErrorCode(r.GetCtx(), 89999999, err.Error()))
			return
		}
		filter := map[string]interface{}{"id": param.IdArr}
		/**--------参数处理 结束--------**/

		/**--------权限验证 开始--------**/
		_, err = service.Action().CheckAuth(r.Context(), "authMenuDelete")
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		/**--------权限验证 结束--------**/

		_, err = service.Menu().Delete(r.Context(), filter)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		utils.HttpSuccessJson(r, map[string]interface{}{}, 0)
	}
}

// 菜单树
func (cMenu *Menu) Tree(r *ghttp.Request) {
	/**--------参数处理 开始--------**/
	var param *apiAuth.MenuTreeReq
	err := r.Parse(&param)
	if err != nil {
		r.Response.Writeln(err.Error())
		return
	}
	filter := gconv.Map(param.Filter)
	if filter == nil {
		filter = map[string]interface{}{}
	}
	/**--------参数处理 结束--------**/

	sceneCode := utils.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	case "platformAdmin":
		/**--------权限验证 开始--------**/
		isAuth, _ := service.Action().CheckAuth(r.Context(), "authMenuLook")
		allowField := []string{"menuId", "menuName", "id"}
		if isAuth {
			allowField = daoAuth.Menu.ColumnArr()
			allowField = append(allowField, "id", "sceneName", "pMenuName")
			//allowField = gset.NewStrSetFrom(allowField).Diff(gset.NewStrSetFrom([]string{"password"})).Slice() //移除敏感字段
		}
		field := allowField
		if len(param.Field) > 0 {
			field = gset.NewStrSetFrom(param.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
			if len(field) == 0 {
				field = allowField
			}
		}

		filter["isStop"] = 0              //补充条件
		field = append(field, "menuTree") //补充字段（菜单树所需）
		/**--------权限验证 结束--------**/

		list, err := service.Menu().List(r.Context(), filter, field, [][2]string{}, 0, 0)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		tree, err := service.Menu().Tree(r.Context(), list, 0)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		utils.HttpSuccessJson(r, map[string]interface{}{"tree": tree}, 0)
	}
}
