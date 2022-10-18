package singleton

import (
	"fmt"
	"sync"
)

// 使用懒惰模式的单例模式，使用双重检查加锁保证线程安全
// Singleton 是单例模式接口，导出的
// 通过该接口可以避免 GetInstance 返回一个包私有类型的指针
type Singleton interface {
	foo()
}

// singleton 是单例模式类，包私有的
type singleton struct{}

func (s singleton) foo() {}

var (
	instance *singleton
	once     sync.Once
)

//GetInstance 用于获取单例模式对象
func GetInstance() Singleton {
	// 多次调用，一次执行
	once.Do(func() {
		fmt.Println("invoked once")
		instance = &singleton{}
	})

	return instance
}
