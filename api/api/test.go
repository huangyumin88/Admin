package api

import "github.com/gogf/gf/v2/frame/g"

type TestMetaReq struct {
	g.Meta `path:"/testMeta" tags:"TestMeta" method:"get" summary:"测试"`
}

type TestMetaRes struct {
	g.Meta   `mime:"text/html" example:"string"`
	UserName string `p:"username"  v:"required|length:4,30#请输入账号|账号长度为:{min}到:{max}位"`
}

type TestReq struct {
	UserName string `p:"username"  v:"required|length:4,30#请输入账号|账号长度为:{min}到:{max}位"`
}