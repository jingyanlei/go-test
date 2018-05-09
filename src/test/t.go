package test

import (
	"time"
	"fmt"
	"strconv"
)

func test()  {
	//fmt.Println("test")
	time.Sleep(time.Nanosecond)

}

func test2()  {
	i := 1
	name := "jingyanlei"
	age := 18
	fmt.Sprintf("i:%d, name:%s, age:%d", i, name, age)
}

func test3()  {
	i := 1
	name := "jingyanlei"
	age := 18
	_ = "i:"+strconv.Itoa(i)+"name:"+name+"age:"+strconv.Itoa(age)
}

