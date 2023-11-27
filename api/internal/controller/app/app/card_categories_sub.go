package controller

import (
	apiApp "api/api/app/app"
	"api/internal/dao"
	daoApp "api/internal/dao/app"
	"context"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/util/gconv"
)

type CardCategoriesSub struct{}

func NewCardCategoriesSub() *CardCategoriesSub {
	return &CardCategoriesSub{}
}

// 列表
func (controllerThis *CardCategoriesSub) List(ctx context.Context, req *apiApp.CardCategoriesSubListReq) (res *apiApp.CardCategoriesSubListRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.MapDeep(req.Filter)
	if filter == nil {
		filter = map[string]interface{}{}
	}
	order := []string{req.Sort}
	page := req.Page
	limit := req.Limit

	allowField := daoApp.CardCategoriesSub.ColumnArr()
	allowField = append(allowField, `id`, `label`)
	field := allowField
	if len(req.Field) > 0 {
		field = gset.NewStrSetFrom(req.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
		if len(field) == 0 {
			field = allowField
		}
	}
	/**--------参数处理 结束--------**/

	daoHandlerThis := dao.NewDaoHandler(ctx, &daoApp.CardCategoriesSub)
	daoHandlerThis.Filter(filter)
	count, err := daoHandlerThis.Count()
	if err != nil {
		return
	}
	list, err := daoHandlerThis.Field(field).Order(order).JoinGroupByPrimaryKey().GetModel().Page(page, limit).All()
	if err != nil {
		return
	}

	res = &apiApp.CardCategoriesSubListRes{Count: count, List: []apiApp.CardCategoriesSubListItem{}}
	list.Structs(&res.List)
	return
}
