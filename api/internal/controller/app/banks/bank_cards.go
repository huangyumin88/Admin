package controller

import (
	"api/api"
	apiBanks "api/api/app/banks"
	"api/internal/dao"
	daoBanks "api/internal/dao/banks"
	daoUser "api/internal/dao/user"
	"api/internal/service"
	"api/internal/utils"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/util/gconv"
)

type BankCards struct{}

func NewBankCards() *BankCards {
	return &BankCards{}
}

// 列表
func (controllerThis *BankCards) List(ctx context.Context, req *apiBanks.BankCardsListReq) (res *apiBanks.BankCardsListRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.MapDeep(req.Filter)
	if filter == nil {
		filter = map[string]interface{}{}
	}
	order := []string{req.Sort}
	page := req.Page
	limit := req.Limit

	//columnsThis := daoBanks.BankCards.Columns()
	allowField := daoBanks.BankCards.ColumnArr()
	allowField = append(allowField, `id`, daoUser.User.Columns().UserId)
	field := allowField
	if len(req.Field) > 0 {
		field = gset.NewStrSetFrom(req.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
		if len(field) == 0 {
			field = allowField
		}
	}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	isAuth, _ := service.AuthAction().CheckAuth(ctx, `banksBankCardsLook`)
	if !isAuth {
		field = []string{`id`}
	}
	/**--------权限验证 结束--------**/

	daoHandlerThis := dao.NewDaoHandler(ctx, &daoBanks.BankCards)
	daoHandlerThis.Filter(filter)
	count, err := daoHandlerThis.Count()
	if err != nil {
		return
	}
	list, err := daoHandlerThis.Field(field).Order(order).JoinGroupByPrimaryKey().GetModel().Page(page, limit).All()
	if err != nil {
		return
	}

	res = &apiBanks.BankCardsListRes{Count: count, List: []apiBanks.BankCardsListItem{}}
	list.Structs(&res.List)
	return
}

// 详情
func (controllerThis *BankCards) Info(ctx context.Context, req *apiBanks.BankCardsInfoReq) (res *apiBanks.BankCardsInfoRes, err error) {
	/**--------参数处理 开始--------**/
	allowField := daoBanks.BankCards.ColumnArr()
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
	//_, err = service.AuthAction().CheckAuth(ctx, `banksBankCardsLook`)
	//if err != nil {
	//	return
	//}
	/**--------权限验证 结束--------**/

	info, err := dao.NewDaoHandler(ctx, &daoBanks.BankCards).Filter(filter).Field(field).JoinGroupByPrimaryKey().GetModel().One()
	if err != nil {
		return
	}
	if info.IsEmpty() {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	res = &apiBanks.BankCardsInfoRes{}
	info.Struct(&res.Info)
	return
}

// 新增
func (controllerThis *BankCards) Create(ctx context.Context, req *apiBanks.BankCardsCreateReq) (res *api.CommonCreateRes, err error) {
	/**--------参数处理 开始--------**/
	data := gconv.MapDeep(req)
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	//_, err = service.AuthAction().CheckAuth(ctx, `banksBankCardsCreate`)
	//if err != nil {
	//	return
	//}
	/**--------权限验证 结束--------**/

	// 先判断是否存在

	data1 := make(map[string]interface{})
	data1["user_id"] = *req.UserId
	data1["bank_id"] = *req.BankId

	info, err := dao.NewDaoHandler(ctx, &daoBanks.BankCards).Filter(data1).JoinGroupByPrimaryKey().GetModel().One()
	if err != nil {
		return
	}
	if info.IsEmpty() == false {
		err = utils.NewErrorCode(ctx, 29991061, ``)
		return
	}

	id, err := service.BanksBankCards().Create(ctx, data)
	if err != nil {
		return
	}

	ids, err := service.BanksBanksUser().Create(ctx, data1)
	if err != nil {
		data2 := make(map[string]interface{})
		data2["id"] = id
		service.BanksBankCards().Delete(ctx, data2)
		return
	}

	fmt.Println("ids", ids)

	res = &api.CommonCreateRes{Id: id}
	return
}

// 修改
func (controllerThis *BankCards) Update(ctx context.Context, req *apiBanks.BankCardsUpdateReq) (res *api.CommonNoDataRes, err error) {
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
	//_, err = service.AuthAction().CheckAuth(ctx, `banksBankCardsUpdate`)
	//if err != nil {
	//	return
	//}
	/**--------权限验证 结束--------**/

	_, err = service.BanksBankCards().Update(ctx, filter, data)
	return
}

// 删除
func (controllerThis *BankCards) Delete(ctx context.Context, req *apiBanks.BankCardsDeleteReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	filter := map[string]interface{}{`id`: req.IdArr}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	//_, err = service.AuthAction().CheckAuth(ctx, `banksBankCardsDelete`)
	//if err != nil {
	//	return
	//}
	/**--------权限验证 结束--------**/

	_, err = service.BanksBankCards().Delete(ctx, filter)

	if err == nil {
		_, err = service.BanksBanksUser().Delete(ctx, filter)
		return
	}

	return
}
