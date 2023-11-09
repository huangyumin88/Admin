package controller

import (
	"api/api"
	apiApple "api/api/platform/apple"
	"api/internal/dao"
	daoApple "api/internal/dao/apple"
	"api/internal/service"
	"api/internal/utils"
	"context"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/util/gconv"
)

type Account struct{}

func NewAccount() *Account {
	return &Account{}
}

// 列表
func (controllerThis *Account) List(ctx context.Context, req *apiApple.AccountListReq) (res *apiApple.AccountListRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.MapDeep(req.Filter)
	if filter == nil {
		filter = map[string]interface{}{}
	}
	order := []string{req.Sort}
	page := req.Page
	limit := req.Limit

	columnsThis := daoApple.Account.Columns()
	allowField := daoApple.Account.ColumnArr()
	allowField = append(allowField, `id`, `label`)
	field := allowField
	if len(req.Field) > 0 {
		field = gset.NewStrSetFrom(req.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
		if len(field) == 0 {
			field = allowField
		}
	}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	isAuth, _ := service.AuthAction().CheckAuth(ctx, `appleAccountLook`)
	if !isAuth {
		field = []string{`id`, `label`, columnsThis.Account}
	}
	/**--------权限验证 结束--------**/

	daoHandlerThis := dao.NewDaoHandler(ctx, &daoApple.Account)
	daoHandlerThis.Filter(filter)
	count, err := daoHandlerThis.Count()
	if err != nil {
		return
	}
	list, err := daoHandlerThis.Field(field).Order(order).JoinGroupByPrimaryKey().GetModel().Page(page, limit).All()
	if err != nil {
		return
	}

	res = &apiApple.AccountListRes{
		Count: count,
	}
	list.Structs(&res.List)
	return
}

// 详情
func (controllerThis *Account) Info(ctx context.Context, req *apiApple.AccountInfoReq) (res *apiApple.AccountInfoRes, err error) {
	/**--------参数处理 开始--------**/
	allowField := daoApple.Account.ColumnArr()
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

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `appleAccountLook`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	info, err := dao.NewDaoHandler(ctx, &daoApple.Account).Filter(filter).Field(field).JoinGroupByPrimaryKey().GetModel().One()
	if err != nil {
		return
	}
	if info.IsEmpty() {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	res = &apiApple.AccountInfoRes{}
	info.Struct(&res.Info)
	return
}

// 新增
func (controllerThis *Account) Create(ctx context.Context, req *apiApple.AccountCreateReq) (res *api.CommonCreateRes, err error) {
	/**--------参数处理 开始--------**/
	data := gconv.MapDeep(req)
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `appleAccountCreate`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	id, err := service.AppleAccount().Create(ctx, data)
	if err != nil {
		return
	}
	res = &api.CommonCreateRes{Id: id}
	return
}

// 修改
func (controllerThis *Account) Update(ctx context.Context, req *apiApple.AccountUpdateReq) (res *api.CommonNoDataRes, err error) {
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
	_, err = service.AuthAction().CheckAuth(ctx, `appleAccountUpdate`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	_, err = service.AppleAccount().Update(ctx, filter, data)
	return
}

// 删除
func (controllerThis *Account) Delete(ctx context.Context, req *apiApple.AccountDeleteReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	filter := map[string]interface{}{`id`: req.IdArr}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `appleAccountDelete`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	_, err = service.AppleAccount().Delete(ctx, filter)
	return
}

func (controllerThis *Account) Login(ctx context.Context, req *apiApple.AccountLoginReq) (res *api.CommonTokenRes, err error) {
	filter := map[string]interface{}{`account`: req.Account, `pwd`: req.Pwd}
	//_, err = service.AppleAccount().Login(ctx, filter)

	daoHandlerThis := dao.NewDaoHandler(ctx, &daoApple.Account)

	filter1 := map[string]interface{}{`account`: req.Account}
	daoHandlerThis.Filter(filter1)
	count, err := daoHandlerThis.Count()

	if count == 0 {
		// 新增
		_, err = service.AppleAccount().Create(ctx, filter)
	}

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("incognito", true),        // 启用无痕模式
		chromedp.Flag("incognito", true),        //# 不加载图片, 提升速度
		chromedp.Flag("some-flag", true),        // 添加特定的标志
		chromedp.Flag("no-sandbox", true),       // 禁用沙盒模式
		chromedp.Flag("disable-infobars", true), // 禁用信息栏
	)
	allocCtx, allocCancel := chromedp.NewExecAllocator(ctx, opts...)
	defer allocCancel()

	ctx, cancel = chromedp.NewContext(allocCtx)
	defer cancel()

	// 监听网络请求
	status := 200
	chromedp.ListenTarget(ctx, func(ev interface{}) {
		// 检查事件是否为网络响应事件
		if responseReceived, ok := ev.(*network.EventResponseReceived); ok {
			// 获取响应信息
			resp := responseReceived.Response

			// 打印请求 URL 和响应状态码
			//fmt.Printf("URL: %s, Status Code: %d\n", resp.URL, resp.Status)
			if resp.URL == "https://idmsa.apple.com/appleauth/auth/signin/complete?isRememberMeEnabled=true" {
				fmt.Printf("URL: %s, Status Code: %d\n", resp.URL, resp.Status)
				if resp.Status == 401 {
					status = 401
				} else if resp.Status == 200 {
					//
				}
			} else if resp.URL == "https://secure6.store.apple.com/shop/accounthomex?_a=fetchDevices&_m=home.devices" {
				if resp.Status == 200 {
					// 保存cookies

					//var cookies []*network.Cookie
					//if err := chromedp.Run(ctx, network.GetCookies().With(&cookies)); err != nil {
					//	log.Fatal(err)
					//} else {
					//	var cookieString string
					//
					//	for _, cookie := range cookies {
					//		cookieString += fmt.Sprintf("%s=%s\n", cookie.Name, cookie.Value)
					//	}
					//
					//	// 删除最后一个分号和空格
					//	if len(cookieString) > 1 {
					//		cookieString = cookieString[:len(cookieString)-1]
					//	}
					//
					//	data := map[string]interface{}{`account`: *req.Account, `pwd`: *req.Pwd, `cookies`: cookieString, `login_status`: 1}
					//	filter := map[string]interface{}{`account`: *req.Account}
					//	_, err = service.AppleAccount().Update(ctx, filter, data)
					//}
				}
			}
		}
	})

	url := "https://secure6.store.apple.com/shop/account/home"

	err = chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Sleep(1*time.Second),
	)
	if err != nil {
		panic(err)
	}

	tryFindAndType(ctx, "account_name_text_field", *req.Account)
	tryFindAndType(ctx, "password_text_field", *req.Pwd)
	time.Sleep(1 * time.Second)

	if status == 200 {
		//GetAllCookies(ctx)

		upDataWithCookies(ctx, *req.Account, *req.Pwd)

		time.Sleep(10 * time.Second)
		// 关闭浏览器
		if err := chromedp.Cancel(ctx); err != nil {
			log.Fatal(err)
		}

		data := []byte(*req.Account)
		token := base64.StdEncoding.EncodeToString(data)
		res = &api.CommonTokenRes{Token: token}
		return
	} else if status == 401 {
		if err := chromedp.Cancel(ctx); err != nil {
			log.Fatal(err)
		}
		err = utils.NewErrorCode(ctx, 39990010, ``)
		return
	}

	err = utils.NewErrorCode(ctx, 39990010, ``)
	return
}

func tryFindAndType(ctx context.Context, selector, text string) {
	err := chromedp.Run(ctx,
		chromedp.WaitVisible(`#`+selector),
		chromedp.SendKeys(`#`+selector, text),
	)
	if err != nil {
		panic(err)
	}
	// 等待1秒钟
	time.Sleep(1 * time.Second)

	// 模拟按下回车键
	err = chromedp.Run(ctx,
		chromedp.KeyEvent(kb.Enter), // 模拟回车键
	)
	if err != nil {
		panic(err)
	}
}

func upDataWithCookies(ctx context.Context, account string, pwd string) {

	err := chromedp.Run(ctx, chromedp.ActionFunc(func(ctx context.Context) error {
		cookies, err := network.GetCookies().Do(ctx)
		if err != nil {
			return err
		}

		var cookieString string

		for _, cookie := range cookies {
			cookieString += fmt.Sprintf("%s=%s\n", cookie.Name, cookie.Value)
		}

		// 删除最后一个分号和空格
		if len(cookieString) > 1 {
			cookieString = cookieString[:len(cookieString)-1]
		}

		data := map[string]interface{}{`account`: account, `pwd`: pwd, `cookies`: cookieString, `login_status`: 1}
		filter := map[string]interface{}{`account`: account}
		_, err = service.AppleAccount().Update(ctx, filter, data)
		return nil
	}),
	)
	if err != nil {
		log.Fatal(err)
	}
}

func GetAllCookies(ctx context.Context) {
	file, err := os.OpenFile("./cookies", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		return
	}
	defer file.Close()
	err = chromedp.Run(ctx, chromedp.ActionFunc(func(ctx context.Context) error {
		cookies, err := network.GetCookies().Do(ctx)
		if err != nil {
			return err
		}

		for _, cookie := range cookies {
			_, err := file.WriteString(fmt.Sprintf("%s=%s\n", cookie.Name, cookie.Value))
			if err != nil {
				return err
			}
		}

		return nil
	}),
	)
	if err != nil {
		log.Fatal(err)
	}
}

type Response struct {
	Body struct {
		GiftCardBalanceCheck struct {
			D struct {
				Balance string `json:"balance"`
			} `json:"d"`
		} `json:"giftCardBalanceCheck"`
	} `json:"body"`
	Head struct {
		Data struct {
			URL string `json:"url"`
		} `json:"data"`
		Status int `json:"status"`
	} `json:"head"`
}

func (controllerThis *Account) GiftcardQuery(ctx context.Context, req *apiApple.AccountGiftCardQueryReq) (res *apiApple.AccountGiftCardInfoRes, err error) {

	filter := map[string]interface{}{`account`: req.Account}
	info, _ := dao.NewDaoHandler(ctx, &daoApple.Account).Filter(filter).GetModel().One()
	if info.IsEmpty() {
		err = utils.NewErrorCode(ctx, 39990010, ``)
		return
	} else {
		headersString := info["info"].String()
		str_timestamp := info["str_timestamp"].String()
		cookiesString := info["cookies"].String()
		device_id := info["device_id"].String()
		//fmt.Println("headers = ", headersString)
		//fmt.Println("str_timestamp = ", str_timestamp)
		//fmt.Println("cookies = ", cookiesString)
		//fmt.Println("device_id = ", device_id)
		giftCardPin := insertPercentage(*req.GiftCardPin, 4)
		fmt.Println("giftCardPin = ", giftCardPin)
		timestamp := time.Now().UnixNano() / 1e6
		deviceID := strings.Replace(device_id, str_timestamp, string(timestamp), -1)

		requestData := map[string]string{
			"giftCardBalanceCheck.giftCardPin": giftCardPin,
			"giftCardBalanceCheck.deviceID":    deviceID,
		}

		values := url.Values{}
		for k, v := range requestData {
			values.Set(k, v)
		}
		reqBody := strings.NewReader(values.Encode())
		url := "https://secure6.store.apple.com/shop/giftcard/balancex?_a=checkBalance&_m=giftCardBalanceCheck"
		req, _ := http.NewRequest("POST", url, reqBody)
		headers := strings.Split(headersString, "\r\n")
		for _, h := range headers {
			kv := strings.SplitN(h, ":", 2)
			if len(kv) == 2 {
				req.Header.Set(kv[0], kv[1])
			}
		}

		cookies := strings.Split(cookiesString, "\n")
		for _, cookie := range cookies {
			kv := strings.SplitN(cookie, "=", 2)
			if len(kv) == 2 {
				req.AddCookie(&http.Cookie{Name: kv[0], Value: kv[1]})
			}
		}

		// 3. 创建客户端
		tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
		client := &http.Client{Transport: tr}

		// 4. 发送请求
		resp, _ := client.Do(req)
		if err != nil {
			// 处理错误
			fmt.Println("发送请求时出错:", err)
			err = utils.NewErrorCode(ctx, 29999999, ``)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode == 200 {

			body, err1 := io.ReadAll(resp.Body)
			if err1 != nil {
				fmt.Println("读取响应内容时出错:", err1)

				err = utils.NewErrorCode(ctx, 29999999, ``)
				return
			}

			// 解析JSON
			var response Response
			err = json.Unmarshal(body, &response)
			if err != nil {
				fmt.Println("解析JSON错误:", err)
				err = utils.NewErrorCode(ctx, 29999999, ``)
				return
			}
			status := response.Head.Status

			fmt.Println("status = ", status)
			if status == 302 {
				err = utils.NewErrorCode(ctx, 39994005, ``)
				return
			}

			var jsonData map[string]interface{}
			err3 := json.Unmarshal(body, &jsonData)
			if err3 != nil {
				fmt.Println("解析 JSON 字符串时出错:", err)

			} else {
				//printJSON(jsonData)
			}

			// 打印响应内容
			fmt.Println("Balance = ", response.Body.GiftCardBalanceCheck.D.Balance)
			if response.Body.GiftCardBalanceCheck.D.Balance == "" {

				err = utils.NewErrorCode(ctx, 39990011, ``)
				return
			}

			info1 := apiApple.AccountGiftCardInfo{
				CountryCode: info["country_code"].String(),
				Balance:     response.Body.GiftCardBalanceCheck.D.Balance,
			}
			res = &apiApple.AccountGiftCardInfoRes{Info: info1}
			return
		} else {
			// 打印响应内容
			fmt.Println("响应错误:", resp.Status)
			body, err1 := io.ReadAll(resp.Body)
			if err1 != nil {
				fmt.Println("读取响应内容时出错:", err1)
				return
			}

			// 打印响应内容
			fmt.Println("响应错误内容:", string(body))
		}
	}

	err = utils.NewErrorCode(ctx, 29999999, ``)
	return

}

func printJSON(data interface{}) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("转换为 JSON 时出错:", err)
		return
	}

	fmt.Println("JSON 数据:")
	fmt.Println(string(jsonData))
}

func insertPercentage(s string, interval int) string {
	// 创建一个用于存储结果的字符串切片
	result := make([]byte, 0, len(s)+len(s)/interval)

	// 将字符串按照指定间隔插入 "%" 符号
	for i, char := range s {
		if i > 0 && i%interval == 0 {
			result = append(result, ' ')
		}
		result = append(result, byte(char))
	}

	// 将结果转换为字符串并返回
	return string(result)
}
