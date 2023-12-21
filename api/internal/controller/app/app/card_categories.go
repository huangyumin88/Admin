package controller

import (
	apiApp "api/api/app/app"
	"api/internal/dao"
	daoApp "api/internal/dao/app"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/container/gset"

	"github.com/gogf/gf/v2/util/gconv"
)

type CardCategories struct{}

func NewCardCategories() *CardCategories {
	return &CardCategories{}
}

// 列表
func (controllerThis *CardCategories) List(ctx context.Context, req *apiApp.CardCategoriesListReq) (res *apiApp.CardCategoriesListRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.MapDeep(req.Filter)
	if filter == nil {
		filter = map[string]interface{}{}
	}
	order := []string{req.Sort}
	page := req.Page
	limit := req.Limit
	fmt.Println("limit", limit)
	fmt.Println("page", page)

	//columnsThis := daoApp.CardCategories.Columns()
	allowField := daoApp.CardCategories.ColumnArr()
	allowField = append(allowField, `id`, `imUserArr`)
	field := allowField
	if len(req.Field) > 0 {
		field = gset.NewStrSetFrom(req.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
		if len(field) == 0 {
			field = allowField
		}
	}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	//isAuth, _ := service.AuthAction().CheckAuth(ctx, `appCardCategoriesLook`)
	//if !isAuth {
	//	field = []string{`id`, columnsThis.Name}
	//}
	/**--------权限验证 结束--------**/

	daoHandlerThis := dao.NewDaoHandler(ctx, &daoApp.CardCategories)
	daoHandlerThis.Filter(filter)
	count, err := daoHandlerThis.Count()
	if err != nil {
		return
	}
	list, err := daoHandlerThis.Field(field).Order(order).JoinGroupByPrimaryKey().GetModel().Page(page, limit).All()
	if err != nil {
		return
	}

	res = &apiApp.CardCategoriesListRes{Count: count, List: []apiApp.CardCategoriesListItem{}}
	list.Structs(&res.List)
	return
}
