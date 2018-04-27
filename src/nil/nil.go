package main

import "fmt"

func main()  {

	// 数组
	var ipv4 [4]uint8
	fmt.Println(ipv4)

	var ipv4_2 [4]uint8 = [4]uint8{1,2}
	ipv4_2[3] = 4
	fmt.Println(ipv4_2)

fmt.Println("------------")

	// 切片
	var ips = []string{}
	fmt.Println(len(ips))
	if ips == nil || len(ips) == 0 {
		fmt.Println("切片空")
	} else {
		fmt.Println("切片非空")
	}
	fmt.Println(ips)

fmt.Println("------------")

	var ips_2 = []string{"192.168.1.1", "192.168.1.2"}
	fmt.Println(len(ips_2))
	if ips_2 == nil {
		fmt.Println("切片空")
	} else {
		fmt.Println("切片非空")
	}
	fmt.Println(ips_2)

fmt.Println("------------")

	// 初始化长度为10的切片
	var ips_3 = make([]string, 10)
	fmt.Println(len(ips_3))
	if ips_3 == nil {
		fmt.Println("切片空")
	} else {
		fmt.Println("切片非空")
	}
	// 重新赋值
	ips_3[0] = "abc"
	fmt.Println(ips_3[0])
	// 增加
	ips_3 = append(ips_3, "增加11")
	fmt.Println(ips_3[10])

fmt.Println("------------")

	var maps = map[string]bool{}

	if maps == nil || len(maps) == 0 {
		fmt.Println("map空")
	} else {
		fmt.Println("map非空")
	}
	fmt.Println(len(maps))
	fmt.Println(maps)

}
