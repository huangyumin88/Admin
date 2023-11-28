package controller

import (
	"api/api"
	apiUser "api/api/platform/user"
	"api/internal/dao"
	daoUser "api/internal/dao/user"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/util/gconv"
)

type Wallets struct{}

func NewWallets() *Wallets {
	return &Wallets{}
}

// 列表
func (controllerThis *Wallets) List(ctx context.Context, req *apiUser.WalletsListReq) (res *apiUser.WalletsListRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.MapDeep(req.Filter)
	if filter == nil {
		filter = map[string]interface{}{}
	}
	order := []string{req.Sort}
	page := req.Page
	limit := req.Limit

	columnsThis := daoUser.Wallets.Columns()
	allowField := daoUser.Wallets.ColumnArr()
	//allowField = append(allowField, `id`, `user_name`, `info`)
	allowField = append(allowField, `id`, `user_name`)
	field := allowField
	if len(req.Field) > 0 {
		field = gset.NewStrSetFrom(req.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
		if len(field) == 0 {
			field = allowField
		}
	}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	isAuth, _ := service.AuthAction().CheckAuth(ctx, `userWalletsLook`)
	if !isAuth {
		field = []string{`id`, columnsThis.WalletId}
	}
	/**--------权限验证 结束--------**/

	daoHandlerThis := dao.NewDaoHandler(ctx, &daoUser.Wallets)
	daoHandlerThis.Filter(filter)
	count, err := daoHandlerThis.Count()
	if err != nil {
		return
	}
	list, err := daoHandlerThis.Field(field).Order(order).JoinGroupByPrimaryKey().GetModel().Page(page, limit).All()
	if err != nil {
		return
	}

	res = &apiUser.WalletsListRes{Count: count, List: []apiUser.WalletsListItem{}}
	list.Structs(&res.List)
	return
}

// 详情
func (controllerThis *Wallets) Info(ctx context.Context, req *apiUser.WalletsInfoReq) (res *apiUser.WalletsInfoRes, err error) {
	/**--------参数处理 开始--------**/
	allowField := daoUser.Wallets.ColumnArr()
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
	_, err = service.AuthAction().CheckAuth(ctx, `userWalletsLook`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	info, err := dao.NewDaoHandler(ctx, &daoUser.Wallets).Filter(filter).Field(field).JoinGroupByPrimaryKey().GetModel().One()
	if err != nil {
		return
	}
	if info.IsEmpty() {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	res = &apiUser.WalletsInfoRes{}
	info.Struct(&res.Info)
	return
}

// 修改
func (controllerThis *Wallets) Update(ctx context.Context, req *apiUser.WalletsUpdateReq) (res *api.CommonNoDataRes, err error) {
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
	_, err = service.AuthAction().CheckAuth(ctx, `userWalletsUpdate`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	_, err = service.UserWallets().Update(ctx, filter, data)
	return
}
