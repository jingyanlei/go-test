package main

import (
	"flag"
	"context"
	"github.com/smallnest/rpcx/server"
	"go-test/rpcx/services"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

type Arith struct{}

// the second parameter is not a pointer
func (t *Arith) Mul(ctx context.Context, args services.Args, reply *services.Reply) error {
	reply.C = args.A * args.B
	return nil
}

func main() {
	flag.Parse()
	
	s := server.NewServer()
	s.Register(new(Arith), "")
	s.Serve("tcp", *addr)
}
