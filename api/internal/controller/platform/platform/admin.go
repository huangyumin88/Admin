package controller

import (
	"api/api"
	apiPlatform "api/api/platform/platform"
	"api/internal/dao"
	daoPlatform "api/internal/dao/platform"

	"api/internal/service"
	"api/internal/utils"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/tencentyun/tls-sig-api-v2-golang/tencentyun"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

type Admin struct{}

func NewAdmin() *Admin {
	return &Admin{}
}

const (
	sdkappid = 1600014653
	key      = "cb21e11dbc877362238d4e27bffeee7c3a02da4c8d86b4f0f1e2591fd5c583b2"
)

type Response struct {
	ActionStatus string `json:"ActionStatus"`
	ErrorInfo    string `json:"ErrorInfo"`
	ErrorCode    int    `json:"ErrorCode"`
}

// 列表
func (controllerThis *Admin) List(ctx context.Context, req *apiPlatform.AdminListReq) (res *apiPlatform.AdminListRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.MapDeep(req.Filter)
	if filter == nil {
		filter = map[string]interface{}{}
	}
	order := []string{req.Sort}
	page := req.Page
	limit := req.Limit

	columnsThis := daoPlatform.Admin.Columns()
	allowField := daoPlatform.Admin.ColumnArr()
	allowField = append(allowField, `id`, `label`)
	allowField = gset.NewStrSetFrom(allowField).Diff(gset.NewStrSetFrom([]string{columnsThis.Password, columnsThis.Salt})).Slice() //移除敏感字段
	field := allowField
	if len(req.Field) > 0 {
		field = gset.NewStrSetFrom(req.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
		if len(field) == 0 {
			field = allowField
		}
	}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	isAuth, _ := service.AuthAction().CheckAuth(ctx, `platformAdminLook`)
	if !isAuth {
		field = []string{`id`, `label`, columnsThis.Phone, columnsThis.Account, columnsThis.AdminId}
	}
	/**--------权限验证 结束--------**/

	daoHandlerThis := dao.NewDaoHandler(ctx, &daoPlatform.Admin)
	daoHandlerThis.Filter(filter)
	count, err := daoHandlerThis.Count()
	if err != nil {
		return
	}
	list, err := daoHandlerThis.Field(field).Order(order).JoinGroupByPrimaryKey().GetModel().Page(page, limit).All()
	if err != nil {
		return
	}

	res = &apiPlatform.AdminListRes{Count: count, List: []apiPlatform.AdminListItem{}}
	list.Structs(&res.List)
	return
}

// 详情
func (controllerThis *Admin) Info(ctx context.Context, req *apiPlatform.AdminInfoReq) (res *apiPlatform.AdminInfoRes, err error) {
	/**--------参数处理 开始--------**/
	allowField := daoPlatform.Admin.ColumnArr()
	allowField = append(allowField, `id`, `label`, `roleIdArr`)
	columnsThis := daoPlatform.Admin.Columns()
	allowField = gset.NewStrSetFrom(allowField).Diff(gset.NewStrSetFrom([]string{columnsThis.Password, columnsThis.Salt})).Slice() //移除敏感字段
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
	_, err = service.AuthAction().CheckAuth(ctx, `platformAdminLook`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	info, err := dao.NewDaoHandler(ctx, &daoPlatform.Admin).Filter(filter).Field(field).JoinGroupByPrimaryKey().GetModel().One()
	if err != nil {
		return
	}
	if info.IsEmpty() {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	res = &apiPlatform.AdminInfoRes{}
	info.Struct(&res.Info)
	return
}

// 新增
func (controllerThis *Admin) Create(ctx context.Context, req *apiPlatform.AdminCreateReq) (res *api.CommonCreateRes, err error) {
	/**--------参数处理 开始--------**/
	data := gconv.MapDeep(req)
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `platformAdminCreate`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/
	id, err := service.PlatformAdmin().Create(ctx, data)
	if err != nil {
		return
	}

	// 注册IM

	// 创建IMUSerID
	// https://cloud.tencent.com/document/product/269/1608

	// 初始化随机数生成器
	rand.Seed(time.Now().UnixNano())

	// 生成随机的 32 位无符号整数
	randomNum := rand.Uint32()

	sig, err := tencentyun.GenUserSig(sdkappid, key, "admin124", 86400*180)
	if err != nil {
		return
	}
	url := fmt.Sprintf("https://console.tim.qq.com/v4/im_open_login_svc/account_import?sdkappid=%d&identifier=admin124&usersig=%s&random=%d&contenttype=json", sdkappid, sig, randomNum)
	jsonStr := ""
	contentType := "application/json" // 内容类型
	imUserId := fmt.Sprintf(`platform_%d`, 1000+id)
	if len(*req.Avatar) > 0 {
		jsonStr = fmt.Sprintf(`{"UserID":"%s", "Nick":"%s", "FaceUrl":"%s"}`, imUserId, *req.Nickname, *req.Avatar)
	} else {
		jsonStr = fmt.Sprintf(`{"UserID":"%s", "Nick":"%s"}`, imUserId, *req.Nickname)
	}

	data3 := []byte(jsonStr)
	// 要发送的数据

	// 创建一个 HTTP POST 请求
	resp, err := http.Post(url, contentType, bytes.NewBuffer(data3))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// 解析 JSON 数据
	var responseObj Response
	if err1 := json.Unmarshal(body, &responseObj); err1 != nil {
		fmt.Println("Error parsing JSON:", err1)
		return
	}

	filter := map[string]interface{}{`adminId`: id}
	//daoThis := daoUser.User

	// 输出 ActionStatus 字段
	fmt.Println("ActionStatus:", responseObj.ActionStatus)
	if responseObj.ActionStatus == "OK" {
		data5 := make(map[string]interface{})
		data5["imUserId"] = imUserId
		_, err = service.PlatformAdmin().Update(ctx, filter, data5)
		if err != nil {
			service.PlatformAdmin().Delete(ctx, g.Map{daoPlatform.Admin.PrimaryKey(): id})
			return
		}
	} else {
		service.PlatformAdmin().Delete(ctx, g.Map{daoPlatform.Admin.PrimaryKey(): id})
		err = utils.NewErrorCode(ctx, 29999999, responseObj.ErrorInfo)
		return
	}

	res = &api.CommonCreateRes{Id: id}
	return
}

// 修改
func (controllerThis *Admin) Update(ctx context.Context, req *apiPlatform.AdminUpdateReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	data := gconv.MapDeep(req)
	delete(data, `idArr`)
	if len(data) == 0 {
		err = utils.NewErrorCode(ctx, 89999999, ``)
		return
	}

	if garray.NewFrom(gconv.SliceAny(req.IdArr)).Contains(g.Cfg().MustGet(ctx, `superPlatformAdminId`).Uint()) { //不能修改平台超级管理员
		err = utils.NewErrorCode(ctx, 30000000, ``)
		return
	}

	filter := map[string]interface{}{`id`: req.IdArr}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `platformAdminUpdate`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	_, err = service.PlatformAdmin().Update(ctx, filter, data)
	return
}

// 删除
func (controllerThis *Admin) Delete(ctx context.Context, req *apiPlatform.AdminDeleteReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	if garray.NewFrom(gconv.SliceAny(req.IdArr)).Contains(g.Cfg().MustGet(ctx, `superPlatformAdminId`).Uint()) { //不能删除平台超级管理员
		err = utils.NewErrorCode(ctx, 30000001, ``)
		return
	}

	filter := map[string]interface{}{`id`: req.IdArr}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `platformAdminDelete`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	_, err = service.PlatformAdmin().Delete(ctx, filter)
	return
}
