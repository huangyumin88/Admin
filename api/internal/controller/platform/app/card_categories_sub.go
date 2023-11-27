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

type CardCategoriesSub struct{}

func NewCardCategoriesSub() *CardCategoriesSub {
	return &CardCategoriesSub{}
}

// 列表
func (controllerThis *CardCategoriesSub) List(ctx context.Context, req *apiApp.CardCategoriesSubListReq) (res *apiApp.CardCategoriesSubListRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.MapDeep(req.Filter)
	if filter == nil {
		filter = map[string]interface{}{}
	}
	order := []string{req.Sort}
	page := req.Page
	limit := req.Limit

	columnsThis := daoApp.CardCategoriesSub.Columns()
	allowField := daoApp.CardCategoriesSub.ColumnArr()
	allowField = append(allowField, `id`)
	field := allowField
	if len(req.Field) > 0 {
		field = gset.NewStrSetFrom(req.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
		if len(field) == 0 {
			field = allowField
		}
	}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	isAuth, _ := service.AuthAction().CheckAuth(ctx, `appCardCategoriesSubLook`)
	if !isAuth {
		field = []string{`id`, columnsThis.Name}
	}
	/**--------权限验证 结束--------**/

	daoHandlerThis := dao.NewDaoHandler(ctx, &daoApp.CardCategoriesSub)
	daoHandlerThis.Filter(filter)
	count, err := daoHandlerThis.Count()
	if err != nil {
		return
	}
	list, err := daoHandlerThis.Field(field).Order(order).JoinGroupByPrimaryKey().GetModel().Page(page, limit).All()
	if err != nil {
		return
	}

	res = &apiApp.CardCategoriesSubListRes{Count: count, List: []apiApp.CardCategoriesSubListItem{}}
	list.Structs(&res.List)
	return
}

// 详情
func (controllerThis *CardCategoriesSub) Info(ctx context.Context, req *apiApp.CardCategoriesSubInfoReq) (res *apiApp.CardCategoriesSubInfoRes, err error) {
	/**--------参数处理 开始--------**/
	allowField := daoApp.CardCategoriesSub.ColumnArr()
	allowField = append(allowField, `id`)
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
	_, err = service.AuthAction().CheckAuth(ctx, `appCardCategoriesSubLook`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	info, err := dao.NewDaoHandler(ctx, &daoApp.CardCategoriesSub).Filter(filter).Field(field).JoinGroupByPrimaryKey().GetModel().One()
	if err != nil {
		return
	}
	if info.IsEmpty() {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	res = &apiApp.CardCategoriesSubInfoRes{}
	info.Struct(&res.Info)
	return
}

// 新增
func (controllerThis *CardCategoriesSub) Create(ctx context.Context, req *apiApp.CardCategoriesSubCreateReq) (res *api.CommonCreateRes, err error) {
	/**--------参数处理 开始--------**/
	data := gconv.MapDeep(req)
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `appCardCategoriesSubCreate`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	id, err := service.AppCardCategoriesSub().Create(ctx, data)
	if err != nil {
		return
	}
	res = &api.CommonCreateRes{Id: id}
	return
}

// 修改
func (controllerThis *CardCategoriesSub) Update(ctx context.Context, req *apiApp.CardCategoriesSubUpdateReq) (res *api.CommonNoDataRes, err error) {
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
	_, err = service.AuthAction().CheckAuth(ctx, `appCardCategoriesSubUpdate`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	_, err = service.AppCardCategoriesSub().Update(ctx, filter, data)
	return
}

// 删除
func (controllerThis *CardCategoriesSub) Delete(ctx context.Context, req *apiApp.CardCategoriesSubDeleteReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	filter := map[string]interface{}{`id`: req.IdArr}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `appCardCategoriesSubDelete`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	_, err = service.AppCardCategoriesSub().Delete(ctx, filter)
	return
}
