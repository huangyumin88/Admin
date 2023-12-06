package api

import (
	"github.com/gogf/gf/v2/frame/g"
)

/*--------获取加密盐 开始--------*/
//type LoginSaltReq struct {
//	g.Meta    `path:"/salt" method:"post" tags:"APP/登录" sm:"获取加密盐"`
//	LoginName string `json:"loginName" v:"required|max-length:30" dc:"账号/手机"`
//}

/*--------获取加密盐 结束--------*/

/*--------登录 开始--------*/
type LoginLoginReq struct {
	g.Meta   `path:"/login" method:"post" tags:"APP/登录" sm:"登录"`
	Email    string `json:"email" v:"required-without:email" dc:"邮箱"`
	Password string `json:"password" v:"required-without:Password" dc:"密码"`
}

/*--------登录 结束--------*/

/*--------注册 开始--------*/
type LoginRegisterReq struct {
	g.Meta `path:"/register" method:"post" tags:"APP/登录" sm:"注册"`
	//Phone  string `json:"phone,omitempty" v:"required-without:Account|max-length:30|phone" dc:"手机"`
	Phone string `json:"phone,omitempty" v:"" dc:"手机"`
	Email string `json:"email,omitempty" v:"required-without:email" dc:"邮箱"`
	//Account string `json:"account,omitempty" v:"required-without:Phone|max-length:30|passport" dc:"账号"`
	// Password string `json:"password" v:"required-with:Account|lsize:32" dc:"密码。加密后发送，公式：md5(密码)"`
	Country      string `json:"country,omitempty" v:"" dc:"国家"`
	Username     string `json:"username" v:"required|max-length:30" dc:"用户姓名"`
	ReferralCode string `json:"referralCode" v:"required|max-length:30" dc:"推荐码"`
	Password     string `json:"password" v:"required|max-length:30" dc:"密码"`
	EmailCode    string `json:"emailCode" v:"required-with:Phone|size:6" dc:"邮箱验证码"`
}

/*--------注册 结束--------*/

/*--------密码找回 开始--------*/
type LoginPasswordRecoveryReq struct {
	g.Meta `path:"/passwordRecovery" method:"post" tags:"APP/登录" sm:"密码找回"`
	Email  string `json:"email,omitempty" v:"required-without:email" dc:"邮箱"`
}

/*--------密码找回 结束--------*/
