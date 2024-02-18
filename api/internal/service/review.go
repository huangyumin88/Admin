// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IReviewImage interface {
		// 新增
		Create(ctx context.Context, data map[string]interface{}) (id int64, err error)
		// 修改
		Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (row int64, err error)
		// 删除
		Delete(ctx context.Context, filter map[string]interface{}) (row int64, err error)
	}
)

var (
	localReviewImage IReviewImage
)

func ReviewImage() IReviewImage {
	if localReviewImage == nil {
		panic("implement not found for interface IReviewImage, forgot register?")
	}
	return localReviewImage
}

func RegisterReviewImage(i IReviewImage) {
	localReviewImage = i
}
