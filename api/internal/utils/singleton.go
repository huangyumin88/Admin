package utils

import "sync"

type singleton struct {
	// 定义你的字段
	isAutomLogin bool
	StopCh       chan struct{}
}

var (
	instance *singleton
	once     sync.Once
)

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{
			// 初始化字段
			isAutomLogin: false,
			StopCh:       make(chan struct{}),
		}
	})
	return instance
}
