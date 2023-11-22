package controller

import (
	"api/api"
	apiCurrent "api/api/app"
	apiConfig "api/api/platform/platform"
	"api/internal/cache"
	"api/internal/dao"
	daoPlatform "api/internal/dao/platform"
	daoUser "api/internal/dao/user"
	"api/internal/utils"
	"api/internal/utils/sms"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"net/smtp"
)

type Sms struct{}

func NewSms() *Sms {
	return &Sms{}
}

// 发送短信
func (controllerThis *Sms) Send(ctx context.Context, req *apiCurrent.SmsSendReq) (res *api.CommonNoDataRes, err error) {
	userColumns := daoUser.User.Columns()
	phone := req.Phone
	switch req.UseScene {
	case 0, 2: //登录，密码找回
		info, _ := dao.NewDaoHandler(ctx, &daoUser.User).Filter(g.Map{userColumns.Phone: phone}).GetModel().One()
		if info.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39990000, ``)
			return
		}
		if info[userColumns.IsStop].Uint() == 1 {
			err = utils.NewErrorCode(ctx, 39990002, ``)
			return
		}
	case 1: //注册
		info, _ := dao.NewDaoHandler(ctx, &daoUser.User).Filter(g.Map{userColumns.Phone: phone}).GetModel().One()
		if !info.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39990004, ``)
			return
		}
	case 3: //密码修改
		loginInfo := utils.GetCtxLoginInfo(ctx)
		if loginInfo.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39994000, ``)
			return
		}
		phone = loginInfo[userColumns.Phone].String()
		if phone != `` {
			err = utils.NewErrorCode(ctx, 39990007, ``)
			return
		}
	case 4: //绑定手机
		loginInfo := utils.GetCtxLoginInfo(ctx)
		if loginInfo.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39994000, ``)
			return
		}
		if loginInfo[userColumns.Phone].String() != `` {
			err = utils.NewErrorCode(ctx, 39990005, ``)
			return
		}
		info, _ := dao.NewDaoHandler(ctx, &daoUser.User).Filter(g.Map{userColumns.Phone: phone}).GetModel().One()
		if !info.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39990006, ``)
			return
		}
	case 5: //解绑手机
		loginInfo := utils.GetCtxLoginInfo(ctx)
		if loginInfo.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39994000, ``)
			return
		}
		phone = loginInfo[userColumns.Phone].String()
		if phone == `` {
			err = utils.NewErrorCode(ctx, 39990007, ``)
			return
		}
	}

	smsCode := grand.Digits(4)
	err = sms.NewSms(ctx).Send(phone, smsCode)
	if err != nil {
		return
	}
	err = cache.NewSms(ctx, phone, req.UseScene).Set(smsCode, 5*60)
	return
}

func (controllerThis *Sms) EmailSend(ctx context.Context, req *apiCurrent.SmsEmailSendReq) (res *api.CommonNoDataRes, err error) {
	userColumns := daoUser.User.Columns()
	email := req.Email

	switch req.UseScene {
	case 1: //注册
		info, _ := dao.NewDaoHandler(ctx, &daoUser.User).Filter(g.Map{userColumns.Email: email}).GetModel().One()
		if !info.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39990004, ``)
			return
		}
	}

	value, err := cache.NewSms(ctx, email, req.UseScene).Get()

	//if value != `` {
	//	err = utils.NewErrorCode(ctx, 39990010, ``)
	//	return
	//}

	println("value", value)

	smsCode := grand.Digits(6)

	println("smsCode", smsCode)

	var configKeyArr []string
	configKeyArr = []string{"smtpHost", "smtpPort", "smtpEmail", "smtpPwd"}
	config, err := daoPlatform.Config.Get(ctx, configKeyArr)
	if err != nil {
		return
	}

	//res = &apiPlatform.ConfigGetRes{}
	var configModel apiConfig.Config
	gconv.Struct(config, &configModel)
	smtpServer := *configModel.SmtpHost

	smtpPort := *configModel.SmtpPort

	senderEmail := *configModel.SmtpEmail

	password := *configModel.SmtpPwd
	//
	//senderEmail := config["smtpEmail"]
	//password := config["smtpPwd"]
	//
	//fmt.Println("smtpHost", smtpServer)
	//
	//fmt.Println("smtpPort", smtpPort)
	//
	//fmt.Println("senderEmail", senderEmail)
	//
	//fmt.Println("password", password)

	// Safely convert config values to strings
	//smtpServer, ok := config["smtpHost"].(string)
	//if !ok {
	//	fmt.Println("smtpHost is not a string")
	//	return
	//}
	//
	//smtpPort, ok := config["smtpPort"].(int)
	//if !ok {
	//	fmt.Println("smtpPort is not a string")
	//	return
	//}
	//
	//senderEmail, ok := config["smtpEmail"].(string)
	//if !ok {
	//	fmt.Println("smtpEmail is not a string")
	//	return
	//}
	//
	//password, ok := config["smtpPwd"].(string)
	//if !ok {
	//	fmt.Println("smtpPwd is not a string")
	//	return
	//}
	//
	//port := string(smtpPort)

	subject := "Subject: Your Subject Here\n"
	body := "Hello, thank you for registering, your verification code is " + smsCode

	// 创建邮件内容
	message := []byte(subject + "\n" + body)

	// 设置身份验证信息
	auth := smtp.PlainAuth("", senderEmail, password, smtpServer)

	// 发送邮件
	err = smtp.SendMail(smtpServer+":"+smtpPort, auth, senderEmail, []string{email}, message)
	if err != nil {
		fmt.Println("Failed to authenticate:", err)
		return
	}

	if err != nil {
		return
	}
	err = cache.NewSms(ctx, email, req.UseScene).Set(smsCode, 5*60)
	return
}
