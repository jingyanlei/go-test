package main

import "fmt"

//43. 旧 slice
//当你从一个已存在的 slice 创建新 slice 时，二者的数据指向相同的底层数组。如果你的程序使用这个特性，那需要注意 "旧"（stale） slice 问题。
//
//某些情况下，向一个 slice 中追加元素而它指向的底层数组容量不足时，将会重新分配一个新数组来存储数据。而其他 slice 还指向原来的旧底层数组。

// 超过容量将重新分配数组来拷贝值、重新存储
func main() {
	s1 := make([]int, 10)
	s1[0] = 1
	s1[1] = 2
	s1[2] = 3
	//s1 := []int{1, 2, 3}
	fmt.Println("s1", len(s1), cap(s1), s1)    // 3 3 [1 2 3 ]
	s2 := s1[1:]
	fmt.Println("s2", len(s2), cap(s2), s2)    // 2 2 [2 3]
	fmt.Println("-----")
	for i := range s2 {
		s2[i] += 20
	}
	// 此时的 s1 与 s2 是指向同一个底层数组的
	fmt.Println("s1", len(s1), cap(s1), s1)        // [1 22 23]
	fmt.Println("s2", len(s2), cap(s2), s2)        // [22 23]
	fmt.Println("-----")

	s2 = append(s2, 4)    // 向容量为 2 的 s2 中再追加元素，此时将分配新数组来存
	for i := range s2 {
		s2[i] += 10
	}
	fmt.Println("s1", len(s1), cap(s1), s1)        // [1 22 23]    // 此时的 s1 不再更新，为旧数据
	fmt.Println("s2", len(s2), cap(s2), s2)        // [32 33 14]
	fmt.Println("-----")

	s3 := s2[2:]
	s3 = append(s3, 5)
	for i := range s3 {
		s3[i] += 10
	}
	// 新分配的slice容量没有增加，新，旧slice值都会更新
	fmt.Println("s2", fmt.Sprintf("%p", s2), len(s2), cap(s2), s2)        // [32 33 14]
	fmt.Println("s3", fmt.Sprintf("%p", s3), len(s3), cap(s3), s3)        // [32 33 14]

	fmt.Println("-----")
	// 新分配的slice容量增加，重新分配置，新slice值更新，旧slice不在更新
	s3 = append(s3, 100,101,102,104,105,106,107,108)
	for i := range s3 {
		s3[i] += 10
	}
	fmt.Println("s2", fmt.Sprintf("%p", s2), len(s2), cap(s2), s2)        // [32 33 14]
	fmt.Println("s3", fmt.Sprintf("%p", s3), len(s3), cap(s3), s3)        // [32 33 14]
}