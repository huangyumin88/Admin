package controller

import (
	"api/api"
	apiImage "api/api/platform/image"
	"api/internal/dao"
	daoImage "api/internal/dao/image"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/util/gconv"
)

type Image struct{}

func NewImage() *Image {
	return &Image{}
}

// 列表
func (controllerThis *Image) List(ctx context.Context, req *apiImage.ImageListReq) (res *apiImage.ImageListRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.MapDeep(req.Filter)
	if filter == nil {
		filter = map[string]interface{}{}
	}

	allowField := daoImage.Image.ColumnArr()
	allowField = append(allowField, `id`)
	field := allowField
	if len(req.Field) > 0 {
		field = gset.NewStrSetFrom(req.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
		if len(field) == 0 {
			field = allowField
		}
	}
	/**--------参数处理 结束--------**/

	daoHandlerThis := dao.NewDaoHandler(ctx, &daoImage.Image).Filter(filter)
	count, err := daoHandlerThis.Count()
	if err != nil {
		return
	}
	list, err := daoHandlerThis.Field(field).Order([]string{req.Sort}).JoinGroupByPrimaryKey().GetModel().Page(req.Page, req.Limit).All()
	if err != nil {
		return
	}

	res = &apiImage.ImageListRes{Count: count, List: []apiImage.ImageListItem{}}
	list.Structs(&res.List)
	return
}

// 详情
func (controllerThis *Image) Info(ctx context.Context, req *apiImage.ImageInfoReq) (res *apiImage.ImageInfoRes, err error) {
	/**--------参数处理 开始--------**/
	allowField := daoImage.Image.ColumnArr()
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

	info, err := dao.NewDaoHandler(ctx, &daoImage.Image).Filter(filter).Field(field).JoinGroupByPrimaryKey().GetModel().One()
	if err != nil {
		return
	}
	if info.IsEmpty() {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	res = &apiImage.ImageInfoRes{}
	info.Struct(&res.Info)
	return
}

// 新增
func (controllerThis *Image) Create(ctx context.Context, req *apiImage.ImageCreateReq) (res *api.CommonCreateRes, err error) {
	/**--------参数处理 开始--------**/
	data := gconv.MapDeep(req)
	/**--------参数处理 结束--------**/

	id, err := service.Image().Create(ctx, data)
	if err != nil {
		return
	}
	res = &api.CommonCreateRes{Id: id}
	return
}

// 修改
func (controllerThis *Image) Update(ctx context.Context, req *apiImage.ImageUpdateReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	data := gconv.MapDeep(req)
	delete(data, `idArr`)
	if len(data) == 0 {
		err = utils.NewErrorCode(ctx, 89999999, ``)
		return
	}
	filter := map[string]interface{}{`id`: req.IdArr}
	/**--------参数处理 结束--------**/

	_, err = service.Image().Update(ctx, filter, data)
	return
}

// 删除
func (controllerThis *Image) Delete(ctx context.Context, req *apiImage.ImageDeleteReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	filter := map[string]interface{}{`id`: req.IdArr}
	/**--------参数处理 结束--------**/

	_, err = service.Image().Delete(ctx, filter)
	return
}
