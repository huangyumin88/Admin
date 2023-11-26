package controller

import (
	"api/api"
	apiBanks "api/api/platform/banks"
	"api/internal/dao"
	daoBanks "api/internal/dao/banks"
	daoUser "api/internal/dao/user"
	"api/internal/service"
	"context"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/util/gconv"
)

type BanksUser struct{}

func NewBanksUser() *BanksUser {
	return &BanksUser{}
}

// 列表
func (controllerThis *BanksUser) List(ctx context.Context, req *apiBanks.BanksUserListReq) (res *apiBanks.BanksUserListRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.MapDeep(req.Filter)
	if filter == nil {
		filter = map[string]interface{}{}
	}
	order := []string{req.Sort}
	page := req.Page
	limit := req.Limit

	//columnsThis := daoBanks.BanksUser.Columns()
	allowField := daoBanks.BanksUser.ColumnArr()
	allowField = append(allowField, `id`, daoUser.User.Columns().Account)
	field := allowField
	if len(req.Field) > 0 {
		field = gset.NewStrSetFrom(req.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
		if len(field) == 0 {
			field = allowField
		}
	}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	isAuth, _ := service.AuthAction().CheckAuth(ctx, `banksBanksUserLook`)
	if !isAuth {
		field = []string{`id`}
	}
	/**--------权限验证 结束--------**/

	daoHandlerThis := dao.NewDaoHandler(ctx, &daoBanks.BanksUser)
	daoHandlerThis.Filter(filter)
	count, err := daoHandlerThis.Count()
	if err != nil {
		return
	}
	list, err := daoHandlerThis.Field(field).Order(order).JoinGroupByPrimaryKey().GetModel().Page(page, limit).All()
	if err != nil {
		return
	}

	res = &apiBanks.BanksUserListRes{Count: count, List: []apiBanks.BanksUserListItem{}}
	list.Structs(&res.List)
	return
}

// 删除
func (controllerThis *BanksUser) Delete(ctx context.Context, req *apiBanks.BanksUserDeleteReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	filter := map[string]interface{}{`id`: req.IdArr}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `banksBanksUserDelete`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	_, err = service.BanksBanksUser().Delete(ctx, filter)
	return
}
