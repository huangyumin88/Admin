// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IBanksBankCards interface {
		// 新增
		Create(ctx context.Context, data map[string]interface{}) (id int64, err error)
		// 修改
		Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (row int64, err error)
		// 删除
		Delete(ctx context.Context, filter map[string]interface{}) (row int64, err error)
	}
	IBanks interface {
		// 新增
		Create(ctx context.Context, data map[string]interface{}) (id int64, err error)
		// 修改
		Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (row int64, err error)
		// 删除
		Delete(ctx context.Context, filter map[string]interface{}) (row int64, err error)
	}
	IBanksBanksUser interface {
		// 新增
		Create(ctx context.Context, data map[string]interface{}) (id int64, err error)
		// 修改
		Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (row int64, err error)
		// 删除
		Delete(ctx context.Context, filter map[string]interface{}) (row int64, err error)
	}
)

var (
	localBanksBanksUser IBanksBanksUser
	localBanksBankCards IBanksBankCards
	localBanks          IBanks
)

func BanksBankCards() IBanksBankCards {
	if localBanksBankCards == nil {
		panic("implement not found for interface IBanksBankCards, forgot register?")
	}
	return localBanksBankCards
}

func RegisterBanksBankCards(i IBanksBankCards) {
	localBanksBankCards = i
}

func Banks() IBanks {
	if localBanks == nil {
		panic("implement not found for interface IBanks, forgot register?")
	}
	return localBanks
}

func RegisterBanks(i IBanks) {
	localBanks = i
}

func BanksBanksUser() IBanksBanksUser {
	if localBanksBanksUser == nil {
		panic("implement not found for interface IBanksBanksUser, forgot register?")
	}
	return localBanksBanksUser
}

func RegisterBanksBanksUser(i IBanksBanksUser) {
	localBanksBanksUser = i
}
