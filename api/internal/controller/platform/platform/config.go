package controller

import (
	"api/api"
	apiPlatform "api/api/platform/platform"
	daoPlatform "api/internal/dao/platform"
	"api/internal/service"
	"api/internal/utils"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
	"net/smtp"
)

type Config struct{}

func NewConfig() *Config {
	return &Config{}
}

func (controllerThis *Config) Test(ctx context.Context, req *apiPlatform.ConfigSmtpTestReq) (res *api.CommonNoDataRes, err error) {

	smtpServer := *req.SmtpHost
	smtpPort := *req.SmtpPort
	//smtpPort, err := strconv.Atoi(*req.SmtpPort)
	//if err != nil {
	//	// handle error
	//	return
	//}

	senderEmail := *req.SmtpEmail
	password := *req.SmtpPwd
	recipient := *req.SmtpTestEmail
	println(smtpServer)
	println(smtpPort)
	println(senderEmail)
	println(password)
	println(recipient)
	// 设置收件人和邮件内容

	subject := "Subject: Your Subject Here\n"
	body := "This is the body of the email."

	// 创建邮件内容
	message := []byte(subject + "\n" + body)

	// 设置身份验证信息
	auth := smtp.PlainAuth("", senderEmail, password, smtpServer)

	// 发送邮件
	err = smtp.SendMail(smtpServer+":"+smtpPort, auth, senderEmail, []string{recipient}, message)
	if err != nil {
		fmt.Println("Failed to authenticate:", err)
		return
	}

	//// 连接SMTP服务器
	//c, err := smtp.Dial(host + ":" + strconv.Itoa(port))
	//if err != nil {
	//	fmt.Println("Failed to connect to SMTP server:", err)
	//	return
	//}
	//
	//auth := smtp.PlainAuth("", addr, pwd, host)
	//
	//if err != nil {
	//	fmt.Println("Failed to authenticate:", err)
	//	return
	//}
	//
	//// 身份验证
	//err = c.Auth(auth)
	//if err != nil {
	//	// handle error
	//
	//	return
	//}
	//
	//// 设置邮箱地址
	//err = c.Mail(addr)
	//err = c.Rcpt("477603590@qq.com")
	//
	//// 构建邮件内容
	//w, err := c.Data()
	//w.Write([]byte("This is the email body."))
	//
	//// 发送邮件
	//err = w.Close()
	//
	//// 关闭连接
	//c.Quit()

	return
}

// 获取
func (controllerThis *Config) Get(ctx context.Context, req *apiPlatform.ConfigGetReq) (res *apiPlatform.ConfigGetRes, err error) {
	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `platformConfigLook`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	config, err := daoPlatform.Config.Get(ctx, *req.ConfigKeyArr)
	if err != nil {
		return
	}

	res = &apiPlatform.ConfigGetRes{}
	gconv.Struct(config, &res.Config)
	return
}

// 保存
func (controllerThis *Config) Save(ctx context.Context, req *apiPlatform.ConfigSaveReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	config := gconv.Map(req)
	if len(config) == 0 {
		err = utils.NewErrorCode(ctx, 89999999, ``)
		return
	}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `platformConfigSave`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	err = daoPlatform.Config.Save(ctx, config)
	return
}
