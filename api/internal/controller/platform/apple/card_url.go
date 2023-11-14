package controller

import (
	"api/api"
	apiApple "api/api/platform/apple"
	"sync"

	//"api/api/platform"
	apiLogin "api/api/platform"
	controllerLogin "api/internal/controller/platform"
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
	allowField = append(allowField, `id`, `label`)
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
	for i := range res.List {
		// 直接通过索引修改 list 中的元素
		res.List[i].Label = res.List[i].CountryCode

		filter1 := map[string]interface{}{}
		filter1["id"] = res.List[i].AccountId

		info, _ := dao.NewDaoHandler(ctx, &daoApple.Account).Filter(filter1).JoinGroupByPrimaryKey().GetModel().One()
		if info != nil {
			info.Struct(&res.List[i].AccountInfo)
		}
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

func (controllerThis *CardUrl) AutomaticLogin(ctx context.Context, req *apiApple.CardUrlAutomaticLoginReq) (res *api.CommonNoDataRes, err error) {
	var IsStop uint = 0
	filter := apiApple.CardUrlListFilter{IsStop: &IsStop}

	var reqs = &apiApple.CardUrlListReq{Limit: 0, Page: 1, Filter: filter}

	reps, _ := controllerThis.List(ctx, reqs)
	println(len(reps.List))
	go processRequests(ctx, reps.List)

	return
}

func processRequests(ctx context.Context, List []apiApple.CardUrlListItem) {
	for _, item := range List {
		var reqs = &apiLogin.LoginAppleReq{Account: item.AccountInfo.Account, Pwd: item.AccountInfo.Pwd, Code: *item.Id}
		loginApi := controllerLogin.NewLogin()
		loginApi.AppleLogin(ctx, reqs)
	}
}

//func (controllerThis *CardUrl) GiftcardAutoQuery(ctx context.Context, req *apiApple.CardUrlQueryReq) (res *apiApple.AccountGiftCardInfoRes, err error) {
//	var IsStop uint = 0
//	filter := apiApple.CardUrlListFilter{IsStop: &IsStop}
//
//	var reqs = &apiApple.CardUrlListReq{Limit: 0, Page: 1, Filter: filter}
//
//	reps, _ := controllerThis.List(ctx, reqs)
//
//	var wg sync.WaitGroup
//
//	results := make(chan apiApple.AccountGiftCardInfoRes)
//	done := make(chan struct{})
//	maxConcurrency := 3
//
//	// 控制并发数量
//	concurrencyControl := make(chan struct{}, maxConcurrency)
//
//	// 启动协程处理请求
//	for _, item := range reps.List {
//		wg.Add(1)
//		concurrencyControl <- struct{}{}
//		go func(data apiApple.CardUrlListItem) {
//			defer func() { <-concurrencyControl }()
//			processCheckRequests(ctx, item, req.GiftCardPin, results, done, &wg)
//		}(item)
//	}
//
//	// 等待结果并检查是否有code == 200
//	go func() {
//		for res := range results {
//			if len(res.Info.Balance) > 0 {
//				close(done)
//				fmt.Println("找到code == 200的响应:", res, res.Info.Balance)
//
//				return
//			}
//		}
//	}()
//
//	// 等待所有请求完成
//	wg.Wait()
//	close(results)
//
//	info := apiApple.AccountGiftCardInfo{
//		CountryCode: "US",
//		Balance:     "100.0",
//	}
//
//	println(len(reps.List))
//	res = &apiApple.AccountGiftCardInfoRes{Info: info}
//	return
//}

//func processCheckRequests(ctx context.Context, item apiApple.CardUrlListItem, giftCardPin *string, results chan<- apiApple.AccountGiftCardInfoRes, done chan struct{}, wg *sync.WaitGroup) {
//	defer wg.Done()
//
//	var reqs = &apiApple.AccountGiftCardQueryReq{Account: item.AccountInfo.Account, Pwd: item.AccountInfo.Pwd, Code: *item.Id, GiftCardPin: giftCardPin}
//
//	account := Account{}
//	response, _ := account.GiftcardQuery(ctx, reqs)
//
//	select {
//	case results <- *response:
//	case <-done:
//		return
//	}
//}

func (controllerThis *CardUrl) GiftcardAutoQuery(ctx context.Context, req *apiApple.CardUrlQueryReq) (res *apiApple.AccountGiftCardInfoRes, err error) {
	// 创建一个可取消的上下文
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var IsStop uint = 0
	filter := apiApple.CardUrlListFilter{IsStop: &IsStop}
	reqs := &apiApple.CardUrlListReq{Limit: 0, Page: 1, Filter: filter}

	reps, err := controllerThis.List(ctx, reqs)
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	results := make(chan apiApple.AccountGiftCardInfoRes)
	maxConcurrency := 3

	// 控制并发数量
	concurrencyControl := make(chan struct{}, maxConcurrency)

	for _, item := range reps.List {
		wg.Add(1)
		concurrencyControl <- struct{}{}
		go func(data apiApple.CardUrlListItem) {
			defer func() { <-concurrencyControl }()
			processCheckRequests(ctx, data, req.GiftCardPin, results, &wg)
		}(item)
	}

	go func() {
		for res := range results {

			if res.Info.Status == 1 {
				cancel() // 满足条件时取消上下文
				break
			}
		}
	}()

	wg.Wait()

	for r := range results {
		println("Status = ", r.Info.Status)
		println("Balance = ", r.Info.Balance)

		if r.Info.Status == 1 {
			close(results)
			info := apiApple.AccountGiftCardInfo{}

			info.Balance = r.Info.Balance
			info.CountryCode = r.Info.CountryCode
			info.Status = 1
			res = &apiApple.AccountGiftCardInfoRes{Info: info}
			return
		}
	}
	close(results)
	err = utils.NewErrorCode(ctx, 39990011, ``)
	return
}

func processCheckRequests(ctx context.Context, item apiApple.CardUrlListItem, giftCardPin *string, results chan<- apiApple.AccountGiftCardInfoRes, wg *sync.WaitGroup) {
	defer wg.Done()

	// 检查上下文是否已被取消
	if ctx.Err() != nil {
		return
	}

	// 执行请求
	reqs := &apiApple.AccountGiftCardQueryReq{Account: item.AccountInfo.Account, Pwd: item.AccountInfo.Pwd, Code: *item.Id, GiftCardPin: giftCardPin}
	account := Account{}
	response, err := account.GiftcardQuery(ctx, reqs)
	if err != nil {
		// 处理错误
		return
	}

	select {
	case results <- *response:
	case <-ctx.Done(): // 检查是否接收到取消信号
		return
	}
}
