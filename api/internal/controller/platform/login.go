package controller

import (
	"api/api"
	apiCurrent "api/api/platform"
	"api/internal/cache"
	"api/internal/dao"
	daoApple "api/internal/dao/apple"
	daoAuth "api/internal/dao/auth"
	daoPlatform "api/internal/dao/platform"
	"api/internal/service"
	"api/internal/utils"
	"context"
	"fmt"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
)

type Login struct{}

func NewLogin() *Login {
	return &Login{}
}

// 获取加密盐
func (controllerThis *Login) Salt(ctx context.Context, req *apiCurrent.LoginSaltReq) (res *api.CommonSaltRes, err error) {
	adminColumns := daoPlatform.Admin.Columns()
	info, _ := dao.NewDaoHandler(ctx, &daoPlatform.Admin).Filter(g.Map{`loginName`: req.LoginName}).GetModel().One()
	if info.IsEmpty() {
		err = utils.NewErrorCode(ctx, 39990000, ``)
		return
	}
	if info[adminColumns.IsStop].Int() == 1 {
		err = utils.NewErrorCode(ctx, 39990002, ``)
		return
	}

	saltDynamic := grand.S(8)
	err = cache.NewSalt(ctx, req.LoginName).Set(saltDynamic, 5)
	if err != nil {
		return
	}
	res = &api.CommonSaltRes{SaltStatic: info[adminColumns.Salt].String(), SaltDynamic: saltDynamic}
	return
}

// 登录
func (controllerThis *Login) Login(ctx context.Context, req *apiCurrent.LoginLoginReq) (res *api.CommonTokenRes, err error) {
	adminColumns := daoPlatform.Admin.Columns()
	info, _ := dao.NewDaoHandler(ctx, &daoPlatform.Admin).Filter(g.Map{`loginName`: req.LoginName}).GetModel().One()
	if info.IsEmpty() {
		err = utils.NewErrorCode(ctx, 39990000, ``)
		return
	}
	if info[adminColumns.IsStop].Int() == 1 {
		err = utils.NewErrorCode(ctx, 39990002, ``)
		return
	}

	salt, _ := cache.NewSalt(ctx, req.LoginName).Get()
	if salt == `` || gmd5.MustEncrypt(info[adminColumns.Password].String()+salt) != req.Password {
		err = utils.NewErrorCode(ctx, 39990001, ``)
		return
	}

	claims := utils.CustomClaims{LoginId: info[daoPlatform.Admin.PrimaryKey()].Uint()}
	jwt := utils.NewJWT(ctx, utils.GetCtxSceneInfo(ctx)[daoAuth.Scene.Columns().SceneConfig].Map())
	token, err := jwt.CreateToken(claims)
	if err != nil {
		return
	}
	// cache.NewToken(ctx, claims.LoginId).Set(token, int64(jwt.ExpireTime)) //缓存token（限制多地登录，多设备登录等情况下用）

	res = &api.CommonTokenRes{Token: token}
	return
}

func (controllerThis *Login) AppleLogin(ctx context.Context, req *apiCurrent.LoginAppleReq) (res *api.CommonTokenRes, err error) {

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
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
	status := 999
	var deviceID string
	var str_timestamp string
	var headersString string
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
					status = 200
				}
			} else if strings.Contains(resp.URL, "https://secure6.store.apple.com/shop/signIn/idms/authx") {

			} else if resp.URL == "https://secure6.store.apple.com/shop/accounthomex?_a=fetchDevices&_m=home.devices" {

			} else if resp.URL == "https://secure6.store.apple.com/shop/account/home" {

			}
		}

	})

	chromedp.ListenTarget(ctx, func(ev interface{}) {
		if req, ok := ev.(*network.EventRequestWillBeSent); ok {
			request := req.Request
			if request.URL == "https://secure6.store.apple.com/shop/giftcard/balancex?_a=checkBalance&_m=giftCardBalanceCheck" {
				if request.PostData != "" {
					// 处理 post data
					fmt.Println("request.PostData ", request.PostData)

					// 查找字符串 "giftCardBalanceCheck.deviceID="
					startIndex := strings.Index(request.PostData, "giftCardBalanceCheck.deviceID=")

					if startIndex != -1 {
						// 如果未找到指定的字符串，执行相应的错误处理
						deviceID = request.PostData[startIndex+len("giftCardBalanceCheck.deviceID="):]

						//fmt.Println("Extracted data:", deviceID)

						re := regexp.MustCompile(`\d{13}`)

						// 在输入字符串中查找匹配的时间戳
						str_timestamp = re.FindString(deviceID)

						if str_timestamp != "" {
							fmt.Println("提取到的时间戳:", str_timestamp)
						} else {
							fmt.Println("未找到时间戳")
						}
					} else {
						fmt.Println("String not found.")
					}

					//fmt.Println("request.Headers ", request.Headers)

					for name, value := range request.Headers {
						headersString += name + ": " + fmt.Sprint(value) + "\r\n"
					}
					if len(headersString) > 2 {
						headersString = headersString[:len(headersString)-2]
					}
					fmt.Println("request.Headers \n", headersString)
				} else {
					// 没有 post data
					fmt.Println("没有请求数据 ")
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
	time.Sleep(3 * time.Second)
	if status == 200 {
		//GetAllCookies(ctx)
		var cookieString string
		var stk string
		var countryCode string

		upDataWithCookies(ctx, &cookieString, &stk, &countryCode)
		// 继续处理余额的页面
		time.Sleep(3 * time.Second)
		url = "https://secure.store.apple.com/shop/giftcard/balance"

		err = chromedp.Run(ctx,
			chromedp.Navigate(url),
			chromedp.Sleep(1*time.Second),
		)
		if err != nil {
			panic(err)
		}

		tryFindAndBalanceType(ctx)

		saveDataWith(ctx, *req.Account, *req.Pwd, cookieString, stk, countryCode, deviceID, str_timestamp, headersString)

		time.Sleep(1 * time.Second)

		info1, _ := dao.NewDaoHandler(ctx, &daoApple.Account).Filter(g.Map{`account`: *req.Account}).GetModel().One()

		claims := utils.CustomClaims{LoginId: info1[daoApple.Account.PrimaryKey()].Uint()}
		jwt := utils.NewJWT(ctx, utils.GetCtxSceneInfo(ctx)[daoAuth.Scene.Columns().SceneConfig].Map())
		token, err2 := jwt.CreateToken(claims)
		if err2 != nil {
			return nil, err2
		}

		res = &api.CommonTokenRes{Token: token}

		// 关闭浏览器
		if err := chromedp.Cancel(ctx); err != nil {
			log.Fatal(err)
		}

		return
	} else if status == 401 {
		if err := chromedp.Cancel(ctx); err != nil {
			log.Fatal(err)
		}
		err = utils.NewErrorCode(ctx, 39990010, ``)
		return
	}

	chromedp.Cancel(ctx)
	err = utils.NewErrorCode(ctx, 39990010, ``)
	return
}

func saveDataWith(ctx context.Context, account string, pwd string, cookieString string, stk string, countryCode string, deviceID string, str_timestamp string, headersString string) {
	data2 := map[string]interface{}{`account`: account, `pwd`: pwd, `cookies`: cookieString, `login_status`: 1, `stk`: stk, `country_code`: countryCode, `device_id`: deviceID, `str_timestamp`: str_timestamp, `info`: headersString}
	filter2 := map[string]interface{}{`account`: account}

	info, _ := dao.NewDaoHandler(ctx, &daoApple.Account).Filter(g.Map{`account`: account}).GetModel().One()

	if info.IsEmpty() {
		_, err := service.AppleAccount().Create(ctx, data2)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		_, err := service.AppleAccount().Update(ctx, filter2, data2)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func tryFindAndBalanceType(ctx context.Context) {
	err := chromedp.Run(ctx,
		chromedp.SendKeys(`input.form-textbox-input`, "XJ36Q684TG4VGX7Q"),
		chromedp.Sleep(1*time.Second),
		chromedp.Click(`#balanceCheck-balance`),
	)
	time.Sleep(3 * time.Second)
	if err != nil {
		panic(err)
	}
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

func upDataWithCookies(ctx context.Context, cookieString *string, stk *string, countryCode *string) {

	//var aosStk AosStk
	var html string
	var currentURL string

	err := chromedp.Run(ctx, chromedp.ActionFunc(func(ctx context.Context) error {
		cookies, err := network.GetCookies().Do(ctx)
		if err != nil {
			return err
		}

		for i, cookie := range cookies {
			*cookieString += fmt.Sprintf("%s=%s\n", cookie.Name, cookie.Value)

			if i < len(cookies)-1 {
				*cookieString += "\n"
			}
		}

		return nil
	}),
		//chromedp.Evaluate(`document.documentElement.outerHTML`, &html),
		chromedp.Tasks{
			chromedp.OuterHTML("html", &html),
			chromedp.Location(&currentURL),
		},
	)
	time.Sleep(2 * time.Second)
	if err != nil {
		log.Fatal(err)
	}
	// 查找x-aos-stk字段
	re := regexp.MustCompile(`"x-aos-stk":"(.*?)"`)
	match := re.FindStringSubmatch(html)
	if len(match) > 1 {
		stk1 := match[1]
		ix := strings.Index(stk1, `"`)
		if ix > -1 {
			stk1 = stk1[:ix]
		}
		fmt.Printf("子字符串之后的 42 个字符: %s\n", stk1)
		*stk = stk1
	} else {
		fmt.Println("未查找到")
	}

	re1 := regexp.MustCompile(`"countryCode":"(.*?)"`)
	match1 := re1.FindStringSubmatch(html)
	if len(match1) > 1 {
		countryCode1 := match1[1]
		ix := strings.Index(countryCode1, `"`)
		if ix > -1 {
			countryCode1 = countryCode1[:ix]
		}
		fmt.Printf("子字符串之后的 countryCode 个字符: %s\n", countryCode1)
		*countryCode = countryCode1
	} else {
		fmt.Println("未查找到")
		*countryCode = "United States"
	}
}
