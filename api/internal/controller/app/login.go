package controller

import (
	"api/api"
	apiCurrent "api/api/app"
	"api/internal/cache"
	"api/internal/dao"
	daoAuth "api/internal/dao/auth"
	daoUser "api/internal/dao/user"
	"api/internal/service"
	"api/internal/utils"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
)

type Login struct{}

func NewLogin() *Login {
	return &Login{}
}

//// 获取加密盐
//func (controllerThis *Login) Salt(ctx context.Context, req *apiCurrent.LoginSaltReq) (res *api.CommonSaltRes, err error) {
//	if g.Validator().Rules(`phone`).Data(req.LoginName).Run(ctx) != nil && g.Validator().Rules(`passport`).Data(req.LoginName).Run(ctx) != nil {
//		err = utils.NewErrorCode(ctx, 89990000, ``)
//		return
//	}
//
//	userColumns := daoUser.User.Columns()
//	info, _ := dao.NewDaoHandler(ctx, &daoUser.User).Filter(g.Map{`loginName`: req.LoginName}).GetModel().One()
//	if info.IsEmpty() {
//		err = utils.NewErrorCode(ctx, 39990000, ``)
//		return
//	}
//	if info[userColumns.IsStop].Uint() == 1 {
//		err = utils.NewErrorCode(ctx, 39990002, ``)
//		return
//	}
//
//	saltDynamic := grand.S(8)
//	err = cache.NewSalt(ctx, req.LoginName).Set(saltDynamic, 5)
//	if err != nil {
//		return
//	}
//	res = &api.CommonSaltRes{SaltStatic: info[userColumns.Salt].String(), SaltDynamic: saltDynamic}
//	return
//}

// 登录
func (controllerThis *Login) Login(ctx context.Context, req *apiCurrent.LoginLoginReq) (res *api.CommonTokenRes, err error) {
	if g.Validator().Rules(`email`).Data(req.Email).Run(ctx) != nil {
		err = utils.NewErrorCode(ctx, 89990001, ``)
		return
	}

	userColumns := daoUser.User.Columns()
	info, _ := dao.NewDaoHandler(ctx, &daoUser.User).Filter(g.Map{`email`: req.Email}).GetModel().One()
	if info.IsEmpty() {
		err = utils.NewErrorCode(ctx, 39990012, ``)
		return
	}
	if info[userColumns.IsStop].Uint() == 1 {
		err = utils.NewErrorCode(ctx, 39990002, ``)
		return
	}

	if req.Password != `` { //密码
		//salt, _ := cache.NewSalt(ctx, req.Email).Get()
		//if salt == `` || gmd5.MustEncrypt(info[userColumns.Password].String()+salt) != req.Password {
		//	err = utils.NewErrorCode(ctx, 39990001, ``)
		//	return
		//}
		if info[userColumns.Password].String() != req.Password {
			err = utils.NewErrorCode(ctx, 39990001, ``)
			return
		}
	}
	//} else if req.SmsCode != `` { //短信验证码
	//	phone := info[userColumns.Phone].String()
	//	if phone == `` {
	//		err = utils.NewErrorCode(ctx, 39990007, ``)
	//		return
	//	}
	//
	//	smsCode, _ := cache.NewSms(ctx, phone, 0).Get() //使用场景：0登录
	//	if smsCode == `` || smsCode != req.SmsCode {
	//		err = utils.NewErrorCode(ctx, 39990008, ``)
	//		return
	//	}
	//}

	claims := utils.CustomClaims{LoginId: info[daoUser.User.PrimaryKey()].Uint()}
	jwt := utils.NewJWT(ctx, utils.GetCtxSceneInfo(ctx)[daoAuth.Scene.Columns().SceneConfig].Map())
	token, err := jwt.CreateToken(claims)
	if err != nil {
		return
	}
	// cache.NewToken(ctx, claims.LoginId).Set(token, int64(jwt.ExpireTime)) //缓存token（限制多地登录，多设备登录等情况下用）

	res = &api.CommonTokenRes{Token: token}
	return
}

// 注册
func (controllerThis *Login) Register(ctx context.Context, req *apiCurrent.LoginRegisterReq) (res *api.CommonTokenRes, err error) {

	if g.Validator().Rules(`email`).Data(req.Email).Run(ctx) != nil {
		err = utils.NewErrorCode(ctx, 89990001, ``)
		return
	}

	userColumns := daoUser.User.Columns()
	data := g.Map{}
	if req.Email != `` {
		info, _ := dao.NewDaoHandler(ctx, &daoUser.User).Filter(g.Map{userColumns.Account: req.Email}).GetModel().One()
		if !info.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39990004, ``)
			return
		}
		data[userColumns.Email] = req.Email
	}

	emailCode := req.EmailCode

	value, err := cache.NewSms(ctx, req.Email, 1).Get()

	if emailCode != value {

		err = utils.NewErrorCode(ctx, 39990011, ``)
		return
	}

	if req.Password != `` {
		data[userColumns.Password] = req.Password
	}
	if req.Phone != `` {
		//smsCode, _ := cache.NewSms(ctx, req.Phone, 1).Get() //使用场景：1注册
		//if smsCode == `` || smsCode != req.SmsCode {
		//	err = utils.NewErrorCode(ctx, 39990008, ``)
		//	return
		//}
		//
		//info, _ := dao.NewDaoHandler(ctx, &daoUser.User).Filter(g.Map{userColumns.Phone: req.Phone}).GetModel().One()
		//if !info.IsEmpty() {
		//	err = utils.NewErrorCode(ctx, 39990004, ``)
		//	return
		//}
		data[userColumns.Phone] = req.Phone
		//data[userColumns.Nickname] = req.Phone[:3] + `****` + req.Phone[len(req.Phone)-4:]
	}

	if req.Username != `` {
		data[userColumns.Account] = req.Username
		data[userColumns.Nickname] = req.Username
	}

	if req.Country != `` {
		data[userColumns.Country] = req.Country
	}

	if req.ReferralCode != `` {
		data[userColumns.ReferralCode] = req.ReferralCode
	}

	userId, err := dao.NewDaoHandler(ctx, &daoUser.User).Insert(data).GetModel().InsertAndGetId()
	if err != nil {
		return
	}

	data1 := make(map[string]interface{})
	data1["user_id"] = userId

	// 创建钱包
	walletId, err := service.UserWallets().Create(ctx, data1)
	daoThis := daoUser.User
	daoUserWallets := daoUser.Wallets
	if err != nil {

		service.User().Delete(ctx, g.Map{daoThis.PrimaryKey(): userId})
		return
	}

	data2 := make(map[string]interface{})
	data2["wallet_id"] = walletId

	fmt.Println("wallet_id", walletId)

	filter := map[string]interface{}{`userId`: userId}
	/**--------参数处理 结束--------**/

	_, err = service.User().Update(ctx, filter, data2)

	if err != nil {

		service.User().Delete(ctx, g.Map{daoThis.PrimaryKey(): userId})
		service.UserWallets().Delete(ctx, g.Map{daoUserWallets.PrimaryKey(): walletId})
		return
	}

	claims := utils.CustomClaims{LoginId: uint(userId)}
	jwt := utils.NewJWT(ctx, utils.GetCtxSceneInfo(ctx)[daoAuth.Scene.Columns().SceneConfig].Map())
	token, err := jwt.CreateToken(claims)
	if err != nil {
		return
	}
	cache.NewToken(ctx, claims.LoginId).Set(token, int64(jwt.ExpireTime)) //缓存token（限制多地登录，多设备登录等情况下用）

	res = &api.CommonTokenRes{Token: token}
	return
}

// 密码找回
func (controllerThis *Login) PasswordRecovery(ctx context.Context, req *apiCurrent.LoginPasswordRecoveryReq) (res *api.CommonNoDataRes, err error) {

	if g.Validator().Rules(`email`).Data(req.Email).Run(ctx) != nil {
		err = utils.NewErrorCode(ctx, 89990001, ``)
		return
	}

	userColumns := daoUser.User.Columns()
	info, _ := dao.NewDaoHandler(ctx, &daoUser.User).Filter(g.Map{`email`: req.Email}).GetModel().One()
	if info.IsEmpty() {
		err = utils.NewErrorCode(ctx, 39990012, ``)
		return
	}
	if info[userColumns.IsStop].Uint() == 1 {
		err = utils.NewErrorCode(ctx, 39990002, ``)
		return
	}
	msg := "password reset link with instructions has been sent to your mail " + req.Email
	utils.HttpWriteJson(ctx, nil, 0, msg)
	return

}
