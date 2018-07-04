package main

import (
	"fmt"
	"tService"
	"github.com/apache/thrift/lib/go/thrift"
	"os"
)

const NetworkAddr  = "172.17.0.2:9090"

type UserService struct {
}

func (this *UserService)GetUserNameById(userId int32) (r string, err error) {
	fmt.Println("server")
	return fmt.Sprintf("%d:test-go", userId), nil
}

func main() {
	//transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	//protocolFactory := thrift.NewTCompactProtocolFactory()
	
	serverTransport, err := thrift.NewTServerSocket(NetworkAddr)
	if err != nil {
		fmt.Println("Error!", err)
		os.Exit(1)
	}
	
	handler := &UserService{}
	//processor := demo.NewBatuThriftProcessor(handler)
	processor := tService.NewUserServiceProcessor(handler)
	
	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Println("thrift server in", NetworkAddr)
	server.Serve()
}