package controller

import (
	"api/api"
	apiOrders "api/api/app/orders"
	"api/internal/dao"
	daoOrders "api/internal/dao/orders"
	"api/internal/service"
	"api/internal/utils"
	"context"
	"fmt"
	"github.com/google/uuid"
	"time"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/util/gconv"
)

type Orders struct{}

func NewOrders() *Orders {
	return &Orders{}
}

// 列表
func (controllerThis *Orders) List(ctx context.Context, req *apiOrders.OrdersListReq) (res *apiOrders.OrdersListRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.MapDeep(req.Filter)
	if filter == nil {
		filter = map[string]interface{}{}
	}
	order := []string{req.Sort}
	page := req.Page
	limit := req.Limit

	//columnsThis := daoOrders.Orders.Columns()
	allowField := daoOrders.Orders.ColumnArr()
	allowField = append(allowField, `id`)
	field := allowField
	if len(req.Field) > 0 {
		field = gset.NewStrSetFrom(req.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
		if len(field) == 0 {
			field = allowField
		}
	}
	/**--------参数处理 结束--------**/

	///**--------权限验证 开始--------**/
	//isAuth, _ := service.AuthAction().CheckAuth(ctx, `ordersLook`)
	//if !isAuth {
	//	field = []string{`id`, columnsThis.OrderId}
	//}
	///**--------权限验证 结束--------**/

	daoHandlerThis := dao.NewDaoHandler(ctx, &daoOrders.Orders)
	daoHandlerThis.Filter(filter)
	count, err := daoHandlerThis.Count()
	if err != nil {
		return
	}
	list, err := daoHandlerThis.Field(field).Order(order).JoinGroupByPrimaryKey().GetModel().Page(page, limit).All()
	if err != nil {
		return
	}

	res = &apiOrders.OrdersListRes{Count: count, List: []apiOrders.OrdersListItem{}}
	list.Structs(&res.List)
	return
}

// 详情
func (controllerThis *Orders) Info(ctx context.Context, req *apiOrders.OrdersInfoReq) (res *apiOrders.OrdersInfoRes, err error) {
	/**--------参数处理 开始--------**/
	allowField := daoOrders.Orders.ColumnArr()
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

	///**--------权限验证 开始--------**/
	//_, err = service.AuthAction().CheckAuth(ctx, `ordersLook`)
	//if err != nil {
	//	return
	//}
	///**--------权限验证 结束--------**/

	info, err := dao.NewDaoHandler(ctx, &daoOrders.Orders).Filter(filter).Field(field).JoinGroupByPrimaryKey().GetModel().One()
	if err != nil {
		return
	}
	if info.IsEmpty() {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	res = &apiOrders.OrdersInfoRes{}
	info.Struct(&res.Info)
	return
}

// 新增
func (controllerThis *Orders) Create(ctx context.Context, req *apiOrders.OrdersCreateReq) (res *api.CommonCreateRes, err error) {
	/**--------参数处理 开始--------**/
	data := gconv.MapDeep(req)
	/**--------参数处理 结束--------**/

	///**--------权限验证 开始--------**/
	//_, err = service.AuthAction().CheckAuth(ctx, `ordersCreate`)
	//if err != nil {
	//	return
	//}
	///**--------权限验证 结束--------**/

	// 获取当前时间并格式化为年月日时分秒毫秒
	currentTime := time.Now().Format("20060102150405")

	// 生成UUID并获取其后6位
	uuidWithHyphen := uuid.New()
	uuidString := uuidWithHyphen.String()
	lastSixUUID := uuidString[len(uuidString)-6:]

	// 拼接得到订单编号
	orderNumber := fmt.Sprintf("%s%s", currentTime, lastSixUUID)

	data["order_no"] = orderNumber

	loginInfo := utils.GetCtxLoginInfo(ctx)
	data["user_id"] = loginInfo[`loginId`]

	id, err := service.Orders().Create(ctx, data)
	if err != nil {
		return
	}
	res = &api.CommonCreateRes{Id: id}
	return
}

// 修改
func (controllerThis *Orders) Update(ctx context.Context, req *apiOrders.OrdersUpdateReq) (res *api.CommonNoDataRes, err error) {
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
	_, err = service.AuthAction().CheckAuth(ctx, `ordersUpdate`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	_, err = service.Orders().Update(ctx, filter, data)
	return
}

// 删除
func (controllerThis *Orders) Delete(ctx context.Context, req *apiOrders.OrdersDeleteReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	filter := map[string]interface{}{`id`: req.IdArr}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `ordersDelete`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	_, err = service.Orders().Delete(ctx, filter)
	return
}
