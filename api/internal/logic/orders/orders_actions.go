package logic

import (
	"api/internal/dao"
	daoOrders "api/internal/dao/orders"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/util/gconv"
)

type sOrdersOrdersActions struct{}

func NewOrdersOrdersActions() *sOrdersOrdersActions {
	return &sOrdersOrdersActions{}
}

func init() {
	service.RegisterOrdersOrdersActions(NewOrdersOrdersActions())
}

// 新增
func (logicThis *sOrdersOrdersActions) Create(ctx context.Context, data map[string]interface{}) (id int64, err error) {
	daoThis := daoOrders.OrdersActions
	id, err = dao.NewDaoHandler(ctx, &daoThis).Insert(data).GetModel().InsertAndGetId()
	return
}

// 修改
func (logicThis *sOrdersOrdersActions) Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (row int64, err error) {
	daoThis := daoOrders.OrdersActions
	daoHandlerThis := dao.NewDaoHandler(ctx, &daoThis).Filter(filter)
	idArr, _ := daoHandlerThis.GetModel(true).Array(daoThis.PrimaryKey())
	if len(idArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}
	hookData := map[string]interface{}{}

	row, err = daoHandlerThis.Update(data).HookUpdate(hookData, gconv.SliceUint(idArr)...).GetModel().UpdateAndGetAffected()
	return
}

// 删除
func (logicThis *sOrdersOrdersActions) Delete(ctx context.Context, filter map[string]interface{}) (row int64, err error) {
	daoThis := daoOrders.OrdersActions
	daoHandlerThis := dao.NewDaoHandler(ctx, &daoThis).Filter(filter)
	idArr, _ := daoHandlerThis.GetModel(true).Array(daoThis.PrimaryKey())
	if len(idArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	result, err := daoHandlerThis.HookDelete(gconv.SliceUint(idArr)...).GetModel().Delete()
	row, _ = result.RowsAffected()
	return
}
