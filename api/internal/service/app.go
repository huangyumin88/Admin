// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IAppCardCountries interface {
		// 新增
		Create(ctx context.Context, data map[string]interface{}) (id int64, err error)
		// 修改
		Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (row int64, err error)
		// 删除
		Delete(ctx context.Context, filter map[string]interface{}) (row int64, err error)
	}
)

var (
	localAppCardCountries IAppCardCountries
)

func AppCardCountries() IAppCardCountries {
	if localAppCardCountries == nil {
		panic("implement not found for interface IAppCardCountries, forgot register?")
	}
	return localAppCardCountries
}

func RegisterAppCardCountries(i IAppCardCountries) {
	localAppCardCountries = i
}
