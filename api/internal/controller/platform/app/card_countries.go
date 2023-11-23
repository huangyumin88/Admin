package controller

import (
	"api/api"
	apiApp "api/api/platform/app"
	"api/internal/dao"
	daoApp "api/internal/dao/app"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/util/gconv"
)

type CardCountries struct{}

func NewCardCountries() *CardCountries {
	return &CardCountries{}
}

// 列表
func (controllerThis *CardCountries) List(ctx context.Context, req *apiApp.CardCountriesListReq) (res *apiApp.CardCountriesListRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.MapDeep(req.Filter)
	if filter == nil {
		filter = map[string]interface{}{}
	}
	order := []string{req.Sort}
	page := req.Page
	limit := req.Limit

	columnsThis := daoApp.CardCountries.Columns()
	allowField := daoApp.CardCountries.ColumnArr()
	allowField = append(allowField, `id`, `label`)
	field := allowField
	if len(req.Field) > 0 {
		field = gset.NewStrSetFrom(req.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
		if len(field) == 0 {
			field = allowField
		}
	}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	isAuth, _ := service.AuthAction().CheckAuth(ctx, `appCardCountriesLook`)
	if !isAuth {
		field = []string{`id`, `label`, columnsThis.Name}
	}
	/**--------权限验证 结束--------**/

	daoHandlerThis := dao.NewDaoHandler(ctx, &daoApp.CardCountries)
	daoHandlerThis.Filter(filter)
	count, err := daoHandlerThis.Count()
	if err != nil {
		return
	}
	list, err := daoHandlerThis.Field(field).Order(order).JoinGroupByPrimaryKey().GetModel().Page(page, limit).All()
	if err != nil {
		return
	}

	res = &apiApp.CardCountriesListRes{Count: count, List: []apiApp.CardCountriesListItem{}}
	list.Structs(&res.List)
	return
}

// 详情
func (controllerThis *CardCountries) Info(ctx context.Context, req *apiApp.CardCountriesInfoReq) (res *apiApp.CardCountriesInfoRes, err error) {
	/**--------参数处理 开始--------**/
	allowField := daoApp.CardCountries.ColumnArr()
	allowField = append(allowField, `id`, `label`)
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
	_, err = service.AuthAction().CheckAuth(ctx, `appCardCountriesLook`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	info, err := dao.NewDaoHandler(ctx, &daoApp.CardCountries).Filter(filter).Field(field).JoinGroupByPrimaryKey().GetModel().One()
	if err != nil {
		return
	}
	if info.IsEmpty() {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	res = &apiApp.CardCountriesInfoRes{}
	info.Struct(&res.Info)
	return
}

// 新增
func (controllerThis *CardCountries) Create(ctx context.Context, req *apiApp.CardCountriesCreateReq) (res *api.CommonCreateRes, err error) {
	/**--------参数处理 开始--------**/
	data := gconv.MapDeep(req)
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `appCardCountriesCreate`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	id, err := service.AppCardCountries().Create(ctx, data)
	if err != nil {
		return
	}
	res = &api.CommonCreateRes{Id: id}
	return
}

// 修改
func (controllerThis *CardCountries) Update(ctx context.Context, req *apiApp.CardCountriesUpdateReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	data := gconv.MapDeep(req)
	delete(data, `idArr`)
	if len(data) == 0 {
		err = utils.NewErrorCode(ctx, 89999999, ``)
		return
	}
	filter := map[string]interface{}{`id`: req.IdArr}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `appCardCountriesUpdate`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	_, err = service.AppCardCountries().Update(ctx, filter, data)
	return
}

// 删除
func (controllerThis *CardCountries) Delete(ctx context.Context, req *apiApp.CardCountriesDeleteReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	filter := map[string]interface{}{`id`: req.IdArr}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `appCardCountriesDelete`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	_, err = service.AppCardCountries().Delete(ctx, filter)
	return
}
