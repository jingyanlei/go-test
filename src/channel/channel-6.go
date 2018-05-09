package main

import (
	"fmt"
	"time"
)

// mapChan的元素类型属于引用类型，因此接收方对元素值的副本的修改会影响到发送方持有的源值。
var mapChan = make(chan map[string]int, 1)

func main()  {

	syncChan := make(chan struct{}, 2)
	go func() { //用于演示接收操作

		for {
			if elem, ok := <-mapChan; ok {
				//fmt.Printf("The count map: %v.[receiver]\n", elem)
				elem["count"]++
			} else {
				break
			}
		}
		fmt.Println("Stoped.[receiver]")
		syncChan <- struct{}{}
	}()

	go func() { // 用于演示发送操作
		countMap := make(map[string]int)
		for i:=0; i<5; i++ {
			fmt.Printf("The count map: %v.[sender]\n", countMap)
			mapChan <- countMap
			time.Sleep(time.Millisecond)
			fmt.Printf("The count map: %v.[sender]\n", countMap)
			fmt.Println("---------")
		}
		close(mapChan)
		syncChan <- struct{}{}
	}()

	<- syncChan
	<- syncChan
}
