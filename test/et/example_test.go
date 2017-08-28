package et

import "fmt"

func ExampleHello()  {
	fmt.Println("Hello,Golang~")
	//Output:Hello,Golang~
}

func ExampleHello2()  {
	for i:=0; i<3; i++  {
		fmt.Println("Hello,Golang~")
	}

	//Output:Hello,Golang~
	//Hello,Golang~
	//Hello,Golang~
}