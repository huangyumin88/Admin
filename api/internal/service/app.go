// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IAppCardCategories interface {
		// 新增
		Create(ctx context.Context, data map[string]interface{}) (id int64, err error)
		// 修改
		Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (row int64, err error)
		// 删除
		Delete(ctx context.Context, filter map[string]interface{}) (row int64, err error)
	}
	IAppCardCategoriesSub interface {
		// 新增
		Create(ctx context.Context, data map[string]interface{}) (id int64, err error)
		// 修改
		Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (row int64, err error)
		// 删除
		Delete(ctx context.Context, filter map[string]interface{}) (row int64, err error)
	}
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
	localAppCardCategories    IAppCardCategories
	localAppCardCategoriesSub IAppCardCategoriesSub
	localAppCardCountries     IAppCardCountries
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

func AppCardCategories() IAppCardCategories {
	if localAppCardCategories == nil {
		panic("implement not found for interface IAppCardCategories, forgot register?")
	}
	return localAppCardCategories
}

func RegisterAppCardCategories(i IAppCardCategories) {
	localAppCardCategories = i
}

func AppCardCategoriesSub() IAppCardCategoriesSub {
	if localAppCardCategoriesSub == nil {
		panic("implement not found for interface IAppCardCategoriesSub, forgot register?")
	}
	return localAppCardCategoriesSub
}

func RegisterAppCardCategoriesSub(i IAppCardCategoriesSub) {
	localAppCardCategoriesSub = i
}
