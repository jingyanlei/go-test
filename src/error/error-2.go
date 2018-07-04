package main

import (
	"net"
	"fmt"
)

// 断言底层结构类型并使用方法获取更多信息
func main()  {
	addr, err := net.LookupHost("igbd.cn")
	if err != nil {
		if err, ok := err.(*net.DNSError); ok {
			if err.IsTimeout {
				fmt.Println("timeout error")
			} else if err.IsTemporary {
				fmt.Println("Temporary error")
			} else {
				fmt.Println(err)
			}
		}
		return
	}
	fmt.Println(addr)
}
