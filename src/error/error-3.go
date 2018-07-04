package main

import (
	"path/filepath"
	"fmt"
)

func main()  {
	files, err := filepath.Glob("[")
	// 直接比较错误类型
	if err != nil && err == filepath.ErrBadPattern {
		fmt.Println(err)
		return
	}
	fmt.Println("Match files", files)
}