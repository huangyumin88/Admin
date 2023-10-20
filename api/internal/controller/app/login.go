package controller

import (
	"api/api"
	apiCurrent "api/api/app"
	"api/internal/cache"
	"api/internal/dao"
	daoUser "api/internal/dao/user"
	"api/internal/utils"
	"context"

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
	if g.Validator().Rules(`phone`).Data(req.LoginName).Run(ctx) != nil && g.Validator().Rules(`passport`).Data(req.LoginName).Run(ctx) != nil {
		err = utils.NewErrorCode(ctx, 89990000, ``)
		return
	}

	info, _ := dao.NewDaoHandler(ctx, &daoUser.User).Filter(g.Map{`loginName`: req.LoginName}).GetModel().One()
	if info.IsEmpty() {
		err = utils.NewErrorCode(ctx, 39990000, ``)
		return
	}
	if info[`isStop`].Int() == 1 {
		err = utils.NewErrorCode(ctx, 39990002, ``)
		return
	}

	saltDynamic := grand.S(8)
	err = cache.NewSalt(ctx, req.LoginName).Set(saltDynamic, 5)
	if err != nil {
		return
	}
	res = &api.CommonSaltRes{SaltStatic: info[`salt`].String(), SaltDynamic: saltDynamic}
	return
}

// 登录
func (controllerThis *Login) Login(ctx context.Context, req *apiCurrent.LoginLoginReq) (res *api.CommonTokenRes, err error) {
	if g.Validator().Rules(`phone`).Data(req.LoginName).Run(ctx) != nil && g.Validator().Rules(`passport`).Data(req.LoginName).Run(ctx) != nil {
		err = utils.NewErrorCode(ctx, 89990000, ``)
		return
	}

	info, _ := dao.NewDaoHandler(ctx, &daoUser.User).Filter(g.Map{`loginName`: req.LoginName}).GetModel().One()
	if info.IsEmpty() {
		err = utils.NewErrorCode(ctx, 39990000, ``)
		return
	}
	if info[`isStop`].Int() == 1 {
		err = utils.NewErrorCode(ctx, 39990002, ``)
		return
	}
	if req.Password != `` { //密码
		salt, _ := cache.NewSalt(ctx, req.LoginName).Get()
		if salt == `` || gmd5.MustEncrypt(info[`password`].String()+salt) != req.Password {
			err = utils.NewErrorCode(ctx, 39990001, ``)
			return
		}
	} else if req.SmsCode != `` { //短信验证码
		phone := info[`phone`].String()
		if phone == `` {
			err = utils.NewErrorCode(ctx, 39990007, ``)
			return
		}

		smsCode, _ := cache.NewSms(ctx, phone, 0).Get() //使用场景：0登录
		if smsCode == `` || smsCode != req.SmsCode {
			err = utils.NewErrorCode(ctx, 39990008, ``)
			return
		}
	}

	sceneInfo := utils.GetCtxSceneInfo(ctx)
	claims := utils.CustomClaims{LoginId: info[`userId`].Uint()}
	jwt := utils.NewJWT(ctx, sceneInfo[`sceneConfig`].Map())
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
	userDao := daoUser.User
	userColumns := userDao.Columns()
	data := g.Map{}
	if req.Account != `` {
		info, _ := dao.NewDaoHandler(ctx, &userDao).Filter(g.Map{userColumns.Account: req.Account}).GetModel().One()
		if !info.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39990004, ``)
			return
		}
		data[userColumns.Account] = req.Account
		data[userColumns.Nickname] = req.Account
	}
	if req.Password != `` {
		data[userColumns.Password] = req.Password
	}
	if req.Phone != `` {
		smsCode, _ := cache.NewSms(ctx, req.Phone, 1).Get() //使用场景：1注册
		if smsCode == `` || smsCode != req.SmsCode {
			err = utils.NewErrorCode(ctx, 39990008, ``)
			return
		}

		info, _ := dao.NewDaoHandler(ctx, &userDao).Filter(g.Map{userColumns.Phone: req.Phone}).GetModel().One()
		if !info.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39990004, ``)
			return
		}
		data[userColumns.Phone] = req.Phone
		data[userColumns.Nickname] = req.Phone[:3] + `****` + req.Phone[len(req.Phone)-4:]
	}

	userId, err := dao.NewDaoHandler(ctx, &userDao).Insert(data).GetModel().InsertAndGetId()
	if err != nil {
		return
	}

	sceneInfo := utils.GetCtxSceneInfo(ctx)
	claims := utils.CustomClaims{LoginId: uint(userId)}
	jwt := utils.NewJWT(ctx, sceneInfo[`sceneConfig`].Map())
	token, err := jwt.CreateToken(claims)
	if err != nil {
		return
	}
	// cache.NewToken(ctx, claims.LoginId).Set(token, int64(jwt.ExpireTime)) //缓存token（限制多地登录，多设备登录等情况下用）

	res = &api.CommonTokenRes{Token: token}
	return
}
