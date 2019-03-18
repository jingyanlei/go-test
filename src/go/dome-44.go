package main

44. 类型声明与方法
从一个现有的非 interface 类型创建新类型时，并不会继承原有的方法：

// 定义 Mutex 的自定义类型
type myMutex sync.Mutex

func main() {
	var mtx myMutex
	mtx.Lock()
	mtx.UnLock()
}
mtx.Lock undefined (type myMutex has no field or method Lock)...
如果你需要使用原类型的方法，可将原类型以匿名字段的形式嵌到你定义的新 struct 中：

// 类型以字段形式直接嵌入
type myLocker struct {
	sync.Mutex
}

func main() {
	var locker myLocker
	locker.Lock()
	locker.Unlock()
}
interface 类型声明也保留它的方法集：

type myLocker sync.Locker

func main() {
	var locker myLocker
	locker.Lock()
	locker.Unlock()
}