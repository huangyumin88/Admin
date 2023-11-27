package logic

import (
	"api/internal/dao"
	daoUser "api/internal/dao/user"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/util/gconv"
)

type sUserWallets struct{}

func NewUserWallets() *sUserWallets {
	return &sUserWallets{}
}

func init() {
	service.RegisterUserWallets(NewUserWallets())
}

// 新增
func (logicThis *sUserWallets) Create(ctx context.Context, data map[string]interface{}) (id int64, err error) {
	daoThis := daoUser.Wallets
	id, err = dao.NewDaoHandler(ctx, &daoThis).Insert(data).GetModel().InsertAndGetId()
	return
}

// 修改
func (logicThis *sUserWallets) Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (row int64, err error) {
	daoThis := daoUser.Wallets
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
func (logicThis *sUserWallets) Delete(ctx context.Context, filter map[string]interface{}) (row int64, err error) {
	daoThis := daoUser.Wallets
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
