// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IOrdersOrdersActions interface {
		// 新增
		Create(ctx context.Context, data map[string]interface{}) (id int64, err error)
		// 修改
		Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (row int64, err error)
		// 删除
		Delete(ctx context.Context, filter map[string]interface{}) (row int64, err error)
	}
	IOrders interface {
		// 新增
		Create(ctx context.Context, data map[string]interface{}) (id int64, err error)
		// 修改
		Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (row int64, err error)
		// 删除
		Delete(ctx context.Context, filter map[string]interface{}) (row int64, err error)
	}
)

var (
	localOrders              IOrders
	localOrdersOrdersActions IOrdersOrdersActions
)

func Orders() IOrders {
	if localOrders == nil {
		panic("implement not found for interface IOrders, forgot register?")
	}
	return localOrders
}

func RegisterOrders(i IOrders) {
	localOrders = i
}

func OrdersOrdersActions() IOrdersOrdersActions {
	if localOrdersOrdersActions == nil {
		panic("implement not found for interface IOrdersOrdersActions, forgot register?")
	}
	return localOrdersOrdersActions
}

func RegisterOrdersOrdersActions(i IOrdersOrdersActions) {
	localOrdersOrdersActions = i
}
