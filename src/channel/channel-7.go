package main

import (
	"fmt"
	"time"
)

type Counter struct {
	count 	int
	name	string
}

func (counter *Counter)String() string  {
	return fmt.Sprintf("count:%d: name:%s", counter.count, counter.name)
}

// mapChan的元素类型属于引用类型，因此接收方对元素值的副本的修改会影响到发送方持有的源值。
// mapChan2一个结构体类型的值中包含了类型为切片的字段，在这些情况下，需要特别注意。
// 结构体 [[[非指针类型]]] 时，接收方对元素副本的修改 [[[不会影响]]] 到发送方持有的源值
// 结构体 [[[指针类型 ]]] 时，接收方对元素副本的修改 [[[会影响]]] 到发送方持有的源值

// 非指针
//var mapChan2 = make(chan map[string]Counter, 1)
// 指针类型
var mapChan2 = make(chan map[string]*Counter, 1)

func main()  {

	syncChan := make(chan struct{}, 2)
	go func() { //用于演示接收操作

		for {
			if elem, ok := <-mapChan2; ok {
				//fmt.Printf("The count map: %v.[receiver]\n", elem)
				counter := elem["count"]
				counter.count++
				counter.name = "jingyanlei"
				fmt.Printf("The count map: %v.[receiver]\n", counter)
			} else {
				break
			}
		}
		fmt.Println("Stoped.[receiver]")
		syncChan <- struct{}{}
	}()

	go func() { // 用于演示发送操作
		// 非指针
		//countMap2 := map[string]Counter{
		//	"count":Counter{},
		//}
		// 指针
		countMap2 := map[string]*Counter{
			"count":&Counter{},
		}
		for i:=0; i<5; i++ {
			count := countMap2
			fmt.Printf("The count map: %v.[sender]\n", count)
			mapChan2 <- countMap2
			time.Sleep(time.Millisecond)
			count = countMap2
			fmt.Printf("The count map: %v.[sender]\n", count)
			fmt.Println("---------")
		}
		close(mapChan2)
		syncChan <- struct{}{}
	}()

	<- syncChan
	<- syncChan
}
