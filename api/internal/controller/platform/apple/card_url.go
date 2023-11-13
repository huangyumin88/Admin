package controller

import (
	"api/api"
	apiApple "api/api/platform/apple"
	"api/internal/dao"
	daoApple "api/internal/dao/apple"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/util/gconv"
)

type CardUrl struct{}

func NewCardUrl() *CardUrl {
	return &CardUrl{}
}

// 列表
func (controllerThis *CardUrl) List(ctx context.Context, req *apiApple.CardUrlListReq) (res *apiApple.CardUrlListRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.MapDeep(req.Filter)
	if filter == nil {
		filter = map[string]interface{}{}
	}
	order := []string{req.Sort}
	page := req.Page
	limit := req.Limit

	//columnsThis := daoApple.CardUrl.Columns()
	allowField := daoApple.CardUrl.ColumnArr()
	allowField = append(allowField, `id`)
	field := allowField
	if len(req.Field) > 0 {
		target := "label"
		replacement := "country_code"

		// 遍历并替换
		for i, v := range req.Field {
			if v == target {
				req.Field[i] = replacement
				order[0] = "id asc"
				break // 如果只替换第一个匹配项，找到后即可停止循环
			}
		}
		field = gset.NewStrSetFrom(req.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
		if len(field) == 0 {
			field = allowField
		}
	}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	isAuth, _ := service.AuthAction().CheckAuth(ctx, `appleCardUrlLook`)
	if !isAuth {
		field = []string{`id`}
	}
	/**--------权限验证 结束--------**/

	daoHandlerThis := dao.NewDaoHandler(ctx, &daoApple.CardUrl)
	daoHandlerThis.Filter(filter)
	count, err := daoHandlerThis.Count()
	if err != nil {
		return
	}
	list, err := daoHandlerThis.Field(field).Order(order).JoinGroupByPrimaryKey().GetModel().Page(page, limit).All()

	if err != nil {
		return
	}

	res = &apiApple.CardUrlListRes{
		Count: count,
	}

	list.Structs(&res.List)

	//for _, item := range res.List {
	//	//str := "Hello, World!"
	//	//item.CountryCode = &str
	//	println(*item.CountryCode)
	//}
	for i := range res.List {
		// 直接通过索引修改 list 中的元素
		res.List[i].Label = res.List[i].CountryCode
	}
	return
}

// 详情
func (controllerThis *CardUrl) Info(ctx context.Context, req *apiApple.CardUrlInfoReq) (res *apiApple.CardUrlInfoRes, err error) {
	/**--------参数处理 开始--------**/
	allowField := daoApple.CardUrl.ColumnArr()
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
	_, err = service.AuthAction().CheckAuth(ctx, `appleCardUrlLook`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	info, err := dao.NewDaoHandler(ctx, &daoApple.CardUrl).Filter(filter).Field(field).JoinGroupByPrimaryKey().GetModel().One()
	if err != nil {
		return
	}
	if info.IsEmpty() {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	res = &apiApple.CardUrlInfoRes{}
	info.Struct(&res.Info)
	return
}

// 新增
func (controllerThis *CardUrl) Create(ctx context.Context, req *apiApple.CardUrlCreateReq) (res *api.CommonCreateRes, err error) {
	/**--------参数处理 开始--------**/
	data := gconv.MapDeep(req)
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `appleCardUrlCreate`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	id, err := service.AppleCardUrl().Create(ctx, data)
	if err != nil {
		return
	}
	res = &api.CommonCreateRes{Id: id}
	return
}

// 修改
func (controllerThis *CardUrl) Update(ctx context.Context, req *apiApple.CardUrlUpdateReq) (res *api.CommonNoDataRes, err error) {
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
	_, err = service.AuthAction().CheckAuth(ctx, `appleCardUrlUpdate`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	_, err = service.AppleCardUrl().Update(ctx, filter, data)
	return
}

// 删除
func (controllerThis *CardUrl) Delete(ctx context.Context, req *apiApple.CardUrlDeleteReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	filter := map[string]interface{}{`id`: req.IdArr}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `appleCardUrlDelete`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	_, err = service.AppleCardUrl().Delete(ctx, filter)
	return
}
