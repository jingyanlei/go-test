package email

import (
	"go-test/src/alarm"
	"github.com/pkg/errors"
	"fmt"
)

var ADD_ERROR = errors.New("add3 error")

type Email struct {

}

func (e *Email)Update(a alarm.Alarm) {
	println("email send ok ", a.(*alarm.ConcreteAlarm).GetValue())
	var err error
	defer func() {
		fmt.Printf( "defer add1 err:%v\n", err)
		if err != nil {
			e.add1Re()
		}
	}()
	err = e.add1()
	fmt.Printf("add1 err:%v\n", err)
	defer func() {
		fmt.Printf( "defer add2 err:%v\n", err)
		if err != nil {
			e.add2Re()
		}
	}()
	err = e.add2()
	fmt.Printf("add2 err:%v\n", err)
	defer func() {
		fmt.Printf( "defer add3 err:%v\n", err)
		if err != nil {
			e.add3Re()
		}
	}()
	err = e.add3()
	fmt.Printf("add3 err:%v\n", err)
}

func (e *Email)add1() error {
	println("add 1")
	return nil
}

func (e *Email)add2() error {
	println("add 2")
	return nil
}

func (e *Email)add3() error {
	return ADD_ERROR
	//println("add 3")
}

func (e *Email)add1Re()  {
	println("add1 re")
}

func (e *Email)add2Re()  {
	println("add2 re")

}

func (e *Email)add3Re()  {
	println("add3 re")

}