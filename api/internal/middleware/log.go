package middleware

import (
	dao "api/internal/model/dao/log"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
)

func Log(r *ghttp.Request) {
	startTime := gtime.Now().UnixMicro()

	r.Middleware.Next()

	endTime := gtime.Now().UnixMicro()
	runTime := (float64(endTime) - float64(startTime)) / 1000
	data := map[string]interface{}{
		"requestUrl":    r.GetUrl(),
		"requestData":   r.GetMap(),
		"requestHeader": r.Header,
		"runTime":       runTime,
		//"responseBody":  gjson.MustEncodeString(map[string]interface{}{}),
		"responseBody": map[string]interface{}{},
	}
	//dao.Request.Ctx(r.GetCtx()).Handler(dao.Request.ParseInsert(data))
	dao.Request.Ctx(r.GetCtx()).Data(data).Insert()
}
