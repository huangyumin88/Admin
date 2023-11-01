// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IAppleAccount interface {
		// 新增
		Create(ctx context.Context, data map[string]interface{}) (id int64, err error)
		// 修改
		Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (row int64, err error)
		// 删除
		Delete(ctx context.Context, filter map[string]interface{}) (row int64, err error)
	}
)

var (
	localAppleAccount IAppleAccount
)

func AppleAccount() IAppleAccount {
	if localAppleAccount == nil {
		panic("implement not found for interface IAppleAccount, forgot register?")
	}
	return localAppleAccount
}

func RegisterAppleAccount(i IAppleAccount) {
	localAppleAccount = i
}
