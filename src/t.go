package main

import (
	"fmt"
	"unsafe"
	"reflect"
)

func main()  {
	a := 1
	var b int = 444444444444444
	var d interface{} = 555
	var e interface{} = "aaa"

	fmt.Printf("a is %T, size of is %d\n", a, unsafe.Sizeof(a))
	fmt.Printf("b is %T, size of is %d\n", b, unsafe.Sizeof(b))
	fmt.Printf("d is %T, size of is %d\n", d, unsafe.Sizeof(d))
	fmt.Printf("e is %T, size of is %d\n", e, unsafe.Sizeof(e))

	//var c float32 = 1.1
	c := 10 / 1.1
	fmt.Printf("c is %T, value is %v\n", c, c)

	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(d))
	fmt.Println(reflect.TypeOf(e))

	f, ok := e.(int)
	fmt.Println(f, ok)

	g, ok := e.(string)
	fmt.Println(g, ok)

	switch d.(type) {
	case int:
		fmt.Println("d type is int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("unkown")
	}
}
