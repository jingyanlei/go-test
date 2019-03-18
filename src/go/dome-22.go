package main

import "log"

//22. log.Fatal 和 log.Panic 不只是 log

//log 标准库提供了不同的日志记录等级，与其他语言的日志库不同，Go 的 log 包在调用 Fatal*()、Panic*() 时能做更多日志外的事，如中断程序的执行等：

func main() {
	log.Fatalln("Fatal level log: log entry")        // 输出信息后，程序终止执行
	log.Println("Nomal level log: log entry")
}