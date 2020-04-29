package main

import (
	"fmt"
	"go-wire/internal/dao"
	"go-wire/internal/di"
	"go-wire/internal/server/h"
	"go-wire/internal/service"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	db, closeFunc, err := dao.NewDB()
	if err != nil {
		closeFunc()
		fmt.Println("db error", err)
		return
	}
	redis, closeFunc2, err := dao.NewRedis()
	if err != nil {
		closeFunc()
		closeFunc2()
		fmt.Println("redis error", err)
		return
	}
	d, closeFunc3, err := dao.New(db, redis)
	if err != nil {
		closeFunc()
		closeFunc2()
		closeFunc3()
		fmt.Println("daos error", err)
		return
	}
	s, closeFunc4, err := service.New(d)
	if err != nil {
		closeFunc()
		closeFunc2()
		closeFunc3()
		closeFunc4()
		return
	}
	httpServer, closeFunc6, err := h.New()
	if err != nil {
		closeFunc()
		closeFunc2()
		closeFunc3()
		closeFunc4()
		closeFunc6()
	}
	_, closeFunc5, err := di.NewApp(s, httpServer)
	if err != nil {
		closeFunc()
		closeFunc2()
		closeFunc3()
		closeFunc4()
		closeFunc5()
		closeFunc6()
		panic(err)
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		fmt.Println("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			closeFunc()
			closeFunc2()
			closeFunc3()
			closeFunc4()
			closeFunc5()
			closeFunc6()
			fmt.Println("demo1 exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
	// app.Svc.Test()
}
