package main

import "fmt"

func main()  {
	type Talk interface {
		Hello(userName string) string
		Talk(heard string) (saying string, end bool, err error)
	}

	fmt.Println("只要一个数据类型的方法集合中包含Talk接口声明中的所有方法，那么它一定是接口Talk的实现类型。这种接口实现方法完全是非侵入的。")

	var talk Talk = new(myTalk)

	// 若想判断变量talk中的值是否性于*myTalk类型，则可以用类断断言来判断

	if _, ok := talk.(*myTalk); ok != true {
		fmt.Println("*myTalk未实现Talk接口")
	}

	talk.Hello("lei")
	talk.Talk("hi")

}

type myTalk string

func (talk *myTalk) Hello(userName string) string {
	str := "hello:"+userName
	fmt.Println(str)
	return str
}

func (talk *myTalk) Talk(heard string) (saying string, end bool, err error) {
	saying = "说："+heard
	fmt.Println(saying)
	return
}


