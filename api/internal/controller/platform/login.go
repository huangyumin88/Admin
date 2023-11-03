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

		time.Sleep(1 * time.Second)
		// 关闭浏览器
		if err := chromedp.Cancel(ctx); err != nil {
			log.Fatal(err)
		}

		info, _ := dao.NewDaoHandler(ctx, &daoApple.Account).Filter(g.Map{`account`: *req.Account}).GetModel().One()

		claims := utils.CustomClaims{LoginId: info[daoApple.Account.PrimaryKey()].Uint()}
		jwt := utils.NewJWT(ctx, utils.GetCtxSceneInfo(ctx)[daoAuth.Scene.Columns().SceneConfig].Map())
		token, err2 := jwt.CreateToken(claims)
		if err2 != nil {
			return nil, err2
		}
		// cache.NewToken(ctx, claims.LoginId).Set(token, int64(jwt.ExpireTime)) //缓存token（限制多地登录，多设备登录等情况下用）

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

		info, _ := dao.NewDaoHandler(ctx, &daoApple.Account).Filter(g.Map{`account`: account}).GetModel().One()
		if info.IsEmpty() {
			_, err = service.AppleAccount().Create(ctx, data)

		} else {
			_, err = service.AppleAccount().Update(ctx, filter, data)
		}

		return nil
	}),
	)
	if err != nil {
		log.Fatal(err)
	}
}
