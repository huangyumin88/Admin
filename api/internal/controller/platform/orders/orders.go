package controller

import (
	"api/api"
	apiOrders "api/api/platform/orders"
	"api/internal/dao"
	daoOrders "api/internal/dao/orders"
	"api/internal/service"
	"api/internal/utils"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/google/uuid"
	"strings"
	"time"
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

	columnsThis := daoOrders.Orders.Columns()
	allowField := daoOrders.Orders.ColumnArr()
	allowField = append(allowField, `id`, `user_name`, `salesperson_name`)
	field := allowField
	if len(req.Field) > 0 {
		field = gset.NewStrSetFrom(req.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
		if len(field) == 0 {
			field = allowField
		}
	}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	isAuth, _ := service.AuthAction().CheckAuth(ctx, `ordersLook`)
	if !isAuth {
		field = []string{`id`, columnsThis.OrderId}
	}
	/**--------权限验证 结束--------**/

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

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `ordersLook`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

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

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `ordersCreate`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

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
	data["salesperson_id"] = loginInfo[`loginId`]

	id, err := service.Orders().Create(ctx, data)
	if err != nil {
		return
	}

	// 创建操作记录
	data1 := make(map[string]interface{})
	data1["order_id"] = id
	data1["actions_user_id"] = loginInfo[`loginId`]
	data1["backend_status"] = data["backend_status"]
	insertRecord(ctx, data1)

	res = &api.CommonCreateRes{Id: id}
	return
}

func insertRecord(ctx context.Context, data map[string]interface{}) (ids int64, err error) {

	ids, err = service.OrdersOrdersActions().Create(ctx, data)
	if err != nil {
		fmt.Println(err)
	}
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

	backend_status := data["backend_status"]
	if backend_status == "Failed" {
		data["client_status"] = "Failed"
	} else if backend_status == "Completed" {
		data["client_status"] = "Completed"
	} else {
		data["client_status"] = "Pending"
	}

	fmt.Println(backend_status)
	fmt.Println(data["client_status"])

	_, err = service.Orders().Update(ctx, filter, data)

	loginInfo := utils.GetCtxLoginInfo(ctx)

	// 创建操作记录
	data1 := make(map[string]interface{})
	data1["order_id"] = data["order_id"]
	data1["actions_user_id"] = loginInfo[`loginId`]
	data1["backend_status"] = data["backend_status"]
	insertRecord(ctx, data1)
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
	// 创建操作记录

	//if len(req.IdArr) >0 {
	//	loginInfo := utils.GetCtxLoginInfo(ctx)
	//
	//}
	loginInfo := utils.GetCtxLoginInfo(ctx)
	actions_user_id := loginInfo[`loginId`]

	for _, orderID := range req.IdArr {
		data := make(map[string]interface{})
		filter := map[string]interface{}{`id`: req.IdArr}
		data["actions_user_id"] = actions_user_id
		data["order_id"] = orderID
		data["backend_status"] = "Delete"
		info, err1 := dao.NewDaoHandler(ctx, &daoOrders.Orders).Filter(filter).JoinGroupByPrimaryKey().GetModel().One()
		if err1 != nil {
			break
		}
		var builder strings.Builder
		builder.WriteString("user_id：" + info["user_id"].String() + "\n")
		builder.WriteString("salesperson_id：" + info["salesperson_id"].String() + "\n")
		builder.WriteString("client_status：" + info["client_status"].String() + "\n")
		builder.WriteString("backend_status：" + info["backend_status"].String() + "\n")
		builder.WriteString("trade_files：" + info["trade_files"].String() + "\n")
		builder.WriteString("failed_reason：" + info["failed_reason"].String() + "\n")
		builder.WriteString("failed_files：" + info["failed_files"].String() + "\n")
		builder.WriteString("trade_amount：" + info["trade_amount"].String() + "\n")
		builder.WriteString("payable_amount：" + info["payable_amount"].String() + "\n")
		builder.WriteString("card_cate_id：" + info["card_cate_id"].String() + "\n")
		builder.WriteString("card_cate_sub_id：" + info["card_cate_sub_id"].String() + "\n")
		result := builder.String()
		data["remarks"] = result
		insertRecord(ctx, data)
	}

	_, err = service.Orders().Delete(ctx, filter)
	return
}

// 订单状态查询
func (controllerThis *Orders) CheckOrderStatus(ctx context.Context, req *apiOrders.OrdersQueryOrderStatusReq) (res *apiOrders.OrdersQueryOrderStatusRes, err error) {
	orderId := req.OrderId

	//fmt.Println("CheckOrderStatus")
	//_, err = service.AuthAction().CheckAuth(ctx, `ordersLook`)
	//if err != nil {
	//	return
	//}

	if orderId > 0 {
		filter := map[string]interface{}{`order_id`: orderId}
		info, err1 := dao.NewDaoHandler(ctx, &daoOrders.Orders).Filter(filter).JoinGroupByPrimaryKey().GetModel().One()
		if err1 != nil {
			return
		}

		if info.IsEmpty() {
			err = utils.NewErrorCode(ctx, 29999998, ``)
			return
		}

		if req.OrderType == 0 {
			// 后台展示的订单状态
			res = &apiOrders.OrdersQueryOrderStatusRes{List: []apiOrders.OrdersQueryOrderStatusItem{}}

			return
		} else {
			res = &apiOrders.OrdersQueryOrderStatusRes{List: []apiOrders.OrdersQueryOrderStatusItem{}}

			return
		}

	} else {
		if req.OrderType == 0 {
			// 后台展示的订单状态
			//后台订单状态：Pending - 等待审核; Loading - 加载中;  Failed - 加载失败; Pledging - 质押中; Completed - 交易完成;

			item1 := apiOrders.OrdersQueryOrderStatusItem{Status: "Pending", Label: "等待审核"}
			item2 := apiOrders.OrdersQueryOrderStatusItem{Status: "Loading", Label: "加载中"}
			item3 := apiOrders.OrdersQueryOrderStatusItem{Status: "Failed", Label: "加载失败"}
			item4 := apiOrders.OrdersQueryOrderStatusItem{Status: "Pledging", Label: "质押中"}
			item5 := apiOrders.OrdersQueryOrderStatusItem{Status: "Completed", Label: "交易完成"}

			status := req.Status
			if status == "" {
				res = &apiOrders.OrdersQueryOrderStatusRes{List: []apiOrders.OrdersQueryOrderStatusItem{item1, item2, item3, item4, item5}}
			} else if status == "Pending" {
				res = &apiOrders.OrdersQueryOrderStatusRes{List: []apiOrders.OrdersQueryOrderStatusItem{item2, item3, item4, item5}}
			} else if status == "Loading" {
				res = &apiOrders.OrdersQueryOrderStatusRes{List: []apiOrders.OrdersQueryOrderStatusItem{item3, item4, item5}}
			} else if status == "Failed" {
				res = &apiOrders.OrdersQueryOrderStatusRes{List: []apiOrders.OrdersQueryOrderStatusItem{item3}}
			} else if status == "Pledging" {
				res = &apiOrders.OrdersQueryOrderStatusRes{List: []apiOrders.OrdersQueryOrderStatusItem{item3, item5}}
			} else {
				res = &apiOrders.OrdersQueryOrderStatusRes{List: []apiOrders.OrdersQueryOrderStatusItem{item5}}
			}

			return
		} else {
			//用户订单状态：Pending - 交易中; Failed - 交易失败; Completed - 交易完成; Closed - 关闭;
			item1 := apiOrders.OrdersQueryOrderStatusItem{Status: "Pending", Label: "交易中"}
			item3 := apiOrders.OrdersQueryOrderStatusItem{Status: "Failed", Label: "加载失败"}
			item5 := apiOrders.OrdersQueryOrderStatusItem{Status: "Completed", Label: "交易完成"}
			item4 := apiOrders.OrdersQueryOrderStatusItem{Status: "Closed", Label: "关闭"}
			res = &apiOrders.OrdersQueryOrderStatusRes{List: []apiOrders.OrdersQueryOrderStatusItem{item1, item3, item5, item4}}
			return
		}
	}
	res = &apiOrders.OrdersQueryOrderStatusRes{List: []apiOrders.OrdersQueryOrderStatusItem{}}

	return
}
