package main

import (
	"fmt"
	"go-wire/internal/di"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	_, closeApp, err := di.InitApp()
	if err != nil {
		panic(err)
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		fmt.Printf("get a signal %s\n", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			closeApp()
			fmt.Println("demo1 exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
