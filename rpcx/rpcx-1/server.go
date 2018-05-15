package main

import (
	"github.com/smallnest/rpcx/server"
	"flag"
	"go-test/rpcx/services"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)


func main()  {
	flag.Parse()
	s := server.NewServer()
	//s.RegisterName("Arith", new(services.Arith), "")
	s.Register(new(services.Arith), "")
	s.Serve("tcp", *addr)
}
