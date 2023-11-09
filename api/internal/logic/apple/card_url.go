package logic

import (
	"api/internal/dao"
	daoApple "api/internal/dao/apple"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/util/gconv"
)

type sAppleCardUrl struct{}

func NewAppleCardUrl() *sAppleCardUrl {
	return &sAppleCardUrl{}
}

func init() {
	service.RegisterAppleCardUrl(NewAppleCardUrl())
}

// 新增
func (logicThis *sAppleCardUrl) Create(ctx context.Context, data map[string]interface{}) (id int64, err error) {
	daoThis := daoApple.CardUrl
	id, err = dao.NewDaoHandler(ctx, &daoThis).Insert(data).GetModel().InsertAndGetId()
	return
}

// 修改
func (logicThis *sAppleCardUrl) Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (row int64, err error) {
	daoThis := daoApple.CardUrl
	daoHandlerThis := dao.NewDaoHandler(ctx, &daoThis).Filter(filter)
	idArr, _ := daoHandlerThis.GetModel(true).Array(daoThis.PrimaryKey())
	if len(idArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}
	hookData := map[string]interface{}{}

	row, err = daoHandlerThis.Update(data).HookUpdate(hookData, gconv.SliceInt(idArr)...).GetModel().UpdateAndGetAffected()
	return
}

// 删除
func (logicThis *sAppleCardUrl) Delete(ctx context.Context, filter map[string]interface{}) (row int64, err error) {
	daoThis := daoApple.CardUrl
	daoHandlerThis := dao.NewDaoHandler(ctx, &daoThis).Filter(filter)
	idArr, _ := daoHandlerThis.GetModel(true).Array(daoThis.PrimaryKey())
	if len(idArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	result, err := daoHandlerThis.HookDelete(gconv.SliceInt(idArr)...).GetModel().Delete()
	row, _ = result.RowsAffected()
	return
}
