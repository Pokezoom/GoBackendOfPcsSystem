package middleware

import "sync"

// 单例模式
var (
	singleton     = make([]func(), 0)
	singletonLock = sync.Mutex{}
)

// Register 注册单例
// @param initFunc 注册方法
func Register(initFunc func()) {
	singletonLock.Lock()
	defer singletonLock.Unlock()
	singleton = append(singleton, initFunc)
}

// Build 构建单例
func Build() {
	for _, f := range singleton {
		f()
	}
}
