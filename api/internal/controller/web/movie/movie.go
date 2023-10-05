package controller

import (
	apiMovie "api/api/web/movie"
	"api/internal/dao"
	daoMovie "api/internal/dao/movie"
	"api/internal/utils"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/util/gconv"
	"reflect"
	"strings"
)

type Movie struct{}

func NewMovie() *Movie {
	return &Movie{}
}

// 列表
func (controllerThis *Movie) List(ctx context.Context, req *apiMovie.MovieListReq) (res *apiMovie.MovieListRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.MapDeep(req.Filter)
	if filter == nil {
		filter = map[string]interface{}{}
	}
	filter["isStop"] = 0
	filter["isUpdate"] = 1
	order := []string{req.Sort}
	order = []string{}
	//order = append(order, "id ASC")
	order = append(order, "sort_time desc")
	//order = append(order, "score desc")
	page := req.Page
	limit := req.Limit

	allowField := daoMovie.Movie.ColumnArr()
	allowField = append(allowField, `id`, `label`)
	field := allowField
	if len(req.Field) > 0 {
		field = gset.NewStrSetFrom(req.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
		if len(field) == 0 {
			field = allowField
		}
	}
	/**--------参数处理 结束--------**/

	daoHandlerThis := dao.NewDaoHandler(ctx, &daoMovie.Movie)
	daoHandlerThis.Filter(filter)
	count, err := daoHandlerThis.Count()
	if err != nil {
		return
	}
	list, err := daoHandlerThis.Field(field).Order(order).JoinGroupByPrimaryKey().GetModel().Page(page, limit).All()
	if err != nil {
		return
	}

	Movie := daoMovie.Movie
	MovieCol := Movie.Columns()
	for _, item := range list {
		stringsx := item[MovieCol.Actors]
		interfaceValue := stringsx.Interface()
		str, ok := interfaceValue.(string)
		if ok {
			// 变量 str 成功转换为字符串类型
			// 现在可以在 str 中使用它
			fmt.Println("Converted String:", str)
			parts := strings.Split(str, ",")
			partsValue := reflect.ValueOf(parts)
			item["ActorsArray"] = partsValue
		} else {
			// 转换失败，stringsx 不是一个字符串类型
			fmt.Println("Conversion to string failed")
		}

		//parts := strings.Split(stringsx, ",")
		//item.ActorsArray = parts
		//fmt.Println(item.ActorsArray)
		//print(stringsx)
		break
	}

	res = &apiMovie.MovieListRes{
		Count: count,
	}
	list.Structs(&res.List)

	return
}

// 详情
func (controllerThis *Movie) Info(ctx context.Context, req *apiMovie.MovieInfoReq) (res *apiMovie.MovieInfoRes, err error) {
	/**--------参数处理 开始--------**/
	allowField := daoMovie.Movie.ColumnArr()
	allowField = append(allowField, `id`, `label`)
	field := allowField
	if len(req.Field) > 0 {
		field = gset.NewStrSetFrom(req.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
		if len(field) == 0 {
			field = allowField
		}
	}
	filter := map[string]interface{}{`id`: req.Id}
	/**--------参数处理 结束--------**/

	info, err := dao.NewDaoHandler(ctx, &daoMovie.Movie).Filter(filter).Field(field).JoinGroupByPrimaryKey().GetModel().One()
	if err != nil {
		return
	}
	if len(info) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	res = &apiMovie.MovieInfoRes{}
	info.Struct(&res.Info)
	return
}
