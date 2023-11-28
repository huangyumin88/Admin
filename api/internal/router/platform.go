package router

import (
	"github.com/gogf/gf/v2/net/ghttp"

	"api/internal/controller"
	controllerCurrent "api/internal/controller/platform"
	controllerAuth "api/internal/controller/platform/auth"
	controllerMy "api/internal/controller/platform/my"
	controllerPlatform "api/internal/controller/platform/platform"
	controllerUser "api/internal/controller/platform/user"
	controllerApp "api/internal/controller/platform/app"
	controllerBanks "api/internal/controller/platform/banks"
	controllerOrders "api/internal/controller/platform/orders"
	"api/internal/middleware"
)

func InitRouterPlatform(s *ghttp.Server) {
	s.Group(`/platform`, func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Scene)

		// 无需验证登录身份
		group.Group(``, func(group *ghttp.RouterGroup) {
			group.Group(`/login`, func(group *ghttp.RouterGroup) {
				group.Bind(controllerCurrent.NewLogin())
			})
		})

		// 需验证登录身份
		group.Group(``, func(group *ghttp.RouterGroup) {
			group.Middleware(middleware.SceneLoginOfPlatform(true))

			group.Group(`/upload`, func(group *ghttp.RouterGroup) {
				controllerThis := controller.NewUpload()
				group.Bind(controllerThis.Sign)
			})

			group.Group(`/my`, func(group *ghttp.RouterGroup) {
				group.Bind(controllerMy.NewProfile())
				group.Bind(controllerMy.NewMenu())
				group.Bind(controllerMy.NewAction())
			})

			group.Group(`/auth`, func(group *ghttp.RouterGroup) {
				group.Bind(controllerAuth.NewAction())
				group.Bind(controllerAuth.NewMenu())
				group.Bind(controllerAuth.NewRole())
				group.Bind(controllerAuth.NewScene())
			})

			group.Group(`/platform`, func(group *ghttp.RouterGroup) {
				group.Bind(controllerPlatform.NewAdmin())
				group.Bind(controllerPlatform.NewConfig())
			})

			group.Group(`/user`, func(group *ghttp.RouterGroup) {
				group.Bind(controllerUser.NewWallets())
				group.Bind(controllerUser.NewUser())
			})

			group.Group(`/app`, func(group *ghttp.RouterGroup) {
				group.Bind(controllerApp.NewCardCategoriesSub())
				group.Bind(controllerApp.NewCardCategories())
				group.Bind(controllerApp.NewCardCountries())
			})

			group.Group(`/banks`, func(group *ghttp.RouterGroup) {
				group.Bind(controllerBanks.NewBankCards())
				group.Bind(controllerBanks.NewBanksUser())
				group.Bind(controllerBanks.NewBanks())
			})

			group.Group(`/orders`, func(group *ghttp.RouterGroup) {
				group.Bind(controllerOrders.NewOrders())
			})

			/*--------后端路由自动代码生成锚点（不允许修改和删除，否则将不能自动生成路由）--------*/
		})
	})
}
