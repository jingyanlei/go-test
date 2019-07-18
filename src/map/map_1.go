package main

import (
	"fmt"
	"sync"
)

func main()  {
	mapInt := map[int]int{
		1:1,
	}
	fmt.Println("len is", len(mapInt), mapInt)

	mapInt2 := make(map[int]int, 10)
	fmt.Println("len is", len(mapInt2), mapInt2)

	mapInt2[0] = 11
	fmt.Println("len is", len(mapInt2), mapInt2)

	mapInt2[20] = 22
	fmt.Println("len is", len(mapInt2), mapInt2)

	t(mapInt2)
	fmt.Println("len is", len(mapInt2), mapInt2)

	fmt.Println("---")

	mapStr := make(map[string]string)
	fmt.Println("len is", len(mapStr), mapStr)

	mapStr["1"] = "a"
	fmt.Println("len is", len(mapStr), mapStr)

	var wg sync.WaitGroup
	wg.Add(1)
	go func(m map[string]string) {
		t2(m)
		wg.Done()
	}(mapStr)
	wg.Wait()
	fmt.Println("len is", len(mapStr), mapStr)



	a := [...]int{1,2,3}
	fmt.Println(a)
	b := a
	b[0] = 4
	fmt.Println(b)
	fmt.Println(a)

	aa := "aa"
	fmt.Println(aa)
	bb := aa
	bb = "bb"
	fmt.Println(bb)
	fmt.Println(aa)



	aaa := 1
	fmt.Println(aaa)
	t3(&aaa)
	fmt.Println(aaa)




}

func t(m map[int]int) {
	m[5] = 55
}

func t2(m map[string]string) {
	m["5"] = "b"
}

func t3(a *int)  {
	*a++
}