package main

import "fmt"

//25. range 迭代 map
//如果你希望以特定的顺序（如按 key 排序）来迭代 map，要注意每次迭代都可能产生不一样的结果。
//
//Go 的运行时是有意打乱迭代顺序的，所以你得到的迭代结果可能不一致。但也并不总会打乱，得到连续相同的 5 个迭代结果也是可能的，如：

func main() {
	m := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}
	for i:=0; i<1; i++ {
		for k, v := range m {
			fmt.Println(k, v)
		}
		println("----")
	}


}

//如果你去 Go Playground 重复运行上边的代码，输出是不会变的，只有你更新代码它才会重新编译。重新编译后迭代顺序是被打乱的：

