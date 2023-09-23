package router

import (
	controllerMovie "api/internal/controller/web/movie"
	"api/internal/middleware"
	"github.com/gogf/gf/v2/net/ghttp"
)

func InitRouterWeb(s *ghttp.Server) {
	s.Group(`/web`, func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Scene)

		//无需验证登录身份
		group.Group(`/movie`, func(group *ghttp.RouterGroup) {
			group.Bind(controllerMovie.NewMovie())
		})
	})
}
