package main

import (
	"fmt"
	"errors"
)

func main()  {
	fmt.Println("2.3.4")
	func1()
	func2(10, 2)
	result, err := func3(10, 2)
	fmt.Println(fmt.Sprintf("func 有参数，有返回值，未定义返回值变量名，a/b=%d", result), err)
	result2, err2 := func4(2, 50)
	fmt.Println(fmt.Sprintf("func 有参数，有返回值，定义返回值变量名，a*b=%d", result2), err2)


	i1 := myInt(1)
	i2 := i1.add(2)
	fmt.Println(i1)
	fmt.Println(i2)


	a1 := myInt2(1)
	a2 := a1.add(2)
	fmt.Println(a1, a2)

}

func func1() {
	fmt.Println("func 空参数，无返回值方法声明")
}

func func2(a int, b int) {
	str := fmt.Sprintf("func 有参数，无返回值, a+b=%d", a+b)
	fmt.Println(str)
}

// 未定义返回值变量名
func func3(a int, b int) (int, error) {
	result := a / b
	return result, nil
}

// 定义返回值变量名
func func4(a int, b int) (result int, err error) {
	result = a * b
	return
}

//用于定义二元操作函数的类型
type binaryOperation func(operand1 int, operand2 int) (result int, err error)

// 用于以自定义的方式执行二元操作
func operate(op1 int, op2 int, bop binaryOperation) (result int, err error) {
	if bop == nil {
		err = errors.New("inalid binary operation function")
		return
	}
	return bop(op1, op2)
}


// 方法是函数的一种，实际上就是某个数据类型关际在一起的函数

type myInt int

func(i myInt) add(another int) myInt {
	i = i + myInt(another)
	return i
}

type myInt2 int

func(i2 *myInt2) add(another int) myInt2 {
	*i2 = *i2 + myInt2(another)
	return *i2
}