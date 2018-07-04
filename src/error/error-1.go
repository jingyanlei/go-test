package main

import (
	"os"
	"fmt"
)

func main()  {
	f, err := os.Open("/test.log")
	if err != nil {
		// 断言错误类型
		if err, ok := err.(*os.PathError); ok {
			fmt.Println("File at path", err.Path, "faild to open")
		}
		return
	}
	fmt.Println(f.Name(), "success")
}