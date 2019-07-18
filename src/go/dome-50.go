package main

50. 阻塞的 gorutinue 与资源泄露
在 2012 年 Google I/O 大会上，Rob Pike 的 Go Concurrency Patterns 演讲讨论 Go 的几种基本并发模式，如 完整代码 中从数据集中获取第一条数据的函数：

func First(query string, replicas []Search) Result {
	c := make(chan Result)
	replicaSearch := func(i int) { c <- replicas[i](query) }
	for i := range replicas {
		go replicaSearch(i)
	}
	return <-c
}
在搜索重复时依旧每次都起一个 goroutine 去处理，每个 goroutine 都把它的搜索结果发送到结果 channel 中，channel 中收到的第一条数据会直接返回。

返回完第一条数据后，其他 goroutine 的搜索结果怎么处理？他们自己的协程如何处理？

在 First() 中的结果 channel 是无缓冲的，这意味着只有第一个 goroutine 能返回，由于没有 receiver，其他的 goroutine 会在发送上一直阻塞。如果你大量调用，则可能造成资源泄露。

为避免泄露，你应该确保所有的 goroutine 都能正确退出，有 2 个解决方法：

使用带缓冲的 channel，确保能接收全部 goroutine 的返回结果：
func First(query string, replicas ...Search) Result {
	c := make(chan Result,len(replicas))
	searchReplica := func(i int) { c <- replicas[i](query) }
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}
使用 select 语句，配合能保存一个缓冲值的 channel default 语句：
default 的缓冲 channel 保证了即使结果 channel 收不到数据，也不会阻塞 goroutine

func First(query string, replicas ...Search) Result {
	c := make(chan Result,1)
	searchReplica := func(i int) {
		select {
		case c <- replicas[i](query):
		default:
		}
	}
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}
使用特殊的废弃（cancellation） channel 来中断剩余 goroutine 的执行：
func First(query string, replicas ...Search) Result {
	c := make(chan Result)
	done := make(chan struct{})
	defer close(done)
	searchReplica := func(i int) {
		select {
		case c <- replicas[i](query):
		case <- done:
		}
	}
	for i := range replicas {
		go searchReplica(i)
	}

	return <-c
}
Rob Pike 为了简化演示，没有提及演讲代码中存在的这些问题。不过对于新手来说，可能会不加思考直接使用