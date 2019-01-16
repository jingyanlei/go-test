package main

import "fmt"

type Redis struct {
	Host string
	Port int
}

// 使用值接收器的方法。
func (r Redis)setHost(host string)  {
	r.Host = host
	fmt.Println(r.Host)
}

// 使用指针接收器的方法。
func (r *Redis)setPort(port int)  {
	r.Port = port
}

func (r *Redis)display()  {
	fmt.Println("host:", r.Host, "port:", r.Port)
}

func main()  {
	redis := Redis{
		Host: "127.0.0.1",
		Port: 6379,
	}

	fmt.Println("before host:", redis.Host)
	fmt.Println("before redis:", redis)
	redis.display()
	redis.setHost("10.18.5.130")
	fmt.Println("after host:", redis.Host)
	fmt.Println("before redis:", redis)
	redis.display()
fmt.Println("---")
	fmt.Println("before port:", redis.Port)
	fmt.Println("before redis:", redis)
	redis.display()
	redis.setPort(16379)
	fmt.Println("after port:", redis.Port)
	fmt.Println("before redis:", redis)
	redis.display()


	a := 1
	fmt.Printf("a is address:%v, value: %d\n", &a, a)

	b := a
	fmt.Printf("b is address:%v, value: %d\n", &b, b)

	c := &a
	fmt.Printf("c is address:%v, value: %v\n", c, *c)

	*c = 11
	fmt.Printf("a is address:%v, value: %d\n", &a, a)
	fmt.Printf("c is address:%v, value: %v\n", c, *c)




}

func t3(a int)  {
	a = 2
}