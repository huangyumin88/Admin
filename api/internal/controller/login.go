package controller

import (
	"api/api"
	apiPlatform "api/api/platform/platform"
	"api/internal/packed"
	"api/internal/service"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

type Login struct{}

func NewLogin() *Login {
	return &Login{}
}

// 获取登录加密字符串(前端登录操作用于加密密码后提交)
func (c *Login) EncryptStr(r *ghttp.Request) {
	sceneCode := packed.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	case `platform`:
		/**--------参数处理 开始--------**/
		var param *api.LoginEncryptReq
		err := r.Parse(&param)
		if err != nil {
			packed.HttpFailJson(r, packed.NewErrorCode(r.GetCtx(), 89999999, err.Error()))
			return
		}
		/**--------参数处理 结束--------**/

		encryptStr, err := service.Login().EncryptStr(r.GetCtx(), sceneCode, param.Account)
		if err != nil {
			packed.HttpFailJson(r, err)
			return
		}
		packed.HttpSuccessJson(r, map[string]interface{}{`encryptStr`: encryptStr}, 0)
	}
}

// 登录
func (c *Login) Login(r *ghttp.Request) {
	sceneCode := packed.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	case `platform`:
		/**--------参数处理 开始--------**/
		var param *api.LoginLoginReq
		err := r.Parse(&param)
		if err != nil {
			packed.HttpFailJson(r, packed.NewErrorCode(r.GetCtx(), 89999999, err.Error()))
			return
		}
		/**--------参数处理 结束--------**/

		token, err := service.Login().Login(r.GetCtx(), sceneCode, param.Account, param.Password)
		if err != nil {
			packed.HttpFailJson(r, err)
			return
		}
		packed.HttpSuccessJson(r, map[string]interface{}{`token`: token}, 0)
	}
}

// 登录用户详情
func (c *Login) Info(r *ghttp.Request) {
	sceneCode := packed.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	case `platform`:
		loginInfo := packed.GetCtxLoginInfo(r.GetCtx())
		packed.HttpSuccessJson(r, map[string]interface{}{`info`: loginInfo}, 0)
	}
}

// 修改个人信息
func (c *Login) Update(r *ghttp.Request) {
	sceneCode := packed.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	case `platform`:
		/**--------参数处理 开始--------**/
		var param *apiPlatform.AdminUpdateSelfReq
		err := r.Parse(&param)
		if err != nil {
			packed.HttpFailJson(r, packed.NewErrorCode(r.GetCtx(), 89999999, err.Error()))
			return
		}
		data := gconv.Map(param)
		if len(data) == 0 {
			packed.HttpFailJson(r, packed.NewErrorCode(r.GetCtx(), 89999999, ``))
			return
		}
		loginInfo := packed.GetCtxLoginInfo(r.GetCtx())
		filter := map[string]interface{}{`id`: loginInfo[`adminId`]}
		/**--------参数处理 结束--------**/

		_, err = service.Admin().Update(r.GetCtx(), data, filter)
		if err != nil {
			packed.HttpFailJson(r, err)
			return
		}
		packed.HttpSuccessJson(r, map[string]interface{}{}, 0)
	}
}

// 用户菜单树
func (c *Login) MenuTree(r *ghttp.Request) {
	sceneCode := packed.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	case `platform`:
		loginInfo := packed.GetCtxLoginInfo(r.GetCtx())
		sceneInfo := packed.GetCtxSceneInfo(r.GetCtx())
		filter := map[string]interface{}{}
		filter[`selfMenu`] = map[string]interface{}{
			`sceneCode`: sceneCode,
			`sceneId`:   sceneInfo[`sceneId`].Int(),
			`loginId`:   loginInfo[`adminId`].Int(),
		}
		field := []string{`menuTree`, `showMenu`}

		list, err := service.Menu().List(r.GetCtx(), filter, field, [][2]string{}, 0, 0)
		if err != nil {
			packed.HttpFailJson(r, err)
			return
		}
		tree := packed.Tree(list, 0, `menuId`, `pid`)
		packed.HttpSuccessJson(r, map[string]interface{}{`tree`: tree}, 0)
	}
}
