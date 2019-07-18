package main

import "fmt"

// 什么是头等函数？
// 支持头等函数的语言允许将函数分配给变量，作为参数传递给其他函数并从其他函数返回。 Go就支持这一功能。

//用户定义的函数类型
type add func(a int, b int) int

// 将函数作为参数传递给其他函数
func simple(f func(a, b int) int) {
	fmt.Println(f(2,2))
}

// 从其他函数返回函数
func simple2() func(a, b int) int {
	f := func(a, b int) int {
		return a+b
	}
	return f
}

func appendStr() func(string) string {
	t := "hello"
	c := func(b string) string {
		t = t+"_"+b
		return t
	}
	return c
}

func main()  {

	// 匿名函数
	a := func() {
		fmt.Println("hello world first class function")
	}
	a()

	fmt.Printf("%T\n", a)


	// 匿名函数直接执行
	func() {
		fmt.Println("hello world first class function")
	}()

	//参数传递给匿名函数
	func(n string) {
		fmt.Println("Welcome", n)
	}("L")

	// 用户定义的函数类型
	var add add = func(a int, b int) int {
		return a + b
	}
	s := add(1,2)
	fmt.Println("sum", s)

	fmt.Println("---")

	// 高阶函数
	//从维基百科中关于高阶函数的定义我们可以知道，高阶函数它至少有以下的一个功能
	//
	//接受一个或多个函数作为输入
	//输出一个函数
	f := func(a, b int) int {
		return a + b
	}
	fmt.Println("高阶函数")
	simple(f)
	fmt.Println("---")

	fmt.Println("从其他函数返回函数")
	s2 := simple2()
	fmt.Println(s2(2,3))
	fmt.Println("---")

	fmt.Println("闭包\n闭包是一个特殊的匿名函数。 它是一个可以访问函数体外的变量的匿名函数。每个闭包都绑定到它自己的周围变量")
	s3 := 5
	func() {
		fmt.Println("s3=", s3)
		s3 = 55
	}()
	fmt.Println("s3=", s3)

	a1 := appendStr()
	b1 := appendStr()
	fmt.Println(a1("A"))
	fmt.Println(b1("B"))

	fmt.Println(a1("A1"))
	fmt.Println(b1("B1"))

	a2 := appendStr()
	fmt.Println(a2("A2"))


}