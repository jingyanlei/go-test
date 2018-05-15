package main

import (
	"fmt"
	"errors"
)

func main()  {

	fmt.Println("只要一个数据类型的方法集合中包含Talk接口声明中的所有方法，那么它一定是接口Talk的实现类型。这种接口实现方法完全是非侵入的。")

	var talk Talk = new(myTalk)

	// 若想判断变量talk中的值是否性于*myTalk类型，则可以用类断断言来判断

	if _, ok := talk.(*myTalk); ok != true {
		fmt.Println("*myTalk未实现Talk接口")
	}

	talk.Hello("lei")
	talk.Talk("hi")

	// 嵌入字段
	p := Peo{}
	p.age = 1
	p.name = "test"
	fmt.Println(p.name)
	fmt.Println(p.age)
	fmt.Println(p)

	var chatbot Chatbot = new(myTalk)

	if _, ok := chatbot.(*myTalk); ok != true {
		fmt.Println("*myTalk未实现接口Chatbot")
	}


	//  结构体，类型可以是接口
	simpelCn := simpelCn{
		name:"abc",
		talk:talk,
	}
	fmt.Println(simpelCn)
	fmt.Println(simpelCn.talk.Hello("zi chen"))

}

type Talk interface {
	Hello(userName string) string
	Talk(heard string) (saying string, end bool, err error)
}

type Chatbot interface {
	Name() string
	Begin() (string, error)
	Talk
	ReportError(err error) string
	End() error
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

func (talk *myTalk)Name() string  {
	return "MyName is Name"
}

func (talk *myTalk)Begin() (string, error)  {
	return "MyName is Begin", errors.New("err")
}

func (talk *myTalk)ReportError(err error) string  {
	return "MyName is ReportError"
}

func (talk *myTalk)End() error  {
	return errors.New("err")
}

// 继续属性
type Peo struct {
	test
}

type test struct {
	name string
	age int``
}

type simpelCn struct {
	name string
	talk Talk
}
