package pay

import (
	"api/internal/utils/common"
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"errors"
	"sort"
	"strings"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type PayOfAli struct {
	Ctx context.Context
	// Host       string `json:"payOfAliHost"`
	Host       string
	AppId      string `json:"payOfAliAppId"`
	SignType   string `json:"payOfAliSignType"`
	PrivateKey string `json:"payOfAliPrivateKey"`
	PublicKey  string `json:"payOfAliPublicKey"`
	NotifyUrl  string `json:"payOfAliNotifyUrl"`
	ReturnUrl  string `json:"payOfAliReturnUrl"`
}

func NewPayOfAli(ctx context.Context, config map[string]interface{}) *PayOfAli {
	payOfAliObj := PayOfAli{
		Ctx:  ctx,
		Host: `https://openapi.alipay.com/gateway.do`,
	}
	gconv.Struct(config, &payOfAliObj)
	return &payOfAliObj
}

func (payThis *PayOfAli) App(payData PayData) (orderInfo PayInfo, err error) {
	data := map[string]string{
		`app_id`: payThis.AppId,
		`method`: `alipay.trade.app.pay`,
		// `format`: `JSON`,
		`charset`:    `utf-8`,
		`sign_type`:  payThis.SignType,
		`timestamp`:  gtime.Now().Format(`Y-m-d H:i:s`),
		`version`:    `1.0`,
		`notify_url`: payThis.NotifyUrl,
		// `app_auth_token`:     ``,
	}
	bizContent := map[string]interface{}{
		`out_trade_no`: payData.OrderNo,
		`total_amount`: payData.Amount,
		`subject`:      payData.Desc,
	}
	data[`biz_content`] = gjson.MustEncodeString(bizContent)

	sign, err := payThis.CreateSign(data)
	if err != nil {
		return
	}
	data[`sign`] = sign

	res, err := g.Client().Post(payThis.Ctx, payThis.Host, data)
	if err != nil {
		return
	}
	defer res.Close()
	resData := gjson.New(res.ReadAllString())
	/* //验证签名（一般不用再验证了）
	err = payThis.VerifySign(resData.Var().MapStrStr(), resData.Get(`sign`).String())
	if err != nil {
		return
	} */
	// if !(resData.Contains(`alipay_trade_app_pay_response.code`) && resData.Get(`alipay_trade_app_pay_response.code`).Int() == 10000) {
	if resData.Get(`alipay_trade_app_pay_response.code`).Int() != 10000 {
		err = errors.New(resData.Get(`alipay_trade_app_pay_response.msg`).String())
		return
	}

	orderInfo.PayStr = resData.Get(`alipay_trade_app_pay_response.trade_no`).String()
	return
}

func (payThis *PayOfAli) Notify() (notifyInfo NotifyInfo, err error) {
	r := g.RequestFromCtx(payThis.Ctx)
	sign := r.Get(`sign`).String()
	data := r.GetFormMapStrStr()
	/* dataTmp := r.Request.Form
	sign := dataTmp.Get(`sign`)
	data := map[string]string{}
	for key := range dataTmp {
		data[key] = dataTmp.Get(key)
	} */
	err = payThis.VerifySign(data, sign)
	if err != nil {
		return
	}

	/* tradeStatus := r.Get(`trade_status`).String()
	if tradeStatus != `TRADE_SUCCESS` && tradeStatus != `TRADE_FINISHED` {
		err = errors.New(`支付失败`)
		return
	} */

	notifyInfo.Amount = r.Get(`total_amount`).Float64()
	notifyInfo.OrderNo = r.Get(`out_trade_no`).String()
	notifyInfo.OrderNoOfThird = r.Get(`trade_no`).String()
	return
}

func (payThis *PayOfAli) NotifyRes(failMsg string) {
	resData := `success` //success:	成功；fail：失败
	if failMsg != `` {
		resData = `fail`
	}
	g.RequestFromCtx(payThis.Ctx).Response.Write(resData)
}

// 拼接签名字符串
func (payThis *PayOfAli) GetDataStr(data map[string]string) (dataStr string) {
	keyArr := []string{}
	for key := range data {
		if key == `sign` /* || key == `sign_type` */ {
			continue
		}
		keyArr = append(keyArr, key)
	}
	sort.Strings(keyArr)
	dataStrArr := []string{}
	for _, key := range keyArr {
		value := strings.TrimSpace(data[key])
		if len(value) > 0 { //过滤空值字段
			dataStrArr = append(dataStrArr, key+`=`+gconv.String(data[key]))
		}
	}
	dataStr = gstr.Join(dataStrArr, `&`)
	return
}

// 生成签名
func (payThis *PayOfAli) CreateSign(data map[string]string) (sign string, err error) {
	privateKey, err := common.ParsePrivateKeyOfRSA(payThis.PrivateKey)
	if err != nil {
		return
	}

	var hashT crypto.Hash
	switch payThis.SignType {
	case `RSA`:
		hashT = crypto.SHA1
	// case `RSA2`:
	default:
		hashT = crypto.SHA256
	}

	hashed := hashT.New()
	hashed.Write([]byte(payThis.GetDataStr(data)))
	hashedData := hashed.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, hashT, hashedData)
	if err != nil {
		return
	}

	sign = base64.StdEncoding.EncodeToString(signature)
	return
}

// 验证签名
func (payThis *PayOfAli) VerifySign(data map[string]string, sign string) (err error) {
	publicKey, err := common.ParsePublicKeyOfRSA(payThis.PublicKey)
	if err != nil {
		return
	}

	var hashT crypto.Hash
	switch payThis.SignType {
	case `RSA`:
		hashT = crypto.SHA1
	// case `RSA2`:
	default:
		hashT = crypto.SHA256
	}

	hashed := hashT.New()
	hashed.Write([]byte(payThis.GetDataStr(data)))
	hashedData := hashed.Sum(nil)

	signature, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return
	}

	err = rsa.VerifyPKCS1v15(publicKey, hashT, hashedData, signature)
	return
}
