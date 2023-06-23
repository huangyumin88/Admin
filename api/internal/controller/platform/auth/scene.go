package controller

import (
	apiAuth "api/api/platform/auth"
	daoAuth "api/internal/dao/auth"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/util/gconv"
)

type Scene struct{}

func NewScene() *Scene {
	return &Scene{}
}

// 列表
func (controllerThis *Scene) List(ctx context.Context, req *apiAuth.SceneListReq) (res *apiAuth.SceneListRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.Map(req.Filter)
	if filter == nil {
		filter = map[string]interface{}{}
	}
	order := []string{req.Sort}
	page := req.Page
	limit := req.Limit

	columnsThis := daoAuth.Scene.Columns()
	allowField := daoAuth.Scene.ColumnArr()
	allowField = append(allowField, `id`, `name`)
	// allowField = gset.NewStrSetFrom(allowField).Diff(gset.NewStrSetFrom([]string{columnsThis.Password})).Slice() //移除敏感字段
	field := allowField
	if len(req.Field) > 0 {
		field = gset.NewStrSetFrom(req.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
		if len(field) == 0 {
			field = allowField
		}
	}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	isAuth, _ := service.Action().CheckAuth(ctx, `authSceneLook`)
	if !isAuth {
		field = []string{`id`, `name`, columnsThis.SceneId, columnsThis.SceneName}
	}
	/**--------权限验证 结束--------**/

	count, err := service.Scene().Count(ctx, filter)
	if err != nil {
		return
	}
	list, err := service.Scene().List(ctx, filter, field, order, page, limit)
	if err != nil {
		return
	}
	res = &apiAuth.SceneListRes{
		Count: count,
	}
	list.Structs(&res.List)
	return
}

// 详情
func (controllerThis *Scene) Info(ctx context.Context, req *apiAuth.SceneInfoReq) (res *apiAuth.SceneInfoRes, err error) {
	/**--------参数处理 开始--------**/
	allowField := daoAuth.Scene.ColumnArr()
	allowField = append(allowField, `id`, `name`)
	//allowField = gset.NewStrSetFrom(allowField).Diff(gset.NewStrSetFrom([]string{`password`})).Slice() //移除敏感字段
	field := allowField
	if len(req.Field) > 0 {
		field = gset.NewStrSetFrom(req.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
		if len(field) == 0 {
			field = allowField
		}
	}
	filter := map[string]interface{}{`id`: req.Id}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.Action().CheckAuth(ctx, `authSceneLook`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	info, err := service.Scene().Info(ctx, filter, field)
	if err != nil {
		return
	}
	res = &apiAuth.SceneInfoRes{}
	info.Struct(&res.Info)
	return
}

// 新增
func (controllerThis *Scene) Create(ctx context.Context, req *apiAuth.SceneCreateReq) (res *apiAuth.SceneCreateRes, err error) {
	/**--------参数处理 开始--------**/
	data := gconv.Map(req)
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.Action().CheckAuth(ctx, `authSceneCreate`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	id, err := service.Scene().Create(ctx, data)
	if err != nil {
		return
	}
	res = &apiAuth.SceneCreateRes{
		Id: id,
	}
	return
}

// 修改
func (controllerThis *Scene) Update(ctx context.Context, req *apiAuth.SceneUpdateReq) (res *apiAuth.SceneUpdateRes, err error) {
	/**--------参数处理 开始--------**/
	data := gconv.Map(req)
	delete(data, `idArr`)
	if len(data) == 0 {
		err = utils.NewErrorCode(ctx, 89999999, ``)
		return
	}
	filter := map[string]interface{}{`id`: req.IdArr}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.Action().CheckAuth(ctx, `authSceneUpdate`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	_, err = service.Scene().Update(ctx, filter, data)
	if err != nil {
		return
	}
	return
}

// 删除
func (controllerThis *Scene) Delete(ctx context.Context, req *apiAuth.SceneDeleteReq) (res *apiAuth.SceneDeleteRes, err error) {
	/**--------参数处理 开始--------**/
	filter := map[string]interface{}{`id`: req.IdArr}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.Action().CheckAuth(ctx, `authSceneDelete`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	_, err = service.Scene().Delete(ctx, filter)
	if err != nil {
		return
	}
	return
}