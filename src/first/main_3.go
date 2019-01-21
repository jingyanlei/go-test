package main

import "fmt"

func iMap(m []int, f func(int) int) []int {
	var r []int
	for _, v := range m {
		r = append(r, f(v))
	}
	return r
}

func main() {

	m := []int{1,2,3,4,5}
	m1 := iMap(m, func(v int) int {
		return v + 1
	})
	fmt.Println(m1)

	m2 := iMap(m, func(i int) int {
		return i * 2
	})
	fmt.Println(m2)

	m3 := iMap(m, func(i int) int {
		return i * 5
	})
	fmt.Println(m3)

}