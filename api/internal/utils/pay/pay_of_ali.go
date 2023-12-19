package pay

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
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
	return
}

func (payThis *PayOfAli) Notify() (notifyInfo NotifyInfo, err error) {
	return
}

func (payThis *PayOfAli) NotifyRes(failMsg string) {
	resData := `success` //success:	成功；fail：失败
	if failMsg != `` {
		resData = `fail`
	}
	g.RequestFromCtx(payThis.Ctx).Response.Write(resData)
}
