package controller

import (
	"api/api"
	apiCurrent "api/api/platform"
	"api/internal/cache"
	"api/internal/dao"
	daoAuth "api/internal/dao/auth"
	daoPlatform "api/internal/dao/platform"
	"api/internal/utils"
	"context"
	"fmt"
	"github.com/tencentyun/tls-sig-api-v2-golang/tencentyun"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
)

type Login struct{}

func NewLogin() *Login {
	return &Login{}
}

const (
	sdkappid = 1600014653
	key      = "cb21e11dbc877362238d4e27bffeee7c3a02da4c8d86b4f0f1e2591fd5c583b2"
)

// 获取加密盐
func (controllerThis *Login) Salt(ctx context.Context, req *apiCurrent.LoginSaltReq) (res *api.CommonSaltRes, err error) {
	adminColumns := daoPlatform.Admin.Columns()
	info, _ := dao.NewDaoHandler(ctx, &daoPlatform.Admin).Filter(g.Map{`loginName`: req.LoginName}).GetModel().One()
	if info.IsEmpty() {
		err = utils.NewErrorCode(ctx, 39990000, ``)
		return
	}
	if info[adminColumns.IsStop].Uint() == 1 {
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
	if info[adminColumns.IsStop].Uint() == 1 {
		err = utils.NewErrorCode(ctx, 39990002, ``)
		return
	}

	fmt.Println("req.Password", req.Password)

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

	imUserId := info[adminColumns.ImUserId].String()
	fmt.Println("imUserId", imUserId)
	sig, err := tencentyun.GenUserSig(sdkappid, key, imUserId, 86400*180)
	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		fmt.Println(sig)
	}

	res = &api.CommonTokenRes{Token: token, IMUserId: imUserId, IMUserSig: sig}
	// cache.NewToken(ctx, claims.LoginId).Set(token, int64(jwt.ExpireTime)) //缓存token（限制多地登录，多设备登录等情况下用）

	//res = &api.CommonTokenRes{Token: token}
	return
}
