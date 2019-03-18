package main

import (
	"fmt"
	"time"
)

//46. for 语句中的迭代变量与闭包函数
//for 语句中的迭代变量在每次迭代中都会重用，即 for 中创建的闭包函数接收到的参数始终是同一个变量，在 goroutine 开始执行时都会得到同一个迭代值：
//
//func main() {
//	data := []string{"one", "two", "three"}
//
//	for _, v := range data {
//		go func() {
//			fmt.Println(v)
//		}()
//	}
//
//	time.Sleep(3 * time.Second)
//	// 输出 three three three
//}
//最简单的解决方法：无需修改 goroutine 函数，在 for 内部使用局部变量保存迭代值，再传参：
//
//func main() {
//	data := []string{"one", "two", "three"}
//
//	for _, v := range data {
//		vCopy := v
//		go func() {
//			fmt.Println(vCopy)
//		}()
//	}
//
//	time.Sleep(3 * time.Second)
//	// 输出 one two three
//}
//另一个解决方法：直接将当前的迭代值以参数形式传递给匿名函数：
//
//func main() {
//	data := []string{"one", "two", "three"}
//
//	for _, v := range data {
//		go func(in string) {
//			fmt.Println(in)
//		}(v)
//	}
//
//	time.Sleep(3 * time.Second)
//	// 输出 one two three
//}
//注意下边这个稍复杂的 3 个示例区别：

type field struct {
	name string
}

func (p *field) print() {
	fmt.Println(p.name)
}
//
//// 错误示例
//func main() {
//	data := []field{{"one"}, {"two"}, {"three"}}
//	for _, v := range data {
//		go v.print()
//	}
//	time.Sleep(3 * time.Second)
//	// 输出 three three three
//}


//// 正确示例
//func main() {
//	data := []field{{"one"}, {"two"}, {"three"}}
//	for _, v := range data {
//		v := v
//		go v.print()
//	}
//	time.Sleep(3 * time.Second)
//	// 输出 one two three
//}

//// 正确示例
func main() {
	data := []*field{{"one"}, {"two"}, {"three"}}
	for _, v := range data {    // 此时迭代值 v 是三个元素值的地址，每次 v 指向的值不同
		go v.print()
	}
	time.Sleep(3 * time.Second)
	// 输出 one two three
}