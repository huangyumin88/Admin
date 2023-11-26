package controller

import (
	apiBanks "api/api/app/banks"
	"api/internal/dao"
	daoBanks "api/internal/dao/banks"
	"api/internal/service"
	"context"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/util/gconv"
)

type Banks struct{}

func NewBanks() *Banks {
	return &Banks{}
}

// 列表
func (controllerThis *Banks) List(ctx context.Context, req *apiBanks.BanksListReq) (res *apiBanks.BanksListRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.MapDeep(req.Filter)
	if filter == nil {
		filter = map[string]interface{}{}
	}
	order := []string{req.Sort}
	page := req.Page
	limit := req.Limit

	columnsThis := daoBanks.Banks.Columns()
	allowField := daoBanks.Banks.ColumnArr()
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
	isAuth, _ := service.AuthAction().CheckAuth(ctx, `banksLook`)
	if !isAuth {
		field = []string{`id`, `label`, columnsThis.Name}
	}
	/**--------权限验证 结束--------**/

	daoHandlerThis := dao.NewDaoHandler(ctx, &daoBanks.Banks)
	daoHandlerThis.Filter(filter)
	count, err := daoHandlerThis.Count()
	if err != nil {
		return
	}
	list, err := daoHandlerThis.Order(order).JoinGroupByPrimaryKey().GetModel().Page(page, limit).All()
	if err != nil {
		return
	}

	res = &apiBanks.BanksListRes{Count: count, List: []apiBanks.BanksListItem{}}
	list.Structs(&res.List)
	return
}
