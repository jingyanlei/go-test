package main

import (
	"flag"
	"github.com/smallnest/rpcx/client"
	"go-test/rpcx/services"
	"context"
	"log"
)

var (
	addr2 = flag.String("addr", "localhost:8972", "server address")
)

func main()  {
	flag.Parse()
	
	d := client.NewPeer2PeerDiscovery("tcp@"+*addr2, "")
	xclient := client.NewXClient("Arith", client.Failfast, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()
	
	args := services.Args{
		A: 10,
		B: 20,
	}
	
	reply := &services.Reply{}
	err := xclient.Call(context.Background(), "Mul", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}
	log.Printf("%d * %d = %d", args.A, args.B, reply.C)
}