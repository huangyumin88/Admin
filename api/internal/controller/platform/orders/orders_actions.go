package controller

import (
	"api/api"
	apiOrders "api/api/platform/orders"
	"api/internal/dao"
	daoOrders "api/internal/dao/orders"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/util/gconv"
)

type OrdersActions struct{}

func NewOrdersActions() *OrdersActions {
	return &OrdersActions{}
}

// 列表
func (controllerThis *OrdersActions) List(ctx context.Context, req *apiOrders.OrdersActionsListReq) (res *apiOrders.OrdersActionsListRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.MapDeep(req.Filter)
	if filter == nil {
		filter = map[string]interface{}{}
	}
	order := []string{req.Sort}
	page := req.Page
	limit := req.Limit

	//columnsThis := daoOrders.OrdersActions.Columns()
	allowField := daoOrders.OrdersActions.ColumnArr()
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
	isAuth, _ := service.AuthAction().CheckAuth(ctx, `ordersOrdersActionsLook`)
	if !isAuth {
		field = []string{`id`}
	}
	/**--------权限验证 结束--------**/

	daoHandlerThis := dao.NewDaoHandler(ctx, &daoOrders.OrdersActions)
	daoHandlerThis.Filter(filter)
	count, err := daoHandlerThis.Count()
	if err != nil {
		return
	}
	list, err := daoHandlerThis.Field(field).Order(order).JoinGroupByPrimaryKey().GetModel().Page(page, limit).All()
	if err != nil {
		return
	}

	res = &apiOrders.OrdersActionsListRes{Count: count, List: []apiOrders.OrdersActionsListItem{}}
	list.Structs(&res.List)
	return
}

// 详情
func (controllerThis *OrdersActions) Info(ctx context.Context, req *apiOrders.OrdersActionsInfoReq) (res *apiOrders.OrdersActionsInfoRes, err error) {
	/**--------参数处理 开始--------**/
	allowField := daoOrders.OrdersActions.ColumnArr()
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
	_, err = service.AuthAction().CheckAuth(ctx, `ordersOrdersActionsLook`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	info, err := dao.NewDaoHandler(ctx, &daoOrders.OrdersActions).Filter(filter).Field(field).JoinGroupByPrimaryKey().GetModel().One()
	if err != nil {
		return
	}
	if info.IsEmpty() {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	res = &apiOrders.OrdersActionsInfoRes{}
	info.Struct(&res.Info)
	return
}

// 新增
func (controllerThis *OrdersActions) Create(ctx context.Context, req *apiOrders.OrdersActionsCreateReq) (res *api.CommonCreateRes, err error) {
	/**--------参数处理 开始--------**/
	data := gconv.MapDeep(req)
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `ordersOrdersActionsCreate`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	id, err := service.OrdersOrdersActions().Create(ctx, data)
	if err != nil {
		return
	}
	res = &api.CommonCreateRes{Id: id}
	return
}

// 修改
func (controllerThis *OrdersActions) Update(ctx context.Context, req *apiOrders.OrdersActionsUpdateReq) (res *api.CommonNoDataRes, err error) {
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
	_, err = service.AuthAction().CheckAuth(ctx, `ordersOrdersActionsUpdate`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	_, err = service.OrdersOrdersActions().Update(ctx, filter, data)
	return
}

// 删除
func (controllerThis *OrdersActions) Delete(ctx context.Context, req *apiOrders.OrdersActionsDeleteReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	filter := map[string]interface{}{`id`: req.IdArr}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `ordersOrdersActionsDelete`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	_, err = service.OrdersOrdersActions().Delete(ctx, filter)
	return
}
