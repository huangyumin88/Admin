package api

import (
	"github.com/gogf/gf/v2/frame/g"
)

/*--------获取加密盐 开始--------*/
type LoginSaltReq struct {
	g.Meta    `path:"/salt" method:"post" tags:"平台后台/登录" sm:"获取加密盐"`
	LoginName string `json:"loginName" v:"required|length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"账号/手机"`
}

/*--------获取加密盐 结束--------*/

/*--------登录 开始--------*/
type LoginLoginReq struct {
	g.Meta    `path:"/login" method:"post" tags:"平台后台/登录" sm:"登录"`
	LoginName string `json:"loginName" v:"required|length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"账号/手机"`
	Password  string `json:"password" v:"required|size:32" dc:"密码。加密后发送，公式：md5(md5(md5(密码)+静态加密盐)+动态加密盐)"`
}

/*--------登录 结束--------*/

type LoginAppleReq struct {
	g.Meta  `path:"/apple/login" method:"post" tags:"苹果/账号登录" sm:"账号登录"`
	Account *string `json:"account,omitempty" v:"length:1,255" dc:"账号"`
	Pwd     *string `json:"pwd,omitempty" v:"length:1,255" dc:"密码"`
	Code    uint    `json:"code,omitempty" v:"integer|min:1" dc:"ID"`
}
