package logic

import (
	"api/internal/dao"
	daoApp "api/internal/dao/app"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/util/gconv"
)

type sAppCardCountries struct{}

func NewAppCardCountries() *sAppCardCountries {
	return &sAppCardCountries{}
}

func init() {
	service.RegisterAppCardCountries(NewAppCardCountries())
}

// 新增
func (logicThis *sAppCardCountries) Create(ctx context.Context, data map[string]interface{}) (id int64, err error) {
	daoThis := daoApp.CardCountries
	id, err = dao.NewDaoHandler(ctx, &daoThis).Insert(data).GetModel().InsertAndGetId()
	return
}

// 修改
func (logicThis *sAppCardCountries) Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (row int64, err error) {
	daoThis := daoApp.CardCountries
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
func (logicThis *sAppCardCountries) Delete(ctx context.Context, filter map[string]interface{}) (row int64, err error) {
	daoThis := daoApp.CardCountries
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
