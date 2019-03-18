package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"log"
)

//36. 关闭 HTTP 连接
//一些支持 HTTP1.1 或 HTTP1.0 配置了 connection: keep-alive 选项的服务器会保持一段时间的长连接。但标准库 "net/http" 的连接默认只在服务器主动要求关闭时才断开，所以你的程序可能会消耗完 socket 描述符。解决办法有 2 个，请求结束后：
//
//直接设置请求变量的 Close 字段值为 true，每次请求结束后就会主动关闭连接。
//设置 Header 请求头部选项 Connection: close，然后服务器返回的响应头部也会有这个选项，此时 HTTP 标准库会主动断开连接。

func checkError(err error) {
	if err != nil{
		log.Fatalln(err)
	}
}

/*
// 主动关闭连接
func main() {
	req, err := http.NewRequest("GET", "http://golang.org", nil)
	checkError(err)

	req.Close = true
	//req.Header.Add("Connection", "close")    // 等效的关闭方式

	resp, err := http.DefaultClient.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	checkError(err)

	body, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	fmt.Println(string(body))
}
*/

//你可以创建一个自定义配置的 HTTP transport 客户端，用来取消 HTTP 全局的复用连接：

func main() {
	tr := http.Transport{DisableKeepAlives: true}
	client := http.Client{Transport: &tr}

	resp, err := client.Get("https://golang.google.cn/")
	if resp != nil {
		defer resp.Body.Close()
	}
	checkError(err)

	fmt.Println(resp.StatusCode)    // 200

	body, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	fmt.Println(len(string(body)))
}
//根据需求选择使用场景：
//
//若你的程序要向同一服务器发大量请求，使用默认的保持长连接。
//若你的程序要连接大量的服务器，且每台服务器只请求一两次，那收到请求后直接关闭连接。或增加最大文件打开数 fs.file-max 的值。