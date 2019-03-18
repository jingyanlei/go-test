package main

import "fmt"

//6. 不能使用简短声明来设置字段的值
//struct 的变量字段不能使用 := 来赋值以使用预定义的变量来避免解决：



type info struct {
	result int
}

func work() (int, error) {
	return 3, nil
}

/*
// 错误示例
func main() {
	var data info
	data.result, err := work()    // error: non-name data.result on left side of :=
	fmt.Printf("info: %+v\n", data)
}
*/

// 正确示例
func main() {
	var data info
	var err error    // err 需要预声明

	data.result, err = work()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("info: %+v\n", data)
}
