package main

import (
	"tService"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"net"
	"os"
	"time"
)

const (
	HOST = "127.0.0.1"
	PORT = "19090"
)

func main() {
	startTime := currentTimeMillis()
	
	//transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	
	transport, err := thrift.NewTSocket(net.JoinHostPort(HOST, PORT))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		os.Exit(1)
	}
	
	useTransport := transportFactory.GetTransport(transport)
	client := tService.NewUserServiceClientFactory(useTransport, protocolFactory)
	if err := transport.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to "+HOST+":"+PORT, " ", err)
		os.Exit(1)
	}
	defer transport.Close()
	
	r1, err := client.GetUserNameById(20)
	fmt.Println(err)
	fmt.Println("GOClient Call->", r1)

	endTime := currentTimeMillis()
	fmt.Printf("本次调用用时:%d-%d=%d毫秒\n", endTime, startTime, (endTime - startTime))
	
}

func currentTimeMillis() int64 {
	return time.Now().UnixNano() / 1000000
}