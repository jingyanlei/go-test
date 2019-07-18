package main

import "os"

//7. 不小心覆盖了变量
//对从动态语言转过来的开发者来说，简短声明很好用，这可能会让人误会 := 是一个赋值操作符。
//
//如果你在新的代码块中像下边这样误用了 :=，编译不会报错，但是变量不会按你的预期工作：

func main() {
	x := 1
	println(x)        // 1
	{
		println(x)    // 1
		x := 2
		println(x)    // 2    // 新的 x 变量的作用域只在代码块内部
		x = 3
		println(x)
	}
	println(x)        // 1

	f, err := os.Open("")
	f, err = os.Open("")
	println(f, err)
}

//这是 Go 开发者常犯的错，而且不易被发现。
//
//可使用 vet 工具来诊断这种变量覆盖，Go 默认不做覆盖检查，添加 -shadow 选项来启用：

//> go tool vet -shadow main.go
//main.go:9: declaration of "x" shadows declaration at main.go:5

//注意 vet 不会报告全部被覆盖的变量，可以使用 go-nyet 来做进一步的检测：
// go get github.com/barakmich/go-nyet
//> $GOPATH/bin/go-nyet main.go
//main.go:10:3:Shadowing variable `x`

//go-nyet ./...
//# or
//go-nyet subpackage
//# or
//go-nyet file.go