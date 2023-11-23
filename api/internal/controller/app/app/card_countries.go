package controller

import (
	apiApp "api/api/app/app"
	"api/internal/dao"
	daoApp "api/internal/dao/app"
	"api/internal/service"
	"context"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/util/gconv"
)

type CardCountries struct{}

func NewCardCountries() *CardCountries {
	return &CardCountries{}
}

// 列表
func (controllerThis *CardCountries) List(ctx context.Context, req *apiApp.CardCountriesListReq) (res *apiApp.CardCountriesListRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.MapDeep(req.Filter)
	if filter == nil {
		filter = map[string]interface{}{}
	}
	order := []string{req.Sort}
	page := req.Page
	limit := req.Limit

	columnsThis := daoApp.CardCountries.Columns()
	allowField := daoApp.CardCountries.ColumnArr()
	allowField = append(allowField, `id`, `currencyCode`, `name`, `flagAvatarID`)
	field := allowField
	if len(req.Field) > 0 {
		field = gset.NewStrSetFrom(req.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
		if len(field) == 0 {
			field = allowField
		}
	}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	isAuth, _ := service.AuthAction().CheckAuth(ctx, `appCardCountriesLook`)
	if !isAuth {
		field = []string{`id`, `label`, columnsThis.Name}
	}
	/**--------权限验证 结束--------**/

	daoHandlerThis := dao.NewDaoHandler(ctx, &daoApp.CardCountries)
	daoHandlerThis.Filter(filter)
	count, err := daoHandlerThis.Count()
	if err != nil {
		return
	}
	list, err := daoHandlerThis.Order(order).JoinGroupByPrimaryKey().GetModel().Page(page, limit).All()
	if err != nil {
		return
	}

	res = &apiApp.CardCountriesListRes{Count: count, List: []apiApp.CardCountriesListItem{}}
	list.Structs(&res.List)
	return
}
